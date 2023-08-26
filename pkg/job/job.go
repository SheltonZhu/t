package job

import (
	"context"
	"time"
)

type Job interface {
	Execute() error
}

// NewJob 创建一个job
func NewJob(opts ...JobOption) Job {
	j := newJob()
	for _, opt := range opts {
		opt(j)
	}
	return j
}

// RetryIntervalFunc 重试间隔函数
type RetryIntervalFunc func(retryTimes uint, lastWaitDuration time.Duration) time.Duration

type job struct {
	ctx               context.Context
	jobTimeout        time.Duration
	jobFunc           func(context.Context) error
	keepAliveEnable   bool
	keepAliveInterval time.Duration
	keepAliveFunc     func(context.Context) error
	retryMaxTimes     uint
	retryIntervalFunc RetryIntervalFunc
}

func newJob() *job {
	return &job{
		ctx: context.Background(),
		jobFunc: func(ctx context.Context) error {
			// do some job here
			return nil
		},
		jobTimeout:        10 * time.Minute,
		keepAliveInterval: 3 * time.Second,
		keepAliveFunc: func(ctx context.Context) error {
			// do something to keep alive
			return nil
		},
		retryMaxTimes: 0,
		retryIntervalFunc: func(rt uint, lwd time.Duration) time.Duration {
			return time.Duration(1<<rt) * time.Second
		},
	}
}

// Execute 执行job
func (j *job) Execute() error {
	timeoutErr := JobTimeoutErr{timeout: j.jobTimeout}
	ctx, cancelFunc := context.WithTimeoutCause(j.ctx, j.jobTimeout, timeoutErr)
	defer cancelFunc()

	doneChan := make(chan error)

	go j.runJob(ctx, doneChan)
	go j.keepAlive(ctx, doneChan)

	select {
	case <-ctx.Done(): // 如果上下文被取消，则返回
		if err := ctx.Err(); context.Cause(ctx) != timeoutErr {
			return err
		}
		return JobTimeoutErr{timeout: j.jobTimeout} // 超熔断返回超时错误
	case err := <-doneChan: // 任务完成返回任务执行结果
		return err
	}
}

func (j *job) runJob(ctx context.Context, doneChan chan<- error) {
	var err error
	defer func() {
		if err := recover(); err != nil {
			doneChan <- err.(error)
		}
		doneChan <- err
	}()

	// 如果不需要重试，直接执行
	if j.retryMaxTimes == 0 {
		err = j.jobFunc(ctx)
		return
	}

	var nextDelay = time.Duration(0)
	// 出错等待一阵子重试
	for i := uint(0); i < j.retryMaxTimes; i++ {
		err = j.jobFunc(ctx)
		if err == nil {
			return
		}

		// 根据指数退避重试策略计算下一个重试间隔时间
		nextDelay = j.retryIntervalFunc(i, nextDelay)
		select {
		case <-time.After(nextDelay):
			// 等待重试间隔时间后继续重试
		case <-ctx.Done():
			// 如果上下文被取消，则放弃重试
			return
		}
	}

	return
}

func (j *job) keepAlive(ctx context.Context, keepAliveChan chan<- error) {
	defer func() {
		if err := recover(); err != nil {
			keepAliveChan <- err.(error)
		}
	}()
	if !j.keepAliveEnable {
		return
	}

	ticker := time.NewTicker(j.keepAliveInterval)
	for {
		select {
		case <-ctx.Done():
			// 如果上下文被取消，则返回
			return
		case <-ticker.C:
			if err := j.keepAliveFunc(ctx); err != nil {
				keepAliveChan <- err
			}
		}
	}
}
