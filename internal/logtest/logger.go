// Package logtest provides a log.Logger for testing.
package logtest

import (
	"log"
	"testing"

	"go.abhg.dev/gs/internal/ioutil"
)

// New builds a logger that writes messages
// to the given testing.TB.
func New(t testing.TB) *log.Logger {
	return log.New(ioutil.TestLogWriter(t, ""), "", 0)
}