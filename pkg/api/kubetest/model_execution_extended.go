package kubetest

import (
	"encoding/json"
	"time"
)

const (
	// ExecutionStatusQueued status for execution which is added for queue but not get yet by worker
	ExecutionStatusQueued = "queued"
	// ExecutionStatusPending status for execution which is taken by worker
	ExecutionStatusPending = "pending"
	// ExecutionStatusSuceess execution complete with success
	ExecutionStatusSuceess = "success"
	// ExecutionStatusSuceess execution failed
	ExecutionStatusError = "error"
)

func NewExecution(ID string, content string) Execution {
	return Execution{
		Id:            ID,
		ScriptContent: content,
		Status:        ExecutionStatusQueued,
	}
}
func (e *Execution) Start() {
	e.StartTime = time.Now()
}

func (e *Execution) Stop() {
	e.EndTime = time.Now()
}

func (e *Execution) Success() {
	e.Status = ExecutionStatusSuceess
}

func (e *Execution) Error(err error) {
	e.Status = ExecutionStatusError
	e.ErrorMessage = err.Error()
}

func (e *Execution) IsCompleted() bool {
	return e.Status == ExecutionStatusSuceess || e.Status == ExecutionStatusError
}

type ExecuteRequest struct {
	Type     string          `json:"type,omitempty"`
	Name     string          `json:"name,omitempty"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}