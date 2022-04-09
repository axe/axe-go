package job

import "testing"

func TestRunner(t *testing.T) {
	job := New(Job{
		Name: "Test",
		Cost: 5,
	})

	runner := NewRunner(1, 10)
	runner.Add(job)
	runner.Run()
}
