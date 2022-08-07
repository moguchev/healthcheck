package main

import (
	"bytes"
	"strings"
	"testing"
)

// printer_test.go
func TestPrint(t *testing.T) {
	const (
		input = "Hello, World!"
		want  = "hello, world!"
	)

	var buf bytes.Buffer
	LowerPrintV2(&buf, input)

	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Errorf("Expected output to be: %s, but got: %s", want, got)
	}
}

