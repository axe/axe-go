package asset

type Status struct {
	Progress float32
	Started  bool
	Done     bool
	Error    error
}

func (status *Status) Reset() {
	status.Progress = 0
	status.Done = false
	status.Error = nil
	status.Started = false
}
func (status *Status) Start() {
	status.Reset()
	status.Started = true
}
func (status *Status) Fail(err error) {
	status.Done = true
	status.Started = true
	status.Error = err
}
func (status *Status) Success() {
	status.Done = true
	status.Started = true
	status.Progress = 1
}
func (status *Status) IsSuccess() bool {
	return status.Done && status.Error == nil
}
