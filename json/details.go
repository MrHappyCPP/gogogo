package main

type Details struct {
	Active       string       `json:"active"`
	AlertLevel   int64        `json:"alert_level"`
	Groups       Details_sub5 `json:"groups"`
	JobType      string       `json:"job_type"`
	Name         string       `json:"name"`
	Orchestrator string       `json:"orchestrator"`
}

type Details_sub5 struct {
	Blue  Details_sub4 `json:"blue"`
	Green Details_sub4 `json:"green"`
}

type Details_sub2 struct {
	ConsulServiceID string       `json:"consul_service_id"`
	ContainerID     string       `json:"container_id"`
	Environment     []string     `json:"environment"`
	Host            string       `json:"host"`
	ID              string       `json:"id"`
	Image           string       `json:"image"`
	IP              string       `json:"ip"`
	JobVersion      int64        `json:"job_version"`
	Name            string       `json:"name"`
	Port            int64        `json:"port"`
	PresetCPU       int64        `json:"preset_cpu"`
	PresetDisk      int64        `json:"preset_disk"`
	PresetMemory    int64        `json:"preset_memory"`
	Sha             string       `json:"sha"`
	State           Details_sub1 `json:"state"`
	URL             string       `json:"url"`
	UsageCPU        int64        `json:"usage_cpu"`
	UsageMemory     int64        `json:"usage_memory"`
	User            string       `json:"user"`
}

type Details_sub4 struct {
	Container        []Details_sub2 `json:"container"`
	ExternalName     string         `json:"external_name"`
	Fqdn             string         `json:"fqdn"`
	HealthStatus     Details_sub3   `json:"health_status"`
	InstanceCount    int64          `json:"instance_count"`
	Orchestrator     string         `json:"orchestrator"`
	RoutingZone      string         `json:"routing_zone"`
	ServiceImageName string         `json:"service_image_name"`
	ServiceVersion   string         `json:"service_version"`
	Squad            string         `json:"squad"`
	Team             string         `json:"team"`
}

type Details_sub3 struct {
	ContainerCount int64         `json:"container_count"`
	Infos          []interface{} `json:"infos"`
	InstanceCount  int64         `json:"instance_count"`
	State          string        `json:"state"`
}

type Details_sub1 struct {
	Error       string `json:"error"`
	ExitCode    int64  `json:"exit_code"`
	FinishedAt  string `json:"finished_at"`
	HealthCheck string `json:"health_check"`
	HealthState string `json:"health_state"`
	OomKilled   bool   `json:"oom_killed"`
	StartedAt   string `json:"started_at"`
	Status      string `json:"status"`
}
