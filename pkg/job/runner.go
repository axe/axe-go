package job

import (
	"time"

	"github.com/axe/axe-go/pkg/ds"
)

type JobRunner struct {
	Groups     []JobGroup
	Ready      ds.SortableList[*Job]
	Planned    map[int]bool
	Dependency map[int]bool
	// How much we can afford to run
	Budget    int
	TotalCost int
	Profile   bool

	LastDuration int64
	LastRun      int64
	ProfileStart int64
	ProfileEnd   int64
}

func NewRunner(groups int, budget int) *JobRunner {
	return &JobRunner{
		Groups:     make([]JobGroup, groups),
		Budget:     budget,
		LastRun:    currentMillis(),
		Planned:    map[int]bool{},
		Dependency: map[int]bool{},
	}
}

func (r *JobRunner) Add(job *Job) {
	r.Groups[job.Group].Add(job)
	r.Ready.Pad(1)
	job.LastRun = r.LastRun
}
func (r *JobRunner) Create(job Job) *Job {
	created := New(job)
	r.Add(created)
	return created
}
func (r *JobRunner) UnreadyLast() {
	last := r.Ready.Pop()
	r.Groups[last.Group].Ready.Pop()
	r.Unready(last)
}
func (r *JobRunner) Unready(job *Job) {
	r.Planned[job.ID] = false
	r.TotalCost -= job.Cost
}
func (r *JobRunner) UnreadyDependency(id int) {
	for i := 0; i < r.Ready.Size; i++ {
		job := r.Ready.Items[i]
		if job.After == id {
			r.Ready.RemoveAt(i)
			groupReady := r.Groups[job.Group].Ready
			groupReady.RemoveAt(groupReady.IndexOf(job))
			r.Unready(job)
			if r.Dependency[job.ID] {
				r.UnreadyDependency(job.ID)
			}
			i--
		}
	}
}
func (r *JobRunner) SortReady() {
	for _, group := range r.Groups {
		group.Ready.Sort()
	}
	r.Ready.Sort()
}

func (r *JobRunner) Run() {
	now := currentMillis()

	r.TotalCost = 0
	r.Planned = map[int]bool{}
	r.Dependency = map[int]bool{}

	for _, group := range r.Groups {
		group.Ready.Clear()
		for _, job := range group.Jobs {
			job.UpdateWaitTime(now, r.Planned)
			if job.WaitTime >= 0 {
				group.Ready.Add(job)
				r.TotalCost += job.Cost
				r.Planned[job.ID] = true
				if job.After > 0 {
					r.Dependency[job.After] = true
				}
			}
		}
	}

	if r.TotalCost > r.Budget {
		r.Ready.Clear()
		for _, group := range r.Groups {
			group.Ready.Sort()
			r.Ready.AddList(group.Ready.List)
		}
		r.Ready.Sort()

		for r.TotalCost > r.Budget {
			last := r.Ready.Last()
			if last.WaitTime == 0 {
				break
			}
			r.UnreadyLast()

			if r.Dependency[last.ID] {
				r.UnreadyDependency(last.ID)
				r.SortReady()
			}
		}
	}

	r.LastDuration = now - r.LastRun
	r.LastRun = now

	if r.Profile {
		r.ProfileStart = currentNanos()
	}

	for _, group := range r.Groups {
		ready := group.Ready.Items
		readySize := group.Ready.Size
		for i := 0; i < readySize; i++ {
			ready[i].Run(now, r.Profile)
		}
	}

	if r.Profile {
		r.ProfileEnd = currentNanos()
	}
}

func currentMillis() int64 {
	return time.Now().UnixMilli()
}
func currentNanos() int64 {
	return time.Now().UnixNano()
}
