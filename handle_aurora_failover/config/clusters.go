package config

import (
	. "github.com/FiNCDeveloper/finc_server/handle_aurora_failover/libs/common"
)

var AuroraClusters = []AuroraCluster{
	AuroraCluster{
		ClusterName: "recsys-production-aurora",
		DependentServices: []string{
			"RecsysProduction",
			"RecsysProductionInternal",
		},
		DependentWorkers: []string{
			"RecsysProduction_celery",
		},
	},

	AuroraCluster{
		ClusterName: "rankie-production-aurora-cluster",
		DependentServices: []string{
			"RankieProduction",
		},
		DependentWorkers: []string{
			"RankieProduction_shoryuken",
			"RankieProduction_shoryuken_1",
			"RankieProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "lifelog-production-aurora",
		DependentServices: []string{
			"LifelogProductionInternal",
			"LifelogProductionProtect",
		},
		DependentWorkers: []string{
			"LifelogProduction_sidekiq",
			"LifelogProduction_sidekiq_1",
		},
	},

	AuroraCluster{
		ClusterName: "app-navigator-production-aurora",
		DependentServices: []string{
			"AppNavigatorProduction",
			"AppNavigatorProductionInternal",
		},
		DependentWorkers: []string{
			"AppNavigatorProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "fincapp-production-aurora",
		DependentServices: []string{
			"RequlMobileProduction",
			"RequlMobileProductionCom",
			"RequlMobileProductionPushAPI",
			"WellnessAiProduction",
			"WellnessAiProductionInternal",
		},
		DependentWorkers: []string{
			"RequlMobileProduction_delayed_job",
			"RequlMobileProduction_shoryuken",
			"RequlMobileProduction_shoryuken_1",
			"RequlMobileProduction_shoryuken_2",
			"RequlMobileProduction_sidekiq",
			"RequlMobileProduction_sidekiq_1",
			"WellnessAiProduction_delayed_job",
			"WellnessAiProduction_shoryuken",
			"WellnessAiProduction_shoryuken_1",
			"WellnessAiProduction_shoryuken_2",
			"WellnessAiProduction_shoryuken_3",
			"WellnessAiProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "wellness-ai-production-aurora",
		DependentServices: []string{
			"WellnessAiProduction",
			"WellnessAiProductionInternal",
		},
		DependentWorkers: []string{
			"WellnessAiProduction_delayed_job",
			"WellnessAiProduction_shoryuken",
			"WellnessAiProduction_shoryuken_1",
			"WellnessAiProduction_shoryuken_2",
			"WellnessAiProduction_shoryuken_3",
			"WellnessAiProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "push-notifier-production-aurora",
		DependentServices: []string{
			"PushNotifierProductionInternal",
		},
		DependentWorkers: []string{
		},
	},

	AuroraCluster{
		ClusterName: "ai-chat-production-aurora-cluster",
		DependentServices: []string{
			"AiChatProduction",
		},
		DependentWorkers: []string{
			"AiChatProduction_shoryuken",
			"AiChatProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "finc-account-manager-production-aurora",
		DependentServices: []string{
			"FincAccountManagerProduction",
			"FincAccountManagerProductionInternal",
		},
		DependentWorkers: []string{
			"FincAccountManagerProduction_delayed_job",
		},
	},

	AuroraCluster{
		ClusterName: "common-production-aurora-multiaz",
		DependentServices: []string{
			"AmbassadorAdminProduction",
			"CgmPatrolProduction",
			"FincConnectProduction",
			"FincHelpProduction",
			"FoodCategorizingProduction",
			"FoodLogProduction",
			"HealthAnalysisProduction",
			"MailerProduction",
			"O2oProduction",
			"OnlineWorksProduction",
			"PartnersApiProduction",
			"PersonalSupplementProduction",
			"SnsCrawlerProduction",
		},
		DependentWorkers: []string{
			"AmbassadorAdminProduction_sidekiq",
			"CgmPatrolProduction_sidekiq",
			"FincConnectProduction_shoryuken",
			"FincConnectProduction_sidekiq",
			"HealthAnalysisProduction_shoryuken",
			"MailerProduction_delayed_job",
			"O2oProduction_delayed_job",
			"OnlineWorksProduction_delayed_job",
			"OnlineWorksProduction_shoryuken",
			"PersonalSupplementProduction_sidekiq",
			"PersonalSupplementProduction_shoryuken",
		},
	},

	AuroraCluster{
		ClusterName: "finc4biz-common-production-aurora",
		DependentServices: []string{
			"CompanyAccountManagerProduction",
			"CompanyAccountManagerProductionCms",
			"FincsProduction",
			"WellnessDrillProduction",
			"WellnessSurveyProduction",
			"WellnessSurveyProductionCms",
		},
		DependentWorkers: []string{
			"CompanyAccountManagerProduction_delayed_job",
			"FincsProduction_sidekiq",
			"FincsProduction_shoryuken",
			"WellnessDrillProduction_sidekiq",
			"WellnessDrillProduction_shoryuken",
			"WellnessSurveyProduction_delayed_job",
			"WellnessSurveyProduction_shoryuken",
		},
	},

	AuroraCluster{
		ClusterName: "advice-engine-production-aurora-cluster",
		DependentServices: []string{
			"AdviceEngineProduction",
		},
		DependentWorkers: []string{
			"AdviceEngineProduction_shoryuken",
			"AdviceEngineProduction_sidekiq",
			"AdviceEngineProduction_sidekiq_report",
		},
	},

	AuroraCluster{
		ClusterName: "activity-timeline-production-aurora",
		DependentServices: []string{
			"ActivityTimelineProduction",
		},
		DependentWorkers: []string{
			"ActivityTimelineProduction_shoryuken_1",
			"ActivityTimelineProduction_delayed_job",
			"ActivityTimelineProduction_shoryuken",
		},
	},

	AuroraCluster{
		ClusterName: "mission-server-production-aurora",
		DependentServices: []string{
			"MissionServerProduction",
		},
		DependentWorkers: []string{
			"MissionServerProduction_shoryuken",
			"MissionServerProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "fi-chat-production-aurora",
		DependentServices: []string{
			"FiChatProduction",
		},
		DependentWorkers: []string{
			"FiChatProduction_shoryuken",
			"FiChatProduction_delayed_job",
			"FiChatProduction_delayed_job_1",
			"FiChatProduction_delayed_job_2",
			"FiChatProduction_delayed_job_3",
			"FiChatProduction_sidekiq",
			"FiChatProduction_sidekiq_1",
			"FiChatProduction_shoryuken_1",
			"FiChatProduction_shoryuken_2",
		},
	},

	AuroraCluster{
		ClusterName: "finc-payment-production-aurora-cluster",
		DependentServices: []string{
			"FincPaymentProduction",
			"FincPaymentProductionProtect",
			"FincPaymentProductionInternal",
		},
		DependentWorkers: []string{
			"FincPaymentProduction_delayed_job",
		},
	},

	AuroraCluster{
		ClusterName: "finc-point-production-aurora",
		DependentServices: []string{
			"FincPointProduction",
		},
		DependentWorkers: []string{
			"FincPointProduction_shoryuken",
			"FincPointProduction_sidekiq",
		},
	},

	AuroraCluster{
		ClusterName: "common-production-aurora57",
		DependentServices: []string{
			"LetterBoxProduction",
			"LineMessageClientProduction",
		},
		DependentWorkers: []string{
			"LineMessageClientProduction_sidekiq",
			"LineMessageClientProduction_runner",
		},
	},
}
