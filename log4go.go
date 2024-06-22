package log4go

import (
	"bufio"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"log4go/format"
	"log4go/writer"
	"os"
)

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
