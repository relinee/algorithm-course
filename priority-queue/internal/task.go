package internal

import (
	"math/rand/v2"
	"time"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	priority Priority
	duration time.Duration
	id       int
}

func NewTask(priority Priority, duration time.Duration) *Task {
	return &Task{
		priority: priority,
		duration: duration,
		id:       rand.Int(),
	}
}
