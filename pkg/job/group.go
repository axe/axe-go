package job

import "github.com/axe/axe-go/pkg/ds"

type JobGroup struct {
	Jobs  ds.List[*Job]
	Ready ds.SortableList[*Job]
}

func (g *JobGroup) Add(job *Job) {
	g.Ready.Pad(1)
	g.Jobs.Pad(1)
	g.Jobs.Add(job)
}
