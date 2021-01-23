// Package utils is mostly focused in summarizing recurrent functions
// in this binary (and possibly other ones), to avoid repetition in
// different packages
//
package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

// Check function - General purpose error checker
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Splitter function - General purpose string splitter
// takes in a string, a delimiter, and the index (nth position)
// to be retrieved
func Splitter(s, d string, i int) (val string) {
	res := strings.Split(s, d)
	val = res[i]

	return val
}

// Run function - General purpose command-line exec function
// with bytes buffer standard out/err (so that you can
// manipulate the output)
func Run(args ...string) ([]byte, error) {

	cmdPath, err := exec.LookPath(args[0])
	if err != nil {
		return nil, err
	}

	cmdFlags := args[1:]

	c := exec.Command(cmdPath, cmdFlags...)
	var outb, errb bytes.Buffer
	c.Stdout = &outb
	c.Stderr = &errb
	e := c.Run()
	if e != nil {
		return nil, e
	}

	return outb.Bytes(), nil

}

// TrimSuffix function will remove the last part of a string,
// taking the string and suffix as input
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
