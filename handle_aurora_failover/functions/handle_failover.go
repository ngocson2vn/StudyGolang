package main

import (
	"log"
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/FiNCDeveloper/finc_server/handle_aurora_failover/config"
	"github.com/FiNCDeveloper/finc_server/handle_aurora_failover/libs/ecs"

	. "github.com/FiNCDeveloper/finc_server/handle_aurora_failover/libs/common"
)

func findAuroraCluster(auroraClusters []AuroraCluster, clusterName string) AuroraCluster {
	for _, auroraCluster := range auroraClusters {
		if auroraCluster.ClusterName == clusterName {
			return auroraCluster
		}
	}

	return AuroraCluster{}
}

func updateEcsServices(clusterName string, dependentServices []string) {
	for _, serviceName := range dependentServices {
		err := ecs.ForceNewDeployment(clusterName, serviceName)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Printf("Force new deployment: %s\n", serviceName)
		}
	}
}

func updateTaskDefinitions(dependentWorkers []string) {
	for _, taskName := range dependentWorkers {
		err := ecs.UpdateTaskDefinition(taskName)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Printf("Update Task Definition: %s\n", taskName)
		}
	}
}

func handler(ctx context.Context, snsEvent events.SNSEvent) (string, error) {
	ecsClusterName := os.Getenv("ECS_CLUSTER_NAME")
	log.Printf("ECS_CLUSTER_NAME: %s\n", ecsClusterName)

	// For DEBUG
	if len(snsEvent.Records) == 0 {
		out, _ := json.Marshal(config.AuroraClusters)
		log.Print(string(out))

		out, _ = json.MarshalIndent(config.AuroraClusters, "", "  ")
		return string(out), nil
	}

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		snsMessage := SNSMessage{}
		err := json.Unmarshal([]byte(snsRecord.Message), &snsMessage)
		if err != nil {
			return err.Error(), err
		}

		if strings.Contains(snsMessage.EventMessage, "Completed failover") {
			log.Printf("%s, %s, %s\n", snsMessage.EventSource, snsMessage.SourceID, snsMessage.EventMessage)
			auroraCluster := findAuroraCluster(config.AuroraClusters, snsMessage.SourceID)

			// 1. Update ECS services: Force new deployment
			updateEcsServices(ecsClusterName, auroraCluster.DependentServices)

			// 2. Update Task Definitions for asynchronous workers.
			// The daemon_watcher will gracefully restart all workers.
			updateTaskDefinitions(auroraCluster.DependentWorkers)
		}
	}

	return STATUS_SUCCESS, nil
}

func main() {
	lambda.Start(handler)
}
