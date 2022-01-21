package log

import (
	"io"
	"log"
	"os"
)

var (
	INFO  *log.Logger
	DEBUG *log.Logger
	ERROR *log.Logger
)

func init() {
	os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	INFO = log.New(os.Stdout, "INFO-", log.Ldate|log.Ltime|log.Lshortfile)
	DEBUG = log.New(os.Stdout, "DEBUG-", log.Ldate|log.Ltime|log.Lshortfile)
	DEBUG = log.New(io.MultiWriter(os.Stderr), "ERROR-", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(v ...interface{}) {
	INFO.Fatal(v)
}

func Debug(v ...interface{}) {
	DEBUG.Fatal(v)
}

func Error(v ...interface{}) {
	ERROR.Fatal(v)
}
