package main

import (
    "log"
)

func main() {
    log.Println("Hello logger")

    log.Fatalln("An unknown error occurred!")
}
