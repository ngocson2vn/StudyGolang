package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"

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


func listECSServices(clusterName string) ([]string, error) {
	serviceList := []string{}
	sess, err := session.NewSession()
	if err != nil {
		return []string{}, err
	}
	svc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	inputParams := &ecs.ListServicesInput{
		Cluster: aws.String(clusterName),
		MaxResults: aws.Int64(100),
	}
	output, err := svc.ListServices(inputParams)
	if err != nil {
		log.Println(err.Error())
		return []string{}, err
	}

	for _, serviceArn := range output.ServiceArns {
		serviceList = append(serviceList, strings.Split(*serviceArn, "/")[1])
	}

	for output.NextToken != nil {
		log.Println(*output.NextToken)
		inputParams = &ecs.ListServicesInput{
			Cluster: aws.String(clusterName),
			MaxResults: aws.Int64(100),
			NextToken: output.NextToken,
		}

		output, err = svc.ListServices(inputParams)
		if err != nil {
			log.Println(err.Error())
			return []string{}, err
		}

		for _, serviceArn := range output.ServiceArns {
			serviceList = append(serviceList, strings.Split(*serviceArn, "/")[1])
		}
	}

	return serviceList, nil
}


func checkConsistencyECSServices(ecsServices []string, clusterName string) {
	for _, serviceName := range ecsServices {
		currentTaskDefinitionArn := getCurrentTaskDefinition(serviceName, clusterName)
		latestTaskDefinitionArn := describeTaskDefinition(serviceName)

		if len(currentTaskDefinitionArn) > 0 && len(latestTaskDefinitionArn) > 0 {
			if currentTaskDefinitionArn != latestTaskDefinitionArn {
				fmt.Println(fmt.Sprintf("!!! %s != %s", currentTaskDefinitionArn, latestTaskDefinitionArn))
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

func listTaskDefinitionsUsedByServices(clusterName string) ([]string, error) {
	taskDefList := []string{}

	serviceList, err := listECSServices(clusterName)
	if err != nil {
		return []string{}, err
	}

	for _, s := range(serviceList) {
		taskDef := getCurrentTaskDefinition(s, clusterName)
		if len(taskDef) == 0 {
			log.Printf("Could not get current Task Definition of %s\n", s)
		} else {
			taskDefList = append(taskDefList, taskDef)
		}
	}

	return taskDefList, nil
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


func main() {
	clusterName := "Staging"
	taskDefList, err := listTaskDefinitionsUsedByServices(clusterName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, t := range taskDefList {
		fmt.Println(t)
	}
}
