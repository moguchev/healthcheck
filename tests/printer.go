package main

import (
	"fmt"
	"io"
	"strings"
)

// printer.go
//
// How will you write tests?
func LowerPrint(text string) {
	fmt.Println(strings.ToLower(text))
}

func LowerPrintV2(w io.Writer, text string) {
	fmt.Fprintln(w, strings.ToLower(text))
}


