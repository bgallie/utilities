// utilities project utilities.go
package utilities

import (
	"io"
	"log"
	"os"
)

var nullFile *os.File

func CheckFatal(e error) {
	if e != nil && e != io.EOF {
		log.Fatal(e)
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

func TurnOffLogging() {
	var err error

	if nullFile == nil {
		nullFile, err = os.Open(os.DevNull)
		CheckFatal(err)
	}

	log.SetOutput(nullFile)
}

func TurnOnLogging() {
	log.SetOutput(os.Stderr)
}
