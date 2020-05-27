package main

type Services struct {
	Services []Services_sub2 `json:"services"`
}

type Services_sub2 struct {
	ActiveGroup  string        `json:"active_group"`
	HealthStatus Services_sub1 `json:"health_status"`
	JobType      string        `json:"job_type"`
	Link         string        `json:"link"`
	Name         string        `json:"name"`
	Orchestrator string        `json:"orchestrator"`
	Squad        string        `json:"squad"`
	State        int64         `json:"state"`
	Team         string        `json:"team"`
}

type Services_sub1 struct {
	ContainerCount int64    `json:"container_count"`
	Infos          []string `json:"infos"`
	InstanceCount  int64    `json:"instance_count"`
	State          string   `json:"state"`
}
