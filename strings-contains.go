package main

import "strings"
import "fmt"

func handleError(errMsg string) {
	if strings.Contains(errMsg, "ExitCode: 255") && strings.Contains(errMsg, "Connection refused") {
		return
	}

	fmt.Println(errMsg)
}

func main() {
	errMsg := `[DAEMON_WATCHER]
Command [ssh ec2-user@172.32.6.173 -i /home/finc/.ssh/ecs-production.pem docker ps | grep 'FincConnectProduction_sidekiq-62' | cut -f1 -d' '] timeout, ExitCode: 255
stderr: ssh: connect to host 172.32.6.173 port 22: Connection refused`

	handleError(errMsg)
	fmt.Println("DONE!")
}