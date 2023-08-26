package job

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJobDone(t *testing.T) {
	err := NewJob(
		WithJobTimeout(10 * time.Second),
	).Execute()
	assert.NoError(t, err)
}

func TestJobTimeout(t *testing.T) {
	timeout := 500 * time.Millisecond
	err := NewJob(
		WithJobTimeout(timeout),
		WithRetryMaxTimes(2),
		WithKeepAliveEnable(),
		WithKeepAliveInterval(100*time.Millisecond),
		WithJobFunc(
			func(ctx context.Context) error {
				return assert.AnError
			}),
	).Execute()
	timeoutErr := JobTimeoutErr{timeout}
	assert.ErrorIs(t, err, timeoutErr)
	assert.Equal(t, err.Error(), timeoutErr.Error())
}

func TestKeepAliveError(t *testing.T) {
	err := NewJob(
		WithJobTimeout(1*time.Second),
		WithRetryMaxTimes(3),
		WithKeepAliveEnable(),
		WithKeepAliveInterval(100*time.Millisecond),
		WithJobFunc(
			func(ctx context.Context) error {
				time.Sleep(500 * time.Millisecond)
				return nil
			}),
		WithKeepAliveFunc(
			func(context.Context) error {
				return assert.AnError
			}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestRetryMaxTimes(t *testing.T) {
	timeout := 1 * time.Second
	err := NewJob(
		WithJobTimeout(timeout),
		WithRetryMaxTimes(2),
		WithJobFunc(
			func(ctx context.Context) error {
				time.Sleep(100 * time.Millisecond)
				return assert.AnError
			}),
		WithRetryIntervalFunc(func(uint, time.Duration) time.Duration {
			return 100 * time.Millisecond
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestExternalCancel(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()
	err := NewJob(WithContext(ctx), WithJobTimeout(1*time.Second)).Execute()
	assert.ErrorIs(t, err, context.Canceled)
}

func TestJobFuncPaninc(t *testing.T) {
	err := NewJob(
		WithJobTimeout(10*time.Second),
		WithJobFunc(func(ctx context.Context) error {
			panic(assert.AnError)
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestKeepAliveFuncPaninc(t *testing.T) {
	err := NewJob(
		WithJobTimeout(10*time.Second),
		WithJobFunc(func(ctx context.Context) error {
			time.Sleep(1 * time.Second)
			return nil
		}),
		WithKeepAliveEnable(),
		WithKeepAliveInterval(10*time.Millisecond),
		WithKeepAliveFunc(func(ctx context.Context) error {
			panic(assert.AnError)
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestRetryIntervalFuncPanic(t *testing.T) {
	err := NewJob(
		WithJobTimeout(10*time.Second),
		WithJobFunc(func(ctx context.Context) error {
			return assert.AnError
		}),
		WithRetryMaxTimes(4),
		WithRetryIntervalFunc(func(rt uint, lwd time.Duration) time.Duration {
			panic(assert.AnError)
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}

func TestRetryWaitTimeFunc(t *testing.T) {
	t.SkipNow()
	err := NewJob(
		WithJobTimeout(100*time.Second),
		WithJobFunc(func(ctx context.Context) error {
			return assert.AnError
		}),
		WithRetryMaxTimes(4),
		WithRetryIntervalFunc(func(rt uint, lwd time.Duration) time.Duration {
			n := time.Duration(1<<rt) * time.Second
			fmt.Printf("重试: %d, 上次等待时间: %v, 下次等待时间: %v\n", rt, lwd, n)
			return n
		}),
	).Execute()
	assert.ErrorIs(t, err, assert.AnError)
}
