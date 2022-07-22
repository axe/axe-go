package job

import (
	"fmt"
	"testing"
)

type EchoJob string

func (j EchoJob) Run(job *Job) {
	fmt.Println(j)
}
func (j EchoJob) IsActive() bool {
	return true
}
func (j EchoJob) IsAlive() bool {
	return true
}

func TestRunner(t *testing.T) {
	job := New(Job{
		Name:  "Test",
		Cost:  5,
		Logic: EchoJob("Hello World"),
	})

	runner := NewRunner(1, 10)
	runner.Add(job)
	runner.Run()
}
