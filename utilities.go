// Package utilities common utilities used by my projects.
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

// CheckFatal checks for error that are not io.EOF and io.ErrUnexpectedEOF and logs them.
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

// DeferClose with log the given string and then execute the given Close funtion
func DeferClose(s string, rw func() error) {
	if len(s) != 0 {
		log.Println(s)
	}
	err := rw()
	CheckFatal(err)
}

// Trace will log the given string.
// It is used with Un() to log entering and leaving a function by adding
// 'defer Un(Trace("<function name>"))' at the begining of  the function.
func Trace(s string) string {
	log.Printf("entering: %s", s)
	return s
}

// Un will log the given string
func Un(s string) {
	log.Printf("leaving: %s", s)
}

// SetLogFileName will arrange the named file to be used for the log ouput.
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

// TurnOffLogging sets the logging output to os.DevNull
func TurnOffLogging() {
	var err error

	if loggingOn {
		if nullFile == nil {
			nullFile, err = os.Open(os.DevNull)
			CheckFatal(err)
		}

		log.SetOutput(nullFile)
		loggingOn = false
	}
}

// TurnOnLogging sets the logging output to the currently assigned logFile.
func TurnOnLogging() {
	if !loggingOn {
		log.SetOutput(logFile)
		loggingOn = true
	}
}

// IsLoggingOn will return true if logging is currently enabled.
func IsLoggingOn() bool {
	return loggingOn
}
