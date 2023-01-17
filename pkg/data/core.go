package data

// should be able to generate this from the OpenAPI spec somehow

// https://www.jvt.me/posts/2022/04/06/generate-go-struct-openapi/

type State struct {
	Id           string         `json:"id"`
	Type         string         `json:"type"`
	Name         string         `json:"name"`
	Timestamp    string         `json:"timestamp"`
	Message      string         `json:"message"`
	Data         any            `json:"data"`
	StateDetails map[string]any `json:"state_details"`
}

type Flow struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
}

type FlowRun struct {
	Id                       string         `json:"id"`
	Created                  string         `json:"created"`
	Updated                  string         `json:"updated"`
	Name                     string         `json:"name"`
	FlowId                   string         `json:"flow_id"`
	StateId                  string         `json:"state_id"`
	DeploymentId             string         `json:"deployment_id"`
	WorkQueueName            string         `json:"work_queue_name"`
	FlowVersion              float32        `json:"flow_version"`
	Parameters               map[string]any `json:"parameters"`
	IdempotencyKey           string         `json:"idempotency_key"`
	Context                  map[string]any `json:"context"`
	EmpiricalPolicy          map[string]any `json:"empirical_policy"`
	Tags                     []string       `json:"tags"`
	ParentTaskRunId          string         `json:"parent_task_run_id"`
	StateType                string         `json:"state_type"`
	StateName                string         `json:"state_name"`
	RunCount                 int            `json:"run_count"`
	ExpectedStartTime        string         `json:"expected_start_time"`
	NextScheduledStartTime   string         `json:"next_scheduled_start_time"`
	StartTime                string         `json:"start_time"`
	EndTime                  string         `json:"end_time"`
	TotalRunTime             int            `json:"total_run_time"`
	EstimatedRunTime         int            `json:"estimated_run_time"`
	EstimatedStartTimeDelta  int            `json:"estimated_start_time_delta"`
	AutoScheduled            bool           `json:"auto_scheduled"`
	InfrastructureDocumentId string         `json:"infrastructure_document_id"`
	InfrastructurePid        string         `json:"infrastructure_pid"`
	CreatedBy                map[string]any `json:"created_by"`
	WorkPoolName             string         `json:"work_pool_name"`
	WorkPoolQueueName        string         `json:"work_pool_queue_name"`
	State                    State          `json:"state"`
}
