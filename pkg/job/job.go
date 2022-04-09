package job

import (
	"github.com/axe/axe-go/pkg/ds"
)

/*
Example Jobs:
- Get user input
- Read & apply network packets
- Build & write network packet
- Update physics
- Update navigation mesh
- Update path finding
- Update steering behaviors
- Update state machines
- Update space
- Update audio
- Update particles
- Update animations
- Update interface
- Voxel: update lighting
- Voxel: build chunk mesh
- Load & process assets
- Run scripts
- Compute render state
- Render frame
- Display frame
- Recalibrate Jobs (turn on profiling, update costs & budget)
*/

type Job struct {
	// The unique ID of this job
	ID int
	// The name of this job
	Name string
	// The group of this job which dictates ordering
	Group uint8
	// The cost of this job
	Cost int
	// If this job should be executed asynchronously
	Async bool
	// If this job should only be executed after a job with this ID is
	After int
	// If this job is active
	Active bool
	// If this job can be removed
	Remove bool
	// The minimum milliseconds we should wait between execution
	MinWait int64
	// The maximum milliseconds we should wait between execution
	MaxWait int64
	// The last time this job ran
	LastRun      int64
	LastDuration int64
	// The last computed wait time
	WaitTime int64
	// If the job should stricly follow the wait times. So if the wait time is 20 then the job will try to run 50 times a second exactly.
	Strict bool

	Profile      bool
	ProfileStart int64
	ProfileEnd   int64
	// The logic which executes the job
	Logic JobLogic
}

var _ ds.Sortable[*Job] = &Job{}

type JobLogic interface {
	Run(job *Job)
}

var nextJobID = 0

func New(defaults Job) *Job {
	job := defaults
	if job.ID == 0 {
		job.ID = nextJobID
		nextJobID++
	}
	job.Active = true
	return &job
}

func (job *Job) Run(time int64, profile bool) {
	jobProfile := profile || job.Profile
	if jobProfile {
		job.ProfileStart = currentNanos()
	}
	job.Logic.Run(job)
	if jobProfile {
		job.ProfileEnd = currentNanos()
	}
	if job.Strict && job.MinWait > 0 {
		job.LastDuration = job.MinWait
		job.LastRun += job.MinWait
	} else {
		job.LastDuration = time - job.LastRun
		job.LastRun = time
	}
}

func (job *Job) SetFrequency(frequency int64) {
	job.Strict = true
	job.MinWait = frequency
	job.MaxWait = frequency
}

// -1 = not ready
// 0 = can't wait
// X = can wait this many milliseconds
func (job *Job) GetWaitTime(time int64, planned map[int]bool) int64 {
	if !job.Active {
		return -1
	}
	if job.After != 0 && !planned[job.After] {
		return -1
	}
	since := time - job.LastRun
	if since < job.MinWait {
		return -1
	}
	if job.MaxWait > 0 && since > job.MaxWait {
		return 0
	}
	if job.MinWait > 0 {
		return since - job.MinWait
	}
	if job.MaxWait > 0 {
		return since - job.MaxWait
	}
	return 1
}

func (job *Job) UpdateWaitTime(time int64, planned map[int]bool) {
	job.WaitTime = job.GetWaitTime(time, planned)
}

func (job *Job) Less(other *Job) bool {
	return job.WaitTime < other.WaitTime
}

func (job *Job) IsBehind(time int64) bool {
	since := time - job.LastRun
	if job.Strict && job.MinWait > 0 && since > job.MinWait {
		return true
	}
	if job.MaxWait > 0 && since > job.MaxWait {
		return true
	}
	return false
}
