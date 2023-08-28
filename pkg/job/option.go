package job

import (
	"context"
	"time"
)

type JobOption func(*job)

// WithContext 设置context
func WithContext(ctx context.Context) JobOption {
	return func(j *job) {
		j.ctx = ctx
	}
}

// WithJobTimeout 设置job超时时间
func WithJobTimeout(timeout time.Duration) JobOption {
	return func(j *job) {
		j.jobTimeout = timeout
	}
}

// WithJobFunc 设置job函数
func WithJobFunc(jobFunc func(context.Context, *job, uint) error) JobOption {
	return func(j *job) {
		j.jobFunc = jobFunc
	}
}

// WithKeepAliveInterval 设置keep alive间隔
func WithKeepAliveInterval(interval time.Duration) JobOption {
	return func(j *job) {
		j.keepAliveInterval = interval
	}
}

// WithKeepAliveEnable 设置keep alive开关
func WithKeepAliveEnable() JobOption {
	return func(j *job) {
		j.keepAliveEnable = true
	}
}

// WithKeepAliveFunc 设置keep alive函数
func WithKeepAliveFunc(keepAliveFunc func(context.Context, *job) error) JobOption {
	return func(j *job) {
		j.keepAliveFunc = keepAliveFunc
	}
}

// WithRetryMaxTimes 设置最大重试次数
func WithRetryMaxTimes(retryMaxTimes uint) JobOption {
	return func(j *job) {
		j.retryMaxTimes = retryMaxTimes
	}
}

// WithRetryIntervalFunc 设置重试间隔函数
func WithRetryIntervalFunc(retryIntervalFunc func(*job, uint, time.Duration) time.Duration) JobOption {
	return func(j *job) {
		j.retryIntervalFunc = retryIntervalFunc
	}
}
