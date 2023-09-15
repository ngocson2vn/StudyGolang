package ecs

import (
	"os"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func ForceNewDeployment(clusterName string, serviceName string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	serviceParams := &ecs.DescribeServicesInput{
		Services: []*string{
			aws.String(serviceName),
		},
		Cluster: aws.String(clusterName),
	}
	result, err := ecsSvc.DescribeServices(serviceParams)
	if err != nil {
		return err
	}
	if len(result.Services) == 0 {
		return fmt.Errorf("Could not find service %s in cluster %s", serviceName, clusterName)
	}

	// Update Service
	for i := range result.Services {
		service := result.Services[i]
		if *service.ServiceName == serviceName {
			newServiceParams := &ecs.UpdateServiceInput{
				Service:                 service.ServiceName,
				Cluster:                 aws.String(clusterName),
				ForceNewDeployment:      aws.Bool(true),
			}
			_, err := ecsSvc.UpdateService(newServiceParams)
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func UpdateTaskDefinition(taskName string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		return err
	}

	newTaskDefInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    taskDef.TaskDefinition.ContainerDefinitions,
		Cpu:                     taskDef.TaskDefinition.Cpu,
		ExecutionRoleArn:        taskDef.TaskDefinition.ExecutionRoleArn,
		Family:                  taskDef.TaskDefinition.Family,
		Memory:                  taskDef.TaskDefinition.Memory,
		NetworkMode:             taskDef.TaskDefinition.NetworkMode,
		PlacementConstraints:    taskDef.TaskDefinition.PlacementConstraints,
		RequiresCompatibilities: taskDef.TaskDefinition.RequiresCompatibilities,
		TaskRoleArn:             taskDef.TaskDefinition.TaskRoleArn,
		Volumes:                 taskDef.TaskDefinition.Volumes,
	}

	newTaskDef, err := ecsSvc.RegisterTaskDefinition(newTaskDefInput)
	if err != nil {
		return err
	}
	log.Println(*newTaskDef.TaskDefinition.TaskDefinitionArn)

	return nil
}
