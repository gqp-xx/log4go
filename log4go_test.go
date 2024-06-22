package log4go

import (
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
