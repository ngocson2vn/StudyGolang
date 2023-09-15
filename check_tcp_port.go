package main

import (
	"fmt"
	"net"
	"time"
)

func raw_connect(host string, port string) error {
	timeout := 3 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Opened", net.JoinHostPort(host, port))
	}
	return nil
}

func main() {
	err := raw_connect("127.0.0.1", "22")
	if err != nil {
		fmt.Println(err)
	}
	err = raw_connect("127.0.0.1", "80")
	if err != nil {
		fmt.Println(err)
	}

	// python -m http.server 8080
	err = raw_connect("127.0.0.1", "8080")
	if err != nil {
		fmt.Println(err)
	}
}
