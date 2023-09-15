package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/ecs"
)


type ECSService struct {
	TaskDefinition string
	ContainerName  string
}


func listServices(clusterName string) {
	log.Println("List All ECS Services")
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	input := &ecs.ListServicesInput{
		Cluster: aws.String(clusterName),
		MaxResults: aws.Int64(100),
	}

	output, err := ecsSvc.ListServices(input)
	if err != nil {
		log.Fatalln(err.Error())
	}

	count := 0

	for _, arn := range output.ServiceArns {
		count++
		if !strings.Contains(*arn, "SidekiqDashboard") && !strings.Contains(*arn, "Fluentd") {
			log.Println(count, strings.Split(*arn, "/")[1])
		}
	}

	nextToken := output.NextToken
	for nextToken != nil && len(*nextToken) > 0 {
		input = &ecs.ListServicesInput{
			Cluster: aws.String(clusterName),
			MaxResults: aws.Int64(100),
			NextToken: nextToken,
		}

		output, err = ecsSvc.ListServices(input)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for _, arn := range output.ServiceArns {
			count++
			if !strings.Contains(*arn, "SidekiqDashboard") && !strings.Contains(*arn, "Fluentd") {
				log.Println(count, strings.Split(*arn, "/")[1])
			}
		}

		nextToken = output.NextToken
	}
}


func registerTaskDefinition(service *ECSService) {
	log.Println("Create TaskDefinition")
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ecsSvc := ecs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	cloudwatchlogSvc := cloudwatchlogs.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	// Get Current TaskDefinition
	taskParams := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(service.TaskDefinition),
	}
	taskDef, err := ecsSvc.DescribeTaskDefinition(taskParams)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var targetContainer *ecs.ContainerDefinition = nil
	for _, container := range taskDef.TaskDefinition.ContainerDefinitions {
		if *container.Name == service.ContainerName {
			targetContainer = container
			environments := container.Environment
			for _, env := range environments {
				if *env.Name == "UNICORN_WORKER_NUM" {
					*env.Value = "1"
					break
				}
			}
			newLogGroup := fmt.Sprintf("/ecs/%s_batch", *taskDef.TaskDefinition.Family)
			// check if log group exists and create if not
			params := &cloudwatchlogs.CreateLogGroupInput{
				LogGroupName: aws.String(newLogGroup),
			}
			_, err := cloudwatchlogSvc.CreateLogGroup(params)
			if err != nil {
				if awsErr, ok := err.(awserr.Error); ok {
					if awsErr.Code() == "ResourceAlreadyExistsException" {
						log.Println("Log group already exists. Skip creating log group.")
					} else {
						log.Fatalln("Error:", awsErr.Message())
					}
				} else {
					log.Fatalln(err.Error())
				}
			}

			options := container.LogConfiguration.Options
			options["awslogs-group"] = aws.String(newLogGroup)
			if *container.LogConfiguration.LogDriver == "awslogs" {
				container.LogConfiguration.SetOptions(options)
			}
			if *container.LogConfiguration.LogDriver == "fluentd" {
				options["awslogs-stream-prefix"] = aws.String("ecs")
				options["awslogs-region"] = aws.String("ap-northeast-1")
				delete(options, "fluentd-address")
				delete(options, "tag")
				container.LogConfiguration.SetOptions(options)
				container.LogConfiguration.SetLogDriver("awslogs")
			}
		}
	}

	targetContainer.Links = nil
	targetContainer.MountPoints = nil
	targetContainer.VolumesFrom = nil

	batchTaskDefInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    []*ecs.ContainerDefinition{targetContainer},
		Cpu:                     taskDef.TaskDefinition.Cpu,
		ExecutionRoleArn:        taskDef.TaskDefinition.ExecutionRoleArn,
		Family:                  aws.String(fmt.Sprintf("%s_%s", *taskDef.TaskDefinition.Family, "batch")),
		Memory:                  taskDef.TaskDefinition.Memory,
		NetworkMode:             taskDef.TaskDefinition.NetworkMode,
		PlacementConstraints:    taskDef.TaskDefinition.PlacementConstraints,
		RequiresCompatibilities: taskDef.TaskDefinition.RequiresCompatibilities,
		TaskRoleArn:             taskDef.TaskDefinition.TaskRoleArn,
		Volumes:                 taskDef.TaskDefinition.Volumes,
	}
	batchTaskDef, err := ecsSvc.RegisterTaskDefinition(batchTaskDefInput)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(*batchTaskDef.TaskDefinition.TaskDefinitionArn)
}


