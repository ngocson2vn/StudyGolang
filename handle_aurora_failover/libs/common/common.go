package common

const (
	STATUS_SUCCESS = "SUCCESS"
)

type AuroraCluster struct {
	ClusterName            string             `yaml:"cluster_name"`
	DependentServices      []string           `yaml:"dependent_services"`
	DependentWorkers       []string           `yaml:"dependent_workers"`
}

type SNSMessage struct {
	EventSource            string             `json:"Event Source"`
	EventTime              string             `json:"Event Time"`
	IdentifierLink         string             `json:"Identifier Link"`
	SourceID               string             `json:"Source ID"`
	EventID                string             `json:"Event ID"`
	EventMessage           string             `json:"Event Message"`
}

