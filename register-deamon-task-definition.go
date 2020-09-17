package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
		log.Println("registerDaemonTaskDefinition")
		taskName := "RequlMobileStaging"
		newTaskName := "RequlMobileStaging_delayed_job"
		containerName := "requl_mobile"
		daemonCommandInput := "bundle exec bin/delayed_job run"
		memoryReservation := 512

		sess, err := session.NewSession()
		if err != nil {
			log.Fatalln(err.Error())
		}
		svc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

		// Get Current TaskDefinition
		taskParams := &ecs.DescribeTaskDefinitionInput{
			TaskDefinition: aws.String(taskName),
		}
		taskDef, err := svc.DescribeTaskDefinition(taskParams)
		if err != nil {
			log.Fatalln(err.Error())
		}

		// Update Command
		daemonCommand := []*string{}
		for _, cmd := range strings.Fields(daemonCommandInput) {
			cmdStr := cmd
			daemonCommand = append(daemonCommand, &cmdStr)
		}
		var newContainerDef []*ecs.ContainerDefinition
		for i := range taskDef.TaskDefinition.ContainerDefinitions {
			if *taskDef.TaskDefinition.ContainerDefinitions[i].Name == containerName {
				newContainerDef = append(newContainerDef, taskDef.TaskDefinition.ContainerDefinitions[i])
				newContainerDef[0].Command = daemonCommand
				newContainerDef[0].Links = []*string{}
			}
			if *taskDef.TaskDefinition.ContainerDefinitions[i].Name == ("envoy_" + containerName) {
				newContainerDef = append(newContainerDef, taskDef.TaskDefinition.ContainerDefinitions[i])
				// main container needs a link to envoy container
				newContainerDef[0].Links = append(newContainerDef[0].Links, taskDef.TaskDefinition.ContainerDefinitions[i].Name)
			}
		}

		for _, container := range newContainerDef {
			if *container.LogConfiguration.LogDriver == "fluentd" {
				container.LogConfiguration.Options["tag"] = aws.String(fmt.Sprintf("%s_task", *container.LogConfiguration.Options["tag"]))
			}
			if *container.LogConfiguration.LogDriver == "awslogs" {
				container.LogConfiguration.Options["awslogs-stream-prefix"] = aws.String(fmt.Sprintf("%s_task", *container.LogConfiguration.Options["awslogs-stream-prefix"]))
			}
			*container.MemoryReservation = int64(memoryReservation)
		}

		// Register TaskDefinition
		newParams := &ecs.RegisterTaskDefinitionInput{
			ContainerDefinitions:    newContainerDef,
			Cpu:                     taskDef.TaskDefinition.Cpu,
			ExecutionRoleArn:        taskDef.TaskDefinition.ExecutionRoleArn,
			Family:                  aws.String(newTaskName),
			Memory:                  taskDef.TaskDefinition.Memory,
			NetworkMode:             taskDef.TaskDefinition.NetworkMode,
			PlacementConstraints:    taskDef.TaskDefinition.PlacementConstraints,
			RequiresCompatibilities: taskDef.TaskDefinition.RequiresCompatibilities,
			TaskRoleArn:             taskDef.TaskDefinition.TaskRoleArn,
			Volumes:                 taskDef.TaskDefinition.Volumes,
		}

		log.Println(newParams)

		// newTaskDef, err := svc.RegisterTaskDefinition(newParams)
		// if err != nil {
		// 	log.Fatalln(err.Error())
		// }

		// log.Println("==============================================================================================")
		// log.Println(*newTaskDef)
		log.Fatalln("An error occurred when calling RegisterTaskDefinition API!")
	}