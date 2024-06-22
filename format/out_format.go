package format

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

type LogFormat struct {
}

func (l LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	logName, ok := entry.Data["log_name"]
	if ok {
		logName = logName.(string)
	} else {
		logName = "-"
	}
	traceId, ok := entry.Data["trace_id"]
	if ok {
		traceId = traceId.(string)
	} else {
		traceId = "-"
	}
	newLog := fmt.Sprintf("[%s] [%s] [%s] [%s] %s\n", timestamp, entry.Level, logName, traceId, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}
