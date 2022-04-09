package job

import "github.com/axe/axe-go/pkg/ds"

type JobGroup struct {
	Jobs  []*Job
	Ready ds.SortableList[*Job]
}

func (g *JobGroup) Add(job *Job) {
	g.Jobs = append(g.Jobs, job)
	g.Ready.Pad(1)
}
