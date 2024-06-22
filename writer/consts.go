package writer

type WriterType int

const (
	File_Write  = 1
	Os_StdOut   = 2
	File_StdOut = 10
)

type LogLevel int

const (
	InfoLevel = iota + 1
	DebugLevel
	WarnLevel
	ErrorLevel
	PanicLevel
)
