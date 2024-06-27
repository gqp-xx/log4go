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
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Info(args...)
	} else {
		std.Logger.Info(args)
	}
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Infof(format, args...)
	} else {
		std.Logger.Infof(format, args...)
	}
}

func Debug(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Debug(args...)
	} else {
		std.Logger.Debug(args...)
	}
}
func Debugf(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Debugf(format, args...)
	} else {
		std.Logger.Debugf(format, args...)
	}
}

func Warn(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Warn(args...)
	} else {
		std.Logger.Warn(args...)
	}

}
func Warnf(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Warnf(format, args...)
	} else {
		std.Logger.Warnf(format, args...)
	}
}

func Error(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Error(args...)
	} else {
		std.Logger.Error(args...)
	}

}
func Errorf(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		std.WithField("trace_id", ctx.Value("trace_id")).Errorf(format, args...)
	} else {
		std.Logger.Errorf(format, args...)
	}
}

func (proxyLog *ProxyLog) Info(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Info(args...)
	} else {
		proxyLog.Logger.Info(args)
	}
}
func (proxyLog *ProxyLog) Infof(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Infof(format, args...)
	} else {
		proxyLog.Logger.Infof(format, args...)
	}
}

func (proxyLog *ProxyLog) Debug(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Debug(args...)
	} else {
		proxyLog.Logger.Debug(args...)
	}
}
func (proxyLog *ProxyLog) Debugf(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Debugf(format, args...)
	} else {
		proxyLog.Logger.Debugf(format, args...)
	}
}

func (proxyLog *ProxyLog) Warn(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Warn(args...)
	} else {
		proxyLog.Logger.Warn(args)
	}
}
func (proxyLog *ProxyLog) Warnf(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Warnf(format, args...)
	} else {
		proxyLog.Logger.Warnf(format, args...)
	}
}

func (proxyLog *ProxyLog) Error(ctx context.Context, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Error(args...)
	} else {
		proxyLog.Logger.Error(args...)
	}
}
func (proxyLog *ProxyLog) Errorf(ctx context.Context, format string, args ...interface{}) {
	if ctx != nil {
		proxyLog.WithField("log_name", proxyLog.LogName).WithField("trace_id", ctx.Value("trace_id")).Errorf(format, args...)
	} else {
		proxyLog.Logger.Errorf(format, args...)
	}
}