func main() {
	stagingServices := []*ECSService{
		// &ECSService{
		// 	TaskDefinition: "FiChatStaging",
		// 	ContainerName: "fi_chat",
		// },

		// &ECSService{
		// 	TaskDefinition: "MissionServerStaging",
		// 	ContainerName: "mission_server",
		// },

		// &ECSService{
		// 	TaskDefinition: "ActivityTimelineStaging",
		// 	ContainerName: "activity_timeline",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincPointStaging",
		// 	ContainerName: "finc_point",
		// },

		// &ECSService{
		// 	TaskDefinition: "CouponServerStaging",
		// 	ContainerName: "coupon_server",
		// },

		// &ECSService{
		// 	TaskDefinition: "GrouperStaging",
		// 	ContainerName: "grouper",
		// },

		// &ECSService{
		// 	TaskDefinition: "PersonalSupplementStaging",
		// 	ContainerName: "personal_supplement",
		// },

		// &ECSService{
		// 	TaskDefinition: "RankieStaging",
		// 	ContainerName: "rankie",
		// },

		// &ECSService{
		// 	TaskDefinition: "RequlMobileStaging",
		// 	ContainerName: "requl_mobile",
		// },

		// &ECSService{
		// 	TaskDefinition: "TegataFrontStaging",
		// 	ContainerName: "tegata_front",
		// },

		// &ECSService{
		// 	TaskDefinition: "TryServerStaging",
		// 	ContainerName: "try_server",
		// },

		// &ECSService{
		// 	TaskDefinition: "WalkingEventsStaging",
		// 	ContainerName: "walking_events",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincStoreStaging",
		// 	ContainerName: "finc_store",
		// },

		// &ECSService{
		// 	TaskDefinition: "WellnessDrillStaging",
		// 	ContainerName: "wellness_drill",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincHelpStaging",
		// 	ContainerName: "finc_help",
		// },

		// &ECSService{
		// 	TaskDefinition: "WellnessSurveyStaging",
		// 	ContainerName: "wellness_survey",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincsStaging",
		// 	ContainerName: "fincs",
		// },

		// &ECSService{
		// 	TaskDefinition: "MedicalCheckStaging",
		// 	ContainerName: "medical_check",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincProductivityStaging",
		// 	ContainerName: "finc_productivity",
		// },

		// &ECSService{
		// 	TaskDefinition: "SnsCrawlerStaging",
		// 	ContainerName: "sns_crawler",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincPaymentStaging",
		// 	ContainerName: "finc_payment",
		// },

		// &ECSService{
		// 	TaskDefinition: "CompanyAccountManagerStaging",
		// 	ContainerName: "company_account_manager",
		// },

		&ECSService{
			TaskDefinition: "AdviceEngineStaging",
			ContainerName: "advice_engine",
		},

		&ECSService{
			TaskDefinition: "AiChatStaging",
			ContainerName: "ai_chat",
		},

		&ECSService{
			TaskDefinition: "AmbassadorAdminStaging",
			ContainerName: "ambassador_admin",
		},

		&ECSService{
			TaskDefinition: "AmbassadorManagerStaging",
			ContainerName: "ambassador_manager",
		},

		&ECSService{
			TaskDefinition: "AppNavigatorStaging",
			ContainerName: "app_navigator",
		},

		&ECSService{
			TaskDefinition: "FincAccountManagerStaging",
			ContainerName: "finc_account_manager",
		},

		&ECSService{
			TaskDefinition: "FincConnectStaging",
			ContainerName: "finc_connect",
		},

		&ECSService{
			TaskDefinition: "FincPlayStaging",
			ContainerName: "finc_play",
		},

		&ECSService{
			TaskDefinition: "FincPrecaStaging",
			ContainerName: "finc_preca",
		},

		&ECSService{
			TaskDefinition: "HealthAnalysisStaging",
			ContainerName: "health_analysis",
		},

		&ECSService{
			TaskDefinition: "LetterBoxStaging",
			ContainerName: "letter_box",
		},

		&ECSService{
			TaskDefinition: "LifelogStaging",
			ContainerName: "lifelog",
		},

		&ECSService{
			TaskDefinition: "LineMessageClientStaging",
			ContainerName: "line_message_client",
		},

		&ECSService{
			TaskDefinition: "MetisStaging",
			ContainerName: "metis",
		},

		&ECSService{
			TaskDefinition: "OnboardingStaging",
			ContainerName: "onboarding",
		},

		&ECSService{
			TaskDefinition: "OnlineWorksStaging",
			ContainerName: "online_works",
		},

		&ECSService{
			TaskDefinition: "PushNotifierStaging",
			ContainerName: "push_notifier",
		},

		&ECSService{
			TaskDefinition: "TegataStaging",
			ContainerName: "tegata",
		},

		&ECSService{
			TaskDefinition: "WellnessAiStaging",
			ContainerName: "wellness_ai",
		},
	}


	productionServices := []*ECSService{
		// &ECSService{
		// 	TaskDefinition: "FiChatProduction",
		// 	ContainerName: "fi_chat",
		// },

		// &ECSService{
		// 	TaskDefinition: "MissionServerProduction",
		// 	ContainerName: "mission_server",
		// },

		// &ECSService{
		// 	TaskDefinition: "ActivityTimelineProduction",
		// 	ContainerName: "activity_timeline",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincPointProduction",
		// 	ContainerName: "finc_point",
		// },

		// &ECSService{
		// 	TaskDefinition: "CouponServerProduction",
		// 	ContainerName: "coupon_server",
		// },

		// &ECSService{
		// 	TaskDefinition: "GrouperProduction",
		// 	ContainerName: "grouper",
		// },

		// &ECSService{
		// 	TaskDefinition: "PersonalSupplementProduction",
		// 	ContainerName: "personal_supplement",
		// },

		// &ECSService{
		// 	TaskDefinition: "RankieProduction",
		// 	ContainerName: "rankie",
		// },

		// &ECSService{
		// 	TaskDefinition: "RequlMobileProduction",
		// 	ContainerName: "requl_mobile",
		// },

		// &ECSService{
		// 	TaskDefinition: "TegataFrontProduction",
		// 	ContainerName: "tegata_front",
		// },

		// &ECSService{
		// 	TaskDefinition: "TryServerProduction",
		// 	ContainerName: "try_server",
		// },

		// &ECSService{
		// 	TaskDefinition: "WalkingEventsProduction",
		// 	ContainerName: "walking_events",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincStoreProduction",
		// 	ContainerName: "finc_store",
		// },

		// &ECSService{
		// 	TaskDefinition: "WellnessDrillProduction",
		// 	ContainerName: "wellness_drill",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincHelpProduction",
		// 	ContainerName: "finc_help",
		// },

		// &ECSService{
		// 	TaskDefinition: "WellnessSurveyProduction",
		// 	ContainerName: "wellness_survey",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincsProduction",
		// 	ContainerName: "fincs",
		// },

		// &ECSService{
		// 	TaskDefinition: "MedicalCheckProduction",
		// 	ContainerName: "medical_check",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincProductivityProduction",
		// 	ContainerName: "finc_productivity",
		// },

		// &ECSService{
		// 	TaskDefinition: "SnsCrawlerProduction",
		// 	ContainerName: "sns_crawler",
		// },

		// &ECSService{
		// 	TaskDefinition: "FincPaymentProduction",
		// 	ContainerName: "finc_payment",
		// },

		// &ECSService{
		// 	TaskDefinition: "CompanyAccountManagerProduction",
		// 	ContainerName: "company_account_manager",
		// },

		&ECSService{
			TaskDefinition: "AdviceEngineProduction",
			ContainerName: "advice_engine",
		},

		&ECSService{
			TaskDefinition: "AiChatProduction",
			ContainerName: "ai_chat",
		},

		&ECSService{
			TaskDefinition: "AmbassadorAdminProduction",
			ContainerName: "ambassador_admin",
		},

		// &ECSService{
		// 	TaskDefinition: "AmbassadorManagerProduction",
		// 	ContainerName: "ambassador_manager",
		// },

		&ECSService{
			TaskDefinition: "AppNavigatorProduction",
			ContainerName: "app_navigator",
		},

		&ECSService{
			TaskDefinition: "FincAccountManagerProduction",
			ContainerName: "finc_account_manager",
		},

		&ECSService{
			TaskDefinition: "FincConnectProduction",
			ContainerName: "finc_connect",
		},

		&ECSService{
			TaskDefinition: "FincPlayProduction",
			ContainerName: "finc_play",
		},

		&ECSService{
			TaskDefinition: "FincPrecaProduction",
			ContainerName: "finc_preca",
		},

		&ECSService{
			TaskDefinition: "HealthAnalysisProduction",
			ContainerName: "health_analysis",
		},

		&ECSService{
			TaskDefinition: "LetterBoxProduction",
			ContainerName: "letter_box",
		},

		&ECSService{
			TaskDefinition: "LifelogProduction",
			ContainerName: "lifelog",
		},

		&ECSService{
			TaskDefinition: "LineMessageClientProduction",
			ContainerName: "line_message_client",
		},

		&ECSService{
			TaskDefinition: "MetisProduction",
			ContainerName: "metis",
		},

		&ECSService{
			TaskDefinition: "OnboardingProduction",
			ContainerName: "onboarding",
		},

		&ECSService{
			TaskDefinition: "OnlineWorksProduction",
			ContainerName: "online_works",
		},

		&ECSService{
			TaskDefinition: "PushNotifierProduction",
			ContainerName: "push_notifier",
		},

		&ECSService{
			TaskDefinition: "TegataProduction",
			ContainerName: "tegata",
		},

		&ECSService{
			TaskDefinition: "WellnessAiProduction",
			ContainerName: "wellness_ai",
		},
	}


	for i := range(stagingServices) {
		registerTaskDefinition(stagingServices[i])
	}

	for i := range(productionServices) {
		registerTaskDefinition(productionServices[i])
	}
}
