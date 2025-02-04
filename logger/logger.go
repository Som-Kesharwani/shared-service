package logger

import (
	"io"
	"log"
	"os"
)

const (
	//UNSPECIFIED logs nothing
	UNSPECIFIED Level = iota // 0 :
	//Trace logs everything
	TRACE //1
	//INFO logs debug messages
	INFO //2
	//WARNING logs warnings
	WARNING //3
	//ERROR logs errors
	ERROR //4
)

type Level int

var (
	Trace   *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
)

func initLog(traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	isFlag bool) {
	flag := 0

	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}

	// Create log.Logger objects.
	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "ERROR: ", flag)

}

func SetLogLevel(level Level, filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error log file : %s", err.Error())
	}

	switch level {
	case TRACE:
		initLog(f, f, f, f, true)
		return
	case INFO:
		initLog(io.Discard, f, f, f, true)
		return
	case WARNING:
		initLog(io.Discard, io.Discard, f, f, true)
		return
	case ERROR:
		initLog(io.Discard, io.Discard, io.Discard, f, true)
		return
	default:
		initLog(io.Discard, f, f, f, true)
		return

	}
}
