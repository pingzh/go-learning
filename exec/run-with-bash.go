package main

import (
        "io"
        "log"
        "os"
        "os/exec"
)

func main() {
        cmd := exec.Command("bash", "-c", "ps aux | wc -l")
        pipeReader, pipeWriter := io.Pipe()
        defer pipeWriter.Close()

        cmd.Stdout = pipeWriter
        cmd.Stderr = pipeWriter
        go io.Copy(os.Stdout, pipeReader)
        if err := cmd.Run(); err != nil {
                log.Fatal("Error ", err)
        }

}
