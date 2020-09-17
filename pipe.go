package main

import (
        "bytes"
        "io"
        "os"
        "os/exec"
)

func main() {
        //create command
        catCmd := exec.Command("ls")
        wcCmd := exec.Command("wc", "-l")

        //make a pipe
        reader, writer := io.Pipe()
        var buf bytes.Buffer

        //set the output of "cat" command to pipe writer
        catCmd.Stdout = writer
        //set the input of the "wc" command pipe reader

        wcCmd.Stdin = reader

        //cache the output of "wc" to memory
        wcCmd.Stdout = &buf

        //start to execute "cat" command
        catCmd.Start()

        //start to execute "wc" command
        wcCmd.Start()

        //waiting for "cat" command complete and close the writer
        catCmd.Wait()
        writer.Close()

        //waiting for the "wc" command complete and close the reader
        wcCmd.Wait()
        reader.Close()
        //copy the buf to the standard output
        io.Copy( os.Stdout, &buf )
}
