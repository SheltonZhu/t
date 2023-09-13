package job

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJobDone(t *testing.T) {
	t.Parallel()
	err := NewJob(
		WithJobTimeout(10 * time.Second),
	).Execute()
	assert.NoError(t, err)
}

func TestJobTimeout(t *testing.T) {
	t.Parallel()
	timeout := 500 * time.Millisecond
	err := NewJob(
		WithJobTimeout(timeout),
		WithRetryMaxTimes(2),
		WithKeepAliveEnable(),
		WithKeepAliveInterval(100*time.Millisecond),
		WithJobFunc(
			func(context.Context, *Job, uint) error {
				return assert.AnError
			}),
	).Execute()
	timeoutErr := JobTimeoutErr{timeout}
	assert.ErrorIs(t, err, timeoutErr)
	assert.Equal(t, err.Error(), timeoutErr.Error())
}

func TestKeepAliveError(t *testing.T) {
	t.Parallel()
	err := NewJob(
		WithJobTimeout(1*time.Second),
		WithRetryMaxTimes(3),
		WithKeepAliveEnable(),
		WithKeepAliveInterval(100*time.Millisecond),
		WithJobFunc(
			func(context.Context, *Job, uint) error {
				time.Sleep(500 * time.Millisecond)
				return nil
			}),
		WithKeepAliveFunc(
			func(context.Context, *Job) error {
				return assert.AnError
			}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestRetryMaxTimes(t *testing.T) {
	t.Parallel()
	timeout := 1 * time.Second
	err := NewJob(
		WithJobTimeout(timeout),
		WithRetryMaxTimes(2),
		WithJobFunc(
			func(context.Context, *Job, uint) error {
				time.Sleep(100 * time.Millisecond)
				return assert.AnError
			}),
		WithRetryIntervalFunc(func(*Job, uint, time.Duration) time.Duration {
			return 100 * time.Millisecond
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestExternalCancel(t *testing.T) {
	t.Parallel()
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	err := NewJob(WithContext(ctx), WithJobTimeout(1*time.Second)).Execute()
	assert.ErrorIs(t, err, context.Canceled)
}

func TestJobFuncPaninc(t *testing.T) {
	t.Parallel()
	err := NewJob(
		WithJobTimeout(10*time.Second),
		WithJobFunc(func(context.Context, *Job, uint) error {
			panic(assert.AnError)
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestKeepAliveFuncPaninc(t *testing.T) {
	t.Parallel()
	err := NewJob(
		WithJobTimeout(10*time.Second),
		WithJobFunc(func(context.Context, *Job, uint) error {
			time.Sleep(1 * time.Second)
			return nil
		}),
		WithKeepAliveEnable(),
		WithKeepAliveInterval(10*time.Millisecond),
		WithKeepAliveFunc(func(context.Context, *Job) error {
			panic(assert.AnError)
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestRetryIntervalFuncPanic(t *testing.T) {
	t.Parallel()
	err := NewJob(
		WithJobTimeout(10*time.Second),
		WithJobFunc(func(context.Context, *Job, uint) error {
			return assert.AnError
		}),
		WithRetryMaxTimes(4),
		WithRetryIntervalFunc(func(*Job, uint, time.Duration) time.Duration {
			panic(assert.AnError)
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestRetryWaitTimeFunc(t *testing.T) {
	t.Parallel()
	t.SkipNow()
	err := NewJob(
		WithJobTimeout(100*time.Second),
		WithJobFunc(func(context.Context, *Job, uint) error {
			return assert.AnError
		}),
		WithRetryMaxTimes(4),
		WithRetryIntervalFunc(func(j *Job, rt uint, lwd time.Duration) time.Duration {
			n := time.Duration(1<<rt) * time.Second
			fmt.Printf("重试: %d, 上次等待时间: %v, 下次等待时间: %v\n", rt, lwd, n)
			return n
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}
