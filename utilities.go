// utilities project utilities.go
package utilities

import (
	"io"
	"log"
	"os"
)

var (
	nullFile    *os.File
	logFile     = os.Stderr
	loggingOn   = false
	logFileName string
)

func CheckFatal(e error) {
	if e != nil && e != io.EOF && e != io.ErrUnexpectedEOF {
		if loggingOn {
			log.Fatal(e)
		} else {
			TurnOnLogging()
			log.Fatal(e)
			TurnOffLogging()
		}
	}
}

func DeferClose(s string, rw func() error) {
	if len(s) != 0 {
		log.Println(s)
	}
	err := rw()
	CheckFatal(err)
}

func Trace(s string) string {
	log.Printf("entering: %s", s)
	return s
}

func Un(s string) {
	log.Printf("leaving: %s", s)
}

func SetLogFileName(fileName string) {
	var err error
	logFileName = fileName
	if len(logFileName) != 0 {
		logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.SetOutput(os.Stderr)
			log.Fatal(err)
		}
	}
}

func TurnOffLogging() {
	var err error

	if nullFile == nil {
		nullFile, err = os.Open(os.DevNull)
		CheckFatal(err)
	}

	log.SetOutput(nullFile)
	loggingOn = false
}

func TurnOnLogging() {
	// Set the output to os.Stderr so that if opening the log file fails, the
	// error will be sent to Stderr.
	log.SetOutput(logFile)
	loggingOn = true
}

func IsLoggingOn() bool {
	return loggingOn
}
