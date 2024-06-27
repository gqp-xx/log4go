package log4go

import (
	"bufio"
	"github.com/gqp-xx/log4go/format"
	"github.com/gqp-xx/log4go/writer"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

/**
 * 日志组件初始化
 * logFilePath 日志文件路径
 * writeType 输出类型(日志文件 标准输出 文件和标准输出）
 * level 日志级别
 */
func Log4GoInit(logFilePath string, writeType writer.WriterType, level writer.LogLevel) {
	rotatelogWriter, errWriter := writer.GetFileWriter(logFilePath)
	logWriter := writer.GetLogWriter(rotatelogWriter, writeType)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: io.MultiWriter(errWriter, os.Stdout),
		logrus.TraceLevel: logWriter,
		logrus.WarnLevel:  logWriter,
	}, &format.LogFormat{})
	logrus.AddHook(lfHook)

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&format.LogFormat{})
	logrus.SetLevel(getLogRusLever(level))

	// 丢弃logrus输出
	src, _ := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logrus.SetOutput(bufio.NewWriter(src))
}

func getLogRusLever(level writer.LogLevel) logrus.Level {
	if level == 0 {
		return logrus.InfoLevel
	}
	switch level {
	case writer.PanicLevel:
		return logrus.PanicLevel
	case writer.ErrorLevel:
		return logrus.ErrorLevel
	case writer.WarnLevel:
		return logrus.WarnLevel
	case writer.InfoLevel:
		return logrus.InfoLevel
	case writer.DebugLevel:
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}

/**
 * 根据运行环境初始化日志组件
 * profile 运行环境 dev开发环境，debug,输出为(标准输出)控制台+文件 test为测试环境，debug,输出到文件
 * logFilePath 日志文件路径，默认路径为/data/logs/
 */
func InitWithProfile(profile string, logFilePath string) {
	var logLevel writer.LogLevel = writer.InfoLevel
	var writeType writer.WriterType = writer.File_Write
	if profile == "dev" {
		logLevel = writer.DebugLevel
		writeType = writer.File_StdOut
	}
	if profile == "test" {
		logLevel = writer.DebugLevel
		writeType = writer.File_Write
	}
	Log4GoInit(logFilePath, writeType, logLevel)
}
