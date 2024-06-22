package writer

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path"
	"time"
)

func GetLogWriter(w io.Writer, writeType WriterType) io.Writer {
	var logWriter io.Writer
	if writeType == File_Write {
		logWriter = w
	}
	if writeType == Os_StdOut {
		logWriter = os.Stdout
	}
	if writeType == File_StdOut {
		logWriter = io.MultiWriter(os.Stdout, w)
	}
	if logWriter == nil {
		logWriter = os.Stdout
	}
	return logWriter
}

func GetFileWriter(logFilePath string) (*rotatelogs.RotateLogs, *os.File) {
	if logFilePath == "" {
		logFilePath = "/data/logs/"
	}
	fileInfo, err := os.Stat(logFilePath)
	if _, ok := err.(*os.PathError); ok {
		err = os.Mkdir(logFilePath, os.ModePerm)
		fileInfo, err = os.Stat(logFilePath)
	}
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("Log4GoInit error, filePath err or filePath is not a directory")
		panic("Log4GoInit error")
	}
	filename := path.Join(logFilePath, "service.log")
	rotateFileName := filename + "_%Y%m%d"

	rotatelogWriter, err := rotatelogs.New(
		rotateFileName,
		//rotatelogs.WithLinkName(filename),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationCount(30))
	if err != nil {
		panic("rotatelogs error")
	}

	errWriter, err := os.OpenFile(path.Join(logFilePath, "errors.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic("errorlog error")
	}

	return rotatelogWriter, errWriter
}
