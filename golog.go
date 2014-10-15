// package golog implements logging functions that log errors to stderr and
// debug messages to stdout. Trace logging is also supported. Trace logs go to
// stdout as well, but they are only written if the program is built with trace
// enabled, i.e.
//
//   -tags trace
//
package golog

import (
	"fmt"
	"os"
)

// Debug logs to stdout
func Debug(arg interface{}) {
	fmt.Fprintln(os.Stdout, arg)
}

// Debugf logs to stdout
func Debugf(message string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, message+"\n", args...)
}

// Error logs to stderr
func Error(arg interface{}) {
	fmt.Fprintln(os.Stderr, arg)
}

// Errorf logs to stderr
func Errorf(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message+"\n", args...)
}

// Fatal logs to stderr and then exits with status 1
func Fatal(arg interface{}) {
	Error(arg)
	os.Exit(1)
}

// Fatalf logs to stderr and then exits with status 1
func Fatalf(message string, args ...interface{}) {
	Errorf(message, args...)
	os.Exit(1)
}
