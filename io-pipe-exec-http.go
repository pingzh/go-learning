package main

import (
	"io"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("bash", "-c", "ls -alh; lsl")
	pipeReader, pipeWriter := io.Pipe()
	defer pipeWriter.Close()

	cmd.Stdout = pipeWriter
	cmd.Stderr = pipeWriter
	go io.Copy(w, pipeReader)
	cmd.Run()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
