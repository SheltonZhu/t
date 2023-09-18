package executor

import (
	"context"
	"time"
)

type Option func(*JobExecutor)

// WithContext 设置context
func WithContext(ctx context.Context) Option {
	return func(j *JobExecutor) {
		j.ctx = ctx
	}
}

// WithJobTimeout 设置job超时时间
func WithJobTimeout(timeout time.Duration) Option {
	return func(j *JobExecutor) {
		j.jobTimeout = timeout
	}
}

// WithJobFunc 设置job函数
func WithJobFunc(jobFunc func(context.Context, *JobExecutor, uint) error) Option {
	return func(j *JobExecutor) {
		j.jobFunc = jobFunc
	}
}

// WithKeepAliveInterval 设置keep alive间隔
func WithKeepAliveInterval(interval time.Duration) Option {
	return func(j *JobExecutor) {
		j.keepAliveInterval = interval
	}
}

// WithKeepAliveEnable 设置keep alive开关
func WithKeepAliveEnable() Option {
	return func(j *JobExecutor) {
		j.keepAliveEnable = true
	}
}

// WithKeepAliveFunc 设置keep alive函数
func WithKeepAliveFunc(keepAliveFunc func(context.Context, *JobExecutor) error) Option {
	return func(j *JobExecutor) {
		j.keepAliveFunc = keepAliveFunc
	}
}

// WithRetryMaxTimes 设置最大重试次数
func WithRetryMaxTimes(retryMaxTimes uint) Option {
	return func(j *JobExecutor) {
		j.retryMaxTimes = retryMaxTimes
	}
}

// WithRetryIntervalFunc 设置重试间隔函数
func WithRetryIntervalFunc(retryIntervalFunc func(*JobExecutor, uint, time.Duration) time.Duration) Option {
	return func(j *JobExecutor) {
		j.retryIntervalFunc = retryIntervalFunc
	}
}