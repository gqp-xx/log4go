package log4go

import (
	"context"
	"github.com/sirupsen/logrus"
	"log4go/writer"
	"testing"
)

func TestLog4GoInit(t *testing.T) {
	Log4GoInit("/Users/Zhuanz/GolandProjects/log4go/log", writer.Os_StdOut, writer.DebugLevel)
	logrus.Info("******* test log4go init1 ******")
	logrus.Warn("******* test log4go init2 ******")
	logrus.Debug("******* test log4go init3 ******")
	logrus.Error("******* test log4go init4 ******")
}

func TestLog4Go(t *testing.T) {
	Log4GoInit("./log", writer.File_StdOut, writer.DebugLevel)

	ctx := context.WithValue(context.Background(), "trace_id", "12345678")
	Info(ctx, "test info")

	log := NewLog4Go("daluandou_main")
	log.Infof(ctx, "test info")
}
