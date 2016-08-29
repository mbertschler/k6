package lib

import (
	"context"
	"time"
)

type Status struct {
	StartTime time.Time `json:"startTime" yaml:"startTime"`

	Running bool  `json:"running" yaml:"running"`
	VUs     int64 `json:"vus" yaml:"vus"`
}

type Engine struct {
	Runner Runner
	Status Status
}

func (e *Engine) Run(ctx context.Context) error {
	e.Status.StartTime = time.Now()
	e.Status.Running = true
	defer func() {
		e.Status.Running = false
		e.Status.VUs = 0
	}()

	<-ctx.Done()
	time.Sleep(1 * time.Second)

	return nil
}

func (e *Engine) Scale(vus int64) error {
	e.Status.VUs = vus
	return nil
}