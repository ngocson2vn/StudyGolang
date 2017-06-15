package main

import "fmt"
import "strings"
import "os"

func ExecShoryukenByTask(cluster string) {
	label := fmt.Sprintf("%s_shoryuken", os.Getenv("ECS_TASK_NAME"))
	command := os.Getenv("SHORYUKEN_COMMAND")

	if command == "" {
		fmt.Println("There is no shoryuken command in Jenkinsfile.")
		return
	}

	for _, cmd := range strings.Split(command, ", ") {
		if strings.Contains(cmd, "step") {
			label = fmt.Sprintf("%s_%s_shoryuken", os.Getenv("ECS_TASK_NAME"), "sub")
		}

		ExecuteBatchTask(cmd, label, cluster, 1, 256)
		fmt.Println()
	}
}

func ExecuteBatchTask(command string, label string, cluster string, num int64, memory int64) {
	fmt.Println(command)
	fmt.Println(label, cluster, num, memory)
}

func main() {
	os.Setenv("SHORYUKEN_COMMAND", "bundle exec shoryuken -r /home/ubuntu/walking_events/lib/default_sqs_worker.rb, bundle exec shoryuken -r /home/ubuntu/walking_events/lib/step_sqs_worker.rb")
	os.Setenv("ECS_TASK_NAME", "WalkingEvents")
	cluster := "ProductionTask"
	ExecShoryukenByTask(cluster)
}
