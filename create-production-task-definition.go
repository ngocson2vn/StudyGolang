package main

import (
	"fmt"
	"log"
	"os"
	"time"
	// "strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func describeTaskDefinition(serviceName string) string {
	sess, err := session.NewSession()
	if err != nil {
		log.Println(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(serviceName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		fmt.Println(serviceName)
		log.Println(err.Error())
		return ""
	} else {
		return *taskDef.TaskDefinition.TaskDefinitionArn
	}

	// log.Println(*taskDef.TaskDefinition.TaskDefinitionArn)
	// log.Println(*taskDef.TaskDefinition.TaskRoleArn)
}

func updateTaskDefinition(serviceName string) error {
	taskRoleArn := fmt.Sprintf("arn:aws:iam::759549166074:role/ECSTaskRole%s", serviceName)
	log.Println(fmt.Sprintf("updateTaskDefinition: %s", serviceName))
	log.Println(taskRoleArn)

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(serviceName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(*taskDef.TaskDefinition.TaskDefinitionArn)
	log.Println(*taskDef.TaskDefinition.TaskRoleArn)

	newTaskDefInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    taskDef.TaskDefinition.ContainerDefinitions,
		Cpu:                     taskDef.TaskDefinition.Cpu,
		ExecutionRoleArn:        taskDef.TaskDefinition.ExecutionRoleArn,
		Family:                  taskDef.TaskDefinition.Family,
		Memory:                  taskDef.TaskDefinition.Memory,
		NetworkMode:             taskDef.TaskDefinition.NetworkMode,
		PlacementConstraints:    taskDef.TaskDefinition.PlacementConstraints,
		RequiresCompatibilities: taskDef.TaskDefinition.RequiresCompatibilities,
		TaskRoleArn:             aws.String(taskRoleArn),
		Volumes:                 taskDef.TaskDefinition.Volumes,
	}
	newTaskDef, err := ecsSvc.RegisterTaskDefinition(newTaskDefInput)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(*newTaskDef.TaskDefinition.TaskDefinitionArn)

	return nil
}

func updateEnvVarsForTaskDefinition(serviceName string, envMap map[string]string) error {
	log.Println(fmt.Sprintf("updateTaskDefinition: %s", serviceName))

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(serviceName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, container := range taskDef.TaskDefinition.ContainerDefinitions {
		for envName, envValue := range envMap {
				for _, env := range container.Environment {
					if *env.Name == envName {
						*env.Value = envValue
						break
					}
				}
		}

		// fmt.Println(container.Environment)
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
		log.Fatalln(err.Error())
	}
	log.Println(*newTaskDef.TaskDefinition.TaskDefinitionArn)

	return nil
}


func removeEnvVarsFromTaskDefinition(serviceName string, containerName string, envVars []string) error {
	log.Println(fmt.Sprintf("updateTaskDefinition: %s", serviceName))

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(serviceName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, container := range taskDef.TaskDefinition.ContainerDefinitions {
		if *container.Name == containerName {
			fmt.Println(container.Environment)
			fmt.Println("=========================================================")
			
			for _, envName := range envVars {
				for i, env := range container.Environment {
					if *env.Name == envName {
						if (i + 1) < len(container.Environment) {
							container.Environment = append(container.Environment[:i], container.Environment[(i + 1):]...)
						} else {
							container.Environment = append(container.Environment[:i])
						}

						break
					}
				}
			}

			fmt.Println(container.Environment)
		}
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
		log.Fatalln(err.Error())
	}
	log.Println(*newTaskDef.TaskDefinition.TaskDefinitionArn)
	fmt.Println("")

	return nil
}



func removeEnvVarsFromAllContainersInTaskDefinition(serviceName string, envVars []string) error {
	log.Println(fmt.Sprintf("updateTaskDefinition: %s", serviceName))

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(serviceName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, container := range taskDef.TaskDefinition.ContainerDefinitions {
		fmt.Println(*container.Name, len(container.Environment))
		for _, envName := range envVars {
			for i, env := range container.Environment {
				if *env.Name == envName {
					if (i + 1) < len(container.Environment) {
						container.Environment = append(container.Environment[:i], container.Environment[(i + 1):]...)
					} else {
						container.Environment = append(container.Environment[:i])
					}

					fmt.Println(*container.Name, envName)

					break
				}
			}
		}

		fmt.Println(*container.Name, len(container.Environment))
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
		log.Fatalln(err.Error())
	}
	log.Println(*newTaskDef.TaskDefinition.TaskDefinitionArn)
	fmt.Println("")

	return nil
}


func updateSecrets(serviceName string, containerName string, secretMap map[string]string) error {
	log.Println(fmt.Sprintf("updateTaskDefinition: %s", serviceName))

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(serviceName),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	found := false
	for _, container := range taskDef.TaskDefinition.ContainerDefinitions {
		if *container.Name == containerName {
			found = true

			fmt.Println(*container.Name, len(container.Environment))
			for secretName, _ := range secretMap {
				for i, env := range container.Environment {
					if *env.Name == secretName {
						if (i + 1) < len(container.Environment) {
							container.Environment = append(container.Environment[:i], container.Environment[(i + 1):]...)
						} else {
							container.Environment = append(container.Environment[:i])
						}

						fmt.Println(*container.Name, secretName)

						break
					}
				}
			}

			fmt.Println(*container.Name, len(container.Environment))

			for k, v := range secretMap {
				secret := &ecs.Secret{
					Name: aws.String(k),
					ValueFrom: aws.String(v),
				}

				container.Secrets = append(container.Secrets, secret)
			}
		}
	}

	if !found {
		log.Printf("!!! Could not found %s in %s!\n", containerName, serviceName)
	}

	executionRoleArn := fmt.Sprintf("arn:aws:iam::759549166074:role/ECSTaskExecutionRole%s", serviceName)

	newTaskDefInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    taskDef.TaskDefinition.ContainerDefinitions,
		Cpu:                     taskDef.TaskDefinition.Cpu,
		ExecutionRoleArn:        aws.String(executionRoleArn),
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
		log.Fatalln(err.Error())
	}
	log.Println(*newTaskDef.TaskDefinition.TaskDefinitionArn)

	return nil
}


func getCurrentTaskDefinition(serviceName string, clusterName string) string {
	sess, err := session.NewSession()
	if err != nil {
		log.Println(err.Error())
	}
	svc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	serviceParams := &ecs.DescribeServicesInput{
		Services: []*string{
			aws.String(serviceName),
		},
		Cluster: aws.String(clusterName),
	}
	output, err := svc.DescribeServices(serviceParams)
	if err != nil {
		fmt.Println(serviceName)
		log.Println(err.Error())
	}

	if len(output.Services) > 0 {
		return *output.Services[0].TaskDefinition
	}

	return ""
}


func updateECSService(serviceName string, clusterName string, taskDefinitionArn string) {
	log.Println("updateService")
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	svc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	serviceParams := &ecs.DescribeServicesInput{
		Services: []*string{
			aws.String(serviceName),
		},
		Cluster: aws.String(clusterName),
	}
	services, err := svc.DescribeServices(serviceParams)
	if err != nil {
		log.Println(err.Error())
	}

	// Update Service
	for i := range services.Services {
		service := services.Services[i]
		if *service.ServiceName == serviceName {
			newServiceParams := &ecs.UpdateServiceInput{
				Service:                 service.ServiceName,
				Cluster:                 aws.String(clusterName),
				DeploymentConfiguration: service.DeploymentConfiguration,
				DesiredCount:            service.DesiredCount,
				TaskDefinition:          aws.String(taskDefinitionArn),
			}
			output, err := svc.UpdateService(newServiceParams)
			if err != nil {
				log.Println(err.Error())
			}

			log.Println(*output.Service.TaskDefinition)
		}
	}
}


// func rotateSMTP() {
// 	ecsServices := []string{
// 		"AmbassadorRegistrationFormProduction",
// 		"CompanyAccountManagerProduction_shoryuken",
// 		"CompanyAccountManagerStaging3",
// 		"CompanyAccountManagerStaging3Util",
// 		"CompanyAccountManagerStaging3_delayed_job",
// 		"CompanyAccountManagerStaging_shoryuken",
// 		"FincAccountManagerStagingTask_bundle_exec_rake_user_patroller-compare_fid_between_requl_and_fam",
// 		"FincAccountManagerStaging_app_type_patroller-run",
// 		"FincAccountManagerStaging_bundle_exec_bin-delayed_job_run",
// 		"FincAccountManagerStaging_login_notifier_service_for_all_users-run",
// 		"FincAccountManagerStaging_secret",
// 		"FincAccountManagerStaging_shoryuken",
// 		"FincAccountManagerStaging_sidekiq",
// 		"FincAccountManagerStaging_user_patroller-compare_fid_between_requl_and_fam",
// 		"OnlineWorksProduction_shoryuken",
// 		"PersonalSupplementProduction_shoryuken",
// 		"PushNotifierProduction_delayed_job",
// 		"PushNotifierStagingCms",
// 		"PushNotifierStaging_delayed_job",
// 		"PushNotifierStaging_shoryuken",
// 		"RequlMobileProduction_old_sidekiq",
// 		"RequlMobileProduction_old_sidekiq_1",
// 		"RequlMobileRails5Staging",
// 		"RequlMobileRails5StagingPushAPI",
// 		"RequlMobileRails5StagingUtil",
// 		"RequlMobileStagingNewrelic",
// 		"WellnessDrillProduction_rising_dragon",
// 	}

// 	for _, svc := range(ecsServices) {
// 		envMap := make(map[string]string)
// 		envMap["SMTP_USERNAME"] = "AKIA3BWFIQH5G74ACQ3H"
// 		envMap["SMTP_PASSWORD"] = "BNdDzrL7ZHh9l0Hbu5k4ECsQIYeNUwyErbY2yZrJ7CHZ"
// 		updateEnvVarsForTaskDefinition(svc, envMap)
// 	}
// }

// func removeStaticAccessKeys() {
// 	// ecsServices := []string{
// 	// 	"ActivityTimelineProduction",
// 	// 	"AdviceEngineProduction",
// 	// 	"AiChatProduction",
// 	// 	"CompanyAccountManagerProduction",
// 	// 	"FiChatProduction",
// 	// 	"FincAccountManagerProduction",
// 	// 	"FincPointProduction",
// 	// 	"FincsProduction",
// 	// 	"GrouperProduction",
// 	// 	"HealthAnalysisProduction",
// 	// 	"LifelogProduction",
// 	// 	"MetisProduction",
// 	// 	"MissionServerProduction",
// 	// 	"OnboardingProduction",
// 	// 	"OnlineWorksProduction",
// 	// 	"PartnersApiProduction",
// 	// 	"PushNotifierProduction",
// 	// 	"RankieProduction",
// 	// 	"TegataProduction",
// 	// 	"TegataFrontProduction",
// 	// 	"TryServerProduction",
// 	// 	"WalkingEventsProduction",
// 	// 	"WellnessAiProduction",
// 	// 	"WellnessSurveyProduction",
// 	// }

// 	ecsServices := []string{
// 		"ActivityTimelineProductionUtil",
// 		"AdviceEngineProductionUtil",
// 		"AiChatProductionUtil",
// 		"CompanyAccountManagerProductionUtil",
// 		"FiChatProductionUtil",
// 		"FincAccountManagerProductionUtil",
// 		"FincPointProductionUtil",
// 		"FincsProductionUtil",
// 		"GrouperProductionUtil",
// 		"HealthAnalysisProductionUtil",
// 		"LifelogProductionUtil",
// 		"MetisProductionUtil",
// 		"MissionServerProductionUtil",
// 		"OnboardingProductionUtil",
// 		"OnlineWorksProductionUtil",
// 		"PartnersApiProductionUtil",
// 		"PushNotifierProductionUtil",
// 		"RankieProductionUtil",
// 		"TegataProductionUtil",
// 		"TegataFrontProductionUtil",
// 		"TryServerProductionUtil",
// 		"WalkingEventsProductionUtil",
// 		"WellnessAiProductionUtil",
// 		"WellnessSurveyProductionUtil",
// 	}

// 	for _, svc := range(ecsServices) {		
// 		containerName := "unicorn_capacity_to_cloudwatch"
// 		envVars := []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"}
// 		removeEnvVarsFromTaskDefinition(svc, containerName, envVars)
// 	}

// }



func removeStaticAccessKeysFromAllContainers() {
	ecsServices := []string{
		// "AdministokenStaging",
		// "AdministokenStagingUtil",
		// "CompanyAccountManagerStaging",
		// "CompanyAccountManagerStagingUtil",
		// "CompanyAccountManagerStaging_batch",
		// "CompanyAccountManagerStaging_delayed_job",
		// "CouponServerStaging",
		// "CouponServerStagingUtil",
		// "CouponServerStaging_batch",
		// "CouponServerStaging_shoryuken",
		// "CouponServerStaging_sidekiq",
		// "FincHelpStaging_batch",
		// "FincPrecaStaging_batch",
		// "FincProductivityStaging",
		// "FincProductivityStagingUtil",
		// "FincProductivityStaging_batch",
		// "FincProductivityStaging_sidekiq",
		// "FincsStaging",
		// "FincsStagingUtil",
		// "FincsStaging_batch",
		// "FincsStaging_shoryuken",
		// "FincsStaging_sidekiq",
		// "FoodLogStaging",
		// "FoodLogStagingUtil",
		// "HealthAnalysisStaging_batch",
		// "LifelogStagingUtil",
		// "LineMessageClientStaging",
		// "LineMessageClientStagingUtil",
		// "LineMessageClientStaging_batch",
		// "LineMessageClientStaging_runner",
		// "LineMessageClientStaging_sidekiq",
		// "MechanicalTurkFrameworkStaging",
		// "MedicalCheckStaging",
		// "MedicalCheckStagingUtil",
		// "MedicalCheckStaging_batch",
		// "MedicalCheckStaging_rising_dragon",
		// "MedicalCheckStaging_sidekiq",
		// "OnboardingStaging_batch",
		// "OnlineWorksStagingCms",
		// "OnlineWorksStaging_shoryuken",
		// "PersonalSupplementStaging_shoryuken",
		// "PushNotificationInfoRecorderStaging",
		// "RecsysStaging",
		// "RequlMobileCanaryStaging",
		// "RequlMobileStaging_shoryuken_1",
		// "WellnessDrillStaging",
		// "WellnessDrillStagingUtil",
		// "WellnessDrillStaging_batch",
		// "WellnessDrillStaging_shoryuken",
		// "WellnessDrillStaging_sidekiq",
		// "WellnessSurveyStaging",
		// "WellnessSurveyStagingUtil",
		// "WellnessSurveyStaging_batch",
		// "WellnessSurveyStaging_delayed_job",
		// "WellnessSurveyStaging_shoryuken",
	}

	for _, svc := range(ecsServices) {		
		envVars := []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "WS_AWS_ACCESS_KEY_ID", "WS_AWS_SECRET_ACCESS_KEY"}
		removeEnvVarsFromAllContainersInTaskDefinition(svc, envVars)
	}
}


func updateECSServiceToLatestTaskDefinition() {
	ecsServices := []string{
		// "AdministokenStaging",
		// "AdministokenStagingUtil",
		// "CompanyAccountManagerStaging",
		// "CompanyAccountManagerStagingUtil",
		// "CouponServerStaging",
		// "CouponServerStagingUtil",
		// "FincProductivityStaging",
		// "FincProductivityStagingUtil",
		// "FincsStaging",
		// "FincsStagingUtil",
		// "FoodLogStaging",
		// "FoodLogStagingUtil",
		// "LifelogStagingUtil",
		// "LineMessageClientStaging",
		// "LineMessageClientStagingUtil",
		// "MedicalCheckStaging",
		// "MedicalCheckStagingUtil",
		// "RequlMobileCanaryStaging",
		// "WellnessDrillStaging",
		// "WellnessDrillStagingUtil",
		// "WellnessSurveyStaging",
		// "WellnessSurveyStagingUtil",
	}

	for _, svc := range(ecsServices) {
		clusterName := "Staging"
		taskDefinitionArn := describeTaskDefinition(svc)

		if len(taskDefinitionArn) > 0 {
			fmt.Println(fmt.Sprintf("Updating %s with %s", svc, taskDefinitionArn))
			updateECSService(svc, clusterName, taskDefinitionArn)
		}
	}
}


func checkConsistencyECSServices(ecsServices []string, clusterName string) {
	for _, serviceName := range ecsServices {
		currentTaskDefinitionArn := getCurrentTaskDefinition(serviceName, clusterName)
		latestTaskDefinitionArn := describeTaskDefinition(serviceName)

		if len(currentTaskDefinitionArn) > 0 && len(latestTaskDefinitionArn) > 0 {
			if currentTaskDefinitionArn != latestTaskDefinitionArn {
				fmt.Println(fmt.Sprintf("!!! %s != %s", currentTaskDefinitionArn, latestTaskDefinitionArn))
				updateECSService(serviceName, clusterName, latestTaskDefinitionArn)
				fmt.Println(fmt.Sprintf("Updated %s to %s\n", serviceName, latestTaskDefinitionArn))
			} else {
				// fmt.Println(fmt.Sprintf("@@@ %s == %s\n", currentTaskDefinitionArn, latestTaskDefinitionArn))
			}
		}

		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func checkConsistencyStagingECSServices() {
	ecsServices := []string{
    "LetterBoxStaging",
    "WalkingEventsStaging",
    "CgmPatrolStaging",
    "WellnessAiStagingInternal",
    "WellnessSurveyStagingCms",
    "LaerdalStaging",
    "AdministokenStaging",
    "O2oClientStaging",
    "LifelogStaging",
    "OnboardingStaging",
    "FoodCategorizingStaging",
    "CommunityOrganizerStaging",
    "PartnersApiStaging",
    "ReportClientStaging",
    "PersonalSupplementStaging",
    "MetisStaging",
    "CompanyAccountManagerStagingCms",
    "RarecoilStaging",
    "RequlMobileCanaryStaging",
    "AppNavigatorStagingInternal",
    "FincPrecaStaging",
    "RequlMobileStagingPushAPI",
    "OnlineWorksStaging",
    "O2oStaging",
    "WellnessDrillStaging",
    "TryMediaClientStaging",
    "FincAccountManagerStaging",
    "CompanyAccountManagerStaging",
    "MailerStaging",
    "LineMessageClientStaging",
    "RequlMobileStaging",
    "FincHelpStaging",
    "FincPaymentStagingProtect",
    "PushNotifierStagingInternal",
    "LifelogStagingElasticSearch",
    "FincPlayStaging",
    "FincPaymentStaging",
    "FoodLogStaging",
    "CouponServerStaging",
    "HealthAnalysisStaging",
    "SnsCrawlerStaging",
    "AiChatStaging",
    "TegataStaging",
    "WellnessAiStaging",
    "FincConnectStaging",
    "FincPointStaging",
    "FincStoreStaging",
    "LifelogStagingInternal",
    "FichatclientStaging",
    "LifelogStagingElasticSearchTest",
    "GrouperStaging",
    "MedicalCheckStaging",
    "AppNavigatorStaging",
    "TegataFrontStaging",
    "PersonalDataStaging",
    "AmbassadorAdminStaging",
    "AdviceEngineStaging",
    "SnsCrawlerAdminStaging",
    "MetisStagingInternal",
    "FincProductivityStaging",
    "FincAccountManagerStagingInternal",
    "OnlineWorksStagingCms",
    "AppWebGatewayStaging",
    "LifelogStagingProtect",
    "FincPaymentStagingInternal",
    "MissionServerStaging",
    "WellnessSurveyStaging",
    "TryServerStaging",
    "ActivityTimelineStaging",
    "PushNotifierStaging",
    "FincsStaging",
    "RailsSnsStaging",
    "FiChatStaging",
    "RestOnTrialServerStaging",
    "FincAppWebClientStaging",
    "FincStoreFrontStaging",
    "RankieStaging",
	}

	checkConsistencyECSServices(ecsServices, "Staging")
}

func checkConsistencyProductionECSServices() {
	ecsServices := []string{
    "ActivityTimelineProductionUtil",
    "AdministokenProductionUtil",
    "AdviceEngineProductionUtil",
    "AiChatProductionUtil",
    "AmbassadorAdminProductionUtil",
    "AmbassadorRegistrationFormProductionUtil",
    "AppNavigatorProductionUtil",
    "AppWebGatewayProductionUtil",
    "CgmPatrolProductionUtil",
    "CommunityOrganizerProductionUtil",
    "CompanyAccountManagerProductionUtil",
    "CouponServerProductionUtil",
    "FichatclientProductionUtil",
    "FiChatProductionUtil",
    "FincAccountManagerProductionUtil",
    "FincAppWebClientProductionUtil",
    "FincConnectProductionUtil",
    "FincHelpProductionUtil",
    "FincPaymentProductionUtil",
    "FincPlayProductionUtil",
    "FincPointProductionUtil",
    "FincPrecaProductionUtil",
    "FincProductivityProductionUtil",
    "FincsProductionUtil",
    "FincStoreFrontProductionUtil",
    "FincStoreProductionUtil",
    "FoodCategorizingProductionUtil",
    "FoodLogProductionUtil",
    "GrouperProductionUtil",
    "HealthAnalysisProductionUtil",
    "LaerdalProductionUtil",
    "LetterBoxProductionUtil",
    "LifelogProductionUtil",
    "LineMessageClientProductionUtil",
    "MailerProductionUtil",
    "MechanicalTurkFrameworkProductionUtil",
    "MedicalCheckProductionUtil",
    "MetisProductionUtil",
    "MissionServerProductionUtil",
    "NluProductionUtil",
    "O2oProductionUtil",
    "OnboardingProductionUtil",
    "OnlineWorksProductionUtil",
    "PartnersApiProductionUtil",
    "PersonalSupplementProductionUtil",
    "PoseEstimationProductionUtil",
    "PushNotifierProductionUtil",
    "RankieProductionUtil",
    "RarecoilProductionUtil",
    "RecsysProductionUtil",
    "ReportClientProductionUtil",
    "RequlMobileProductionUtil",
    "SmartProductionUtil",
    "SnsCrawlerProductionUtil",
    "TegataFrontProductionUtil",
    "TegataProductionUtil",
    "TryMediaClientProductionUtil",
    "TryServerProductionUtil",
    "WalkingEventsProductionUtil",
    "WellnessAiProductionUtil",
    "WellnessDrillProductionUtil",
    "WellnessSurveyProductionUtil",
	}

	checkConsistencyECSServices(ecsServices, "ProductionUtil")
}

type ServiceSecret struct {
	ServiceName string
	ContainerName string
	SecretName  string
	SecretValue string
}

func updateSecretsForTaskDefinitions() {
	serviceSecrets := []ServiceSecret{
    // ServiceSecret{
    //   ServiceName: "ActivityTimelineStaging",
    //   ContainerName: "activity_timeline",
    //   SecretName: "ACTIVITY_TIMELINE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "ActivityTimelineStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AdministokenStaging",
    //   ContainerName: "administoken",
    //   SecretName: "ADMINSTOKEN_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AdministokenStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AdviceEngineStaging",
    //   ContainerName: "advice_engine",
    //   SecretName: "ADVICE_ENGINE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AdviceEngineStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AiChatStaging",
    //   ContainerName: "ai_chat",
    //   SecretName: "AI_CHAT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AiChatStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AmbassadorAdminStaging",
    //   ContainerName: "ambassador_admin",
    //   SecretName: "AMBASSADOR_ADMIN_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AmbassadorAdminStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AppNavigatorStaging",
    //   ContainerName: "app_navigator",
    //   SecretName: "APP_NAVIGATOR_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AppNavigatorStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "CgmPatrolStaging",
    //   ContainerName: "cgm_patrol",
    //   SecretName: "CGM_PATROL_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "CgmPatrolStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "CompanyAccountManagerStaging",
    //   ContainerName: "company_account_manager",
    //   SecretName: "COMPANY_ACCOUNT_MANAGER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "CompanyAccountManagerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "CouponServerStaging",
    //   ContainerName: "coupon_server",
    //   SecretName: "COUPON_SERVER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "CouponServerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FiChatStaging",
    //   ContainerName: "fi_chat",
    //   SecretName: "FI_CHAT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FiChatStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincHelpStaging",
    //   ContainerName: "finc_help",
    //   SecretName: "FINC_HELP_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincHelpStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPlayStaging",
    //   ContainerName: "finc_play",
    //   SecretName: "FINC_PLAY_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPlayStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPointStaging",
    //   ContainerName: "finc_point",
    //   SecretName: "FINC_POINT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPointStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPrecaStaging",
    //   ContainerName: "finc_preca",
    //   SecretName: "FINC_PRECA_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPrecaStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincProductivityStaging",
    //   ContainerName: "finc_productivity",
    //   SecretName: "FINC_PRODUCTIVITY_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincProductivityStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincsStaging",
    //   ContainerName: "fincs",
    //   SecretName: "FINCS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincsStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincStoreStaging",
    //   ContainerName: "finc_store",
    //   SecretName: "FINC_STORE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincStoreStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FoodCategorizingStaging",
    //   ContainerName: "food_categorizing",
    //   SecretName: "FOOD_CATEGORIZING_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FoodCategorizingStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FoodLogStaging",
    //   ContainerName: "food_log",
    //   SecretName: "FOOD_LOG_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FoodLogStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "HealthAnalysisStaging",
    //   ContainerName: "health_analysis",
    //   SecretName: "HEALTH_ANALYSIS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "HealthAnalysisStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "LetterBoxStaging",
    //   ContainerName: "letter_box",
    //   SecretName: "LETTER_BOX_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "LetterBoxStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "LineMessageClientStaging",
    //   ContainerName: "line_message_client",
    //   SecretName: "LINE_MESSAGE_CLIENT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "LineMessageClientStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MailerStaging",
    //   ContainerName: "mailer",
    //   SecretName: "MAILER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MailerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MedicalCheckStaging",
    //   ContainerName: "medical_check",
    //   SecretName: "MEDICAL_CHECK_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MedicalCheckStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MetisStaging",
    //   ContainerName: "metis",
    //   SecretName: "METIS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MetisStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MissionServerStaging",
    //   ContainerName: "mission_server",
    //   SecretName: "MISSION_SERVER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MissionServerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "O2oStaging",
    //   ContainerName: "o2o",
    //   SecretName: "O2O_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "O2oStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "OnboardingStaging",
    //   ContainerName: "onboarding",
    //   SecretName: "ONBOARDING_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "OnboardingStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "OnlineWorksStaging",
    //   ContainerName: "online_works",
    //   SecretName: "ONLINE_WORKS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "OnlineWorksStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PartnersApiStaging",
    //   ContainerName: "partners_api",
    //   SecretName: "PARTNERS_API_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PartnersApiStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PersonalDataStaging",
    //   ContainerName: "personal_data",
    //   SecretName: "DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PersonalDataStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PersonalSupplementStaging",
    //   ContainerName: "personal_supplement",
    //   SecretName: "PERSONAL_SUPPLEMENT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PersonalSupplementStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PushNotifierStaging",
    //   ContainerName: "push_notifier",
    //   SecretName: "PUSH_NOTIFIER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PushNotifierStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "RankieStaging",
    //   ContainerName: "rankie",
    //   SecretName: "RANKIE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "RankieStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "RequlMobileStaging",
    //   ContainerName: "requl_mobile",
    //   SecretName: "FINC_APP_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "RequlMobileStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "RestOnTrialServerStaging",
    //   ContainerName: "rest_on_trial_server",
    //   SecretName: "DATABASE_WRITE_PASSWORD",
    //   SecretValue: "RestOnTrialServerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "SnsCrawlerAdminStaging",
    //   ContainerName: "sns_crawler",
    //   SecretName: "SNS_CRAWLER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "SnsCrawlerAdminStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "SnsCrawlerStaging",
    //   ContainerName: "sns_crawler",
    //   SecretName: "SNS_CRAWLER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "SnsCrawlerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "TegataFrontStaging",
    //   ContainerName: "tegata_front",
    //   SecretName: "TEGATA_FRONT_EVENTS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "TegataFrontStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "TegataStaging",
    //   ContainerName: "tegata",
    //   SecretName: "TEGATA_EVENTS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "TegataStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "TryServerStaging",
    //   ContainerName: "try_server",
    //   SecretName: "TRY_SERVER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "TryServerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WalkingEventsStaging",
    //   ContainerName: "walking_events",
    //   SecretName: "WALKING_EVENTS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WalkingEventsStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WellnessAiStaging",
    //   ContainerName: "wellness_ai",
    //   SecretName: "WELLNESS_AI_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WellnessAiStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WellnessDrillStaging",
    //   ContainerName: "wellness_drill",
    //   SecretName: "WELLNESS_DRILL_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WellnessDrillStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WellnessSurveyStaging",
    //   ContainerName: "wellness_survey",
    //   SecretName: "WELLNESS_SURVEY_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WellnessSurveyStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincAccountManagerStaging",
    //   ContainerName: "finc_account_manager",
    //   SecretName: "FINC_ACCOUNT_MANAGER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincAccountManagerStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPaymentStaging",
    //   ContainerName: "finc_payment",
    //   SecretName: "FINC_PAYMENT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPaymentStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "LifelogStaging",
    //   ContainerName: "lifelog",
    //   SecretName: "LIFELOG_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "LifelogStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincConnectStaging",
    //   ContainerName: "finc_connect",
    //   SecretName: "FINC_CONNECT_DATABASE_PASSWORD",
    //   SecretValue: "FincConnectStaging_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "GrouperStaging",
    //   ContainerName: "grouper",
    //   SecretName: "GROUPER_DATABASE_PASSWORD",
    //   SecretValue: "GrouperStaging_DATABASE_WRITE_PASSWORD",
    // },

    // Production
    // ServiceSecret{
    //   ServiceName: "ActivityTimelineProduction",
    //   ContainerName: "activity_timeline",
    //   SecretName: "ACTIVITY_TIMELINE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "ActivityTimelineProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AdministokenProduction",
    //   ContainerName: "administoken",
    //   SecretName: "ADMINSTOKEN_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AdministokenProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AdviceEngineProduction",
    //   ContainerName: "advice_engine",
    //   SecretName: "ADVICE_ENGINE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AdviceEngineProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AiChatProduction",
    //   ContainerName: "ai_chat",
    //   SecretName: "AI_CHAT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AiChatProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AmbassadorAdminProduction",
    //   ContainerName: "ambassador_admin",
    //   SecretName: "AMBASSADOR_ADMIN_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AmbassadorAdminProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "AppNavigatorProduction",
    //   ContainerName: "app_navigator",
    //   SecretName: "APP_NAVIGATOR_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "AppNavigatorProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "CgmPatrolProduction",
    //   ContainerName: "cgm_patrol",
    //   SecretName: "CGM_PATROL_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "CgmPatrolProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "CompanyAccountManagerProduction",
    //   ContainerName: "company_account_manager",
    //   SecretName: "COMPANY_ACCOUNT_MANAGER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "CompanyAccountManagerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "CouponServerProduction",
    //   ContainerName: "coupon_server",
    //   SecretName: "COUPON_SERVER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "CouponServerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FiChatProduction",
    //   ContainerName: "fi_chat",
    //   SecretName: "FI_CHAT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FiChatProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincAccountManagerProduction",
    //   ContainerName: "finc_account_manager",
    //   SecretName: "FINC_ACCOUNT_MANAGER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincAccountManagerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincHelpProduction",
    //   ContainerName: "finc_help",
    //   SecretName: "FINC_HELP_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincHelpProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPaymentProduction",
    //   ContainerName: "finc_payment",
    //   SecretName: "FINC_PAYMENT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPaymentProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPlayProduction",
    //   ContainerName: "finc_play",
    //   SecretName: "FINC_PLAY_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPlayProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPointProduction",
    //   ContainerName: "finc_point",
    //   SecretName: "FINC_POINT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPointProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincPrecaProduction",
    //   ContainerName: "finc_preca",
    //   SecretName: "FINC_PRECA_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincPrecaProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincProductivityProduction",
    //   ContainerName: "finc_productivity",
    //   SecretName: "FINC_PRODUCTIVITY_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincProductivityProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincsProduction",
    //   ContainerName: "fincs",
    //   SecretName: "FINCS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincsProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincStoreProduction",
    //   ContainerName: "finc_store",
    //   SecretName: "FINC_STORE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FincStoreProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FoodCategorizingProduction",
    //   ContainerName: "food_categorizing",
    //   SecretName: "FOOD_CATEGORIZING_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FoodCategorizingProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FoodLogProduction",
    //   ContainerName: "food_log",
    //   SecretName: "FOOD_LOG_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "FoodLogProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "HealthAnalysisProduction",
    //   ContainerName: "health_analysis",
    //   SecretName: "HEALTH_ANALYSIS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "HealthAnalysisProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "LetterBoxProduction",
    //   ContainerName: "letter_box",
    //   SecretName: "LETTER_BOX_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "LetterBoxProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "LifelogProduction",
    //   ContainerName: "lifelog",
    //   SecretName: "LIFELOG_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "LifelogProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "LineMessageClientProduction",
    //   ContainerName: "line_message_client",
    //   SecretName: "LINE_MESSAGE_CLIENT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "LineMessageClientProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MailerProduction",
    //   ContainerName: "mailer",
    //   SecretName: "MAILER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MailerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MedicalCheckProduction",
    //   ContainerName: "medical_check",
    //   SecretName: "MEDICAL_CHECK_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MedicalCheckProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MetisProduction",
    //   ContainerName: "metis",
    //   SecretName: "METIS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MetisProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "MissionServerProduction",
    //   ContainerName: "mission_server",
    //   SecretName: "MISSION_SERVER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "MissionServerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "O2oProduction",
    //   ContainerName: "o2o",
    //   SecretName: "O2O_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "O2oProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "OnboardingProduction",
    //   ContainerName: "onboarding",
    //   SecretName: "ONBOARDING_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "OnboardingProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PartnersApiProduction",
    //   ContainerName: "partners_api",
    //   SecretName: "PARTNERS_API_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PartnersApiProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PersonalSupplementProduction",
    //   ContainerName: "personal_supplement",
    //   SecretName: "PERSONAL_SUPPLEMENT_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PersonalSupplementProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "PushNotifierProduction",
    //   ContainerName: "push_notifier",
    //   SecretName: "PUSH_NOTIFIER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "PushNotifierProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "RankieProduction",
    //   ContainerName: "rankie",
    //   SecretName: "RANKIE_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "RankieProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "RequlMobileProduction",
    //   ContainerName: "requl_mobile",
    //   SecretName: "FINC_APP_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "RequlMobileProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "RestOnTrialServerProduction",
    //   ContainerName: "rest_on_trial_server",
    //   SecretName: "DATABASE_WRITE_PASSWORD",
    //   SecretValue: "RestOnTrialServerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "SnsCrawlerAdminProduction",
    //   ContainerName: "sns_crawler",
    //   SecretName: "SNS_CRAWLER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "SnsCrawlerAdminProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "SnsCrawlerProduction",
    //   ContainerName: "sns_crawler",
    //   SecretName: "SNS_CRAWLER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "SnsCrawlerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "TegataFrontProduction",
    //   ContainerName: "tegata_front",
    //   SecretName: "TEGATA_FRONT_EVENTS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "TegataFrontProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "TegataProduction",
    //   ContainerName: "tegata",
    //   SecretName: "TEGATA_EVENTS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "TegataProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "TryServerProduction",
    //   ContainerName: "try_server",
    //   SecretName: "TRY_SERVER_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "TryServerProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WalkingEventsProduction",
    //   ContainerName: "walking_events",
    //   SecretName: "WALKING_EVENTS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WalkingEventsProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WellnessAiProduction",
    //   ContainerName: "wellness_ai",
    //   SecretName: "WELLNESS_AI_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WellnessAiProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WellnessDrillProduction",
    //   ContainerName: "wellness_drill",
    //   SecretName: "WELLNESS_DRILL_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WellnessDrillProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "WellnessSurveyProduction",
    //   ContainerName: "wellness_survey",
    //   SecretName: "WELLNESS_SURVEY_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "WellnessSurveyProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "OnlineWorksProduction",
    //   ContainerName: "online_works",
    //   SecretName: "ONLINE_WORKS_DATABASE_WRITE_PASSWORD",
    //   SecretValue: "OnlineWorksProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "FincConnectProduction",
    //   ContainerName: "finc_connect",
    //   SecretName: "FINC_CONNECT_DATABASE_PASSWORD",
    //   SecretValue: "FincConnectProduction_DATABASE_WRITE_PASSWORD",
    // },
    // ServiceSecret{
    //   ServiceName: "GrouperProduction",
    //   ContainerName: "grouper",
    //   SecretName: "GROUPER_DATABASE_PASSWORD",
    //   SecretValue: "GrouperProduction_DATABASE_WRITE_PASSWORD",
    // },
	}

	for _, ss := range serviceSecrets {
		secretMap := make(map[string]string)
		secretMap[ss.SecretName] = ss.SecretValue
		updateSecrets(ss.ServiceName, ss.ContainerName, secretMap)
		fmt.Println()
	}
}


func main() {
	updateSecretsForTaskDefinitions()
	time.Sleep(time.Duration(100) * time.Millisecond)
}
