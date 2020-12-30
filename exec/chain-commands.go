package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	first := exec.Command("ps", "-ef")
	second := exec.Command("wc", "-l")

	// http://golang.org/pkg/io/#Pipe

	reader, writer := io.Pipe()

	// push first command output to writer
	first.Stdout = writer

	// read from first command output
	second.Stdin = reader

	// prepare a buffer to capture the output
	// after second command finished executing
	var buff bytes.Buffer
	second.Stdout = &buff

	first.Start()
	second.Start()
	first.Wait()
	writer.Close()
	second.Wait()

	total := buff.String() // convert output to string

	fmt.Printf("Total processes running : %s", total)

}
