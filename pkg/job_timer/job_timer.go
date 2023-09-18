package job_timer

import (
	"time"
)

type Logger interface {
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type JobTimer struct {
	Logger  Logger
	start   *time.Time
	elapsed time.Duration
}

func New() *JobTimer {
	return &JobTimer{}
}

func (t *JobTimer) Stop() {
	if t.IsRunning() {
		t.elapsed += time.Since(*t.start)
		t.start = nil
	}
}

func (t *JobTimer) Start() {
	if t.IsRunning() {
		return
	}
	now := time.Now()
	t.start = &now
}

func (t *JobTimer) Reset() {
	t.elapsed = 0
	t.start = nil
}

func (t *JobTimer) Cost() time.Duration {
	return t.elapsed
}

func (t *JobTimer) Count() time.Duration {
	if t.IsRunning() {
		defer func() {
			now := time.Now()
			t.start = &now
		}()
		count := time.Since(*t.start)
		t.elapsed += count
		return count
	}
	return 0
}

func (t *JobTimer) IsRunning() bool {
	return t.start != nil
}
