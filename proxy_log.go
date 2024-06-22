package log4go

import (
	"context"
	"github.com/sirupsen/logrus"
)

var (
	// std is the name of the default logger log4go
	std = newLog4Go()
)

type ProxyLog struct {
	*logrus.Logger
	LogName string
}

func newLog4Go() *ProxyLog {
	return &ProxyLog{
		Logger: logrus.StandardLogger(),
	}
}

func NewLog4Go(logName string) *ProxyLog {
	return &ProxyLog{
		Logger:  logrus.StandardLogger(),
		LogName: logName,
	}
}

func Info(ctx context.Context, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Info(args...)
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Infof(format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Debug(args...)
}
func Debugf(ctx context.Context, format string, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Debugf(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Warn(args...)
}
func Warnf(ctx context.Context, format string, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Warnf(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Error(args...)
}
func Errorf(ctx context.Context, format string, args ...interface{}) {
	std.WithField("trace_id", ctx.Value("trace_id")).Errorf(format, args...)
}

func (proxyLog *ProxyLog) Info(ctx context.Context, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Info(args...)
}
func (proxyLog *ProxyLog) Infof(ctx context.Context, format string, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Infof(format, args...)
}

func (proxyLog *ProxyLog) Debug(ctx context.Context, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Debug(args...)
}
func (proxyLog *ProxyLog) Debugf(ctx context.Context, format string, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Debugf(format, args...)
}

func (proxyLog *ProxyLog) Warn(ctx context.Context, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Warn(args...)
}
func (proxyLog *ProxyLog) Warnf(ctx context.Context, format string, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Warnf(format, args...)
}

func (proxyLog *ProxyLog) Error(ctx context.Context, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Error(args...)
}
func (proxyLog *ProxyLog) Errorf(ctx context.Context, format string, args ...interface{}) {
	proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Errorf(format, args...)
}
