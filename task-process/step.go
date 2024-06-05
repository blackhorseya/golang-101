package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Task is a struct that represents a task.
type Task struct {
	Title string `json:"title,omitempty"`
	Step  Step   `json:"step,omitempty"`
}

// Execute executes the task.
func (t *Task) Execute() error {
	return t.Step.Execute(t)
}

// NewTask creates a new task with the given title.
func NewTask(title string) *Task {
	return &Task{
		Title: title,
		Step:  &StepBacklog{},
	}
}

// Step is an interface for a task step.
type Step interface {
	fmt.Stringer
	json.Marshaler

	Execute(context *Task) error
}

var _ Step = &StepBacklog{}

// StepBacklog is a concrete step for a task backlog.
type StepBacklog struct {
}

func (s *StepBacklog) String() string {
	return "backlog"
}

func (s *StepBacklog) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *StepBacklog) Execute(context *Task) error {
	context.Step = &StepTodo{}
	return nil
}

var _ Step = &StepTodo{}

// StepTodo is a concrete step for a task todo.
type StepTodo struct {
}

func (s *StepTodo) String() string {
	return "todo"
}

func (s *StepTodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *StepTodo) Execute(context *Task) error {
	context.Step = &StepInProgress{}
	return nil
}

var _ Step = &StepInProgress{}

// StepInProgress is a concrete step for a task in progress.
type StepInProgress struct {
}

func (s *StepInProgress) String() string {
	return "in_progress"
}

func (s *StepInProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *StepInProgress) Execute(context *Task) error {
	context.Step = &StepDone{}
	return nil
}

var _ Step = &StepDone{}

// StepDone is a concrete step for a task done.
type StepDone struct {
}

func (s *StepDone) String() string {
	return "done"
}

func (s *StepDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *StepDone) Execute(context *Task) error {
	return errors.New("task is already done")
}
