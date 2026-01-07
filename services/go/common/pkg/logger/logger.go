package logger

import (
	"fmt"
	"log"
	"os"
)

// New returns a basic stdout logger with a service-prefixed label.
func New(service string) *log.Logger {
	prefix := fmt.Sprintf("[%s] ", service)
	return log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile)
}
