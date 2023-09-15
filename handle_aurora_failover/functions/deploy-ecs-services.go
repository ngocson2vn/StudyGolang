package main

import (
	"log"

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

func main() {
	ecsClusterName := "Production"
	auroraClusterName := "advice-engine-production-aurora-cluster"

	auroraCluster := findAuroraCluster(config.AuroraClusters, auroraClusterName)

	// 1. Update ECS services: Force new deployment
	updateEcsServices(ecsClusterName, auroraCluster.DependentServices)

	// 2. Update Task Definitions for asynchronous workers.
	// The daemon_watcher will gracefully restart all workers.
	updateTaskDefinitions(auroraCluster.DependentWorkers)
}
