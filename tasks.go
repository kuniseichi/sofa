package sofa

// Task structs contain information about a single task on the server. Examples
// of tasks represented include database compaction & indexing.
type Task struct {
	ChangesDone  int64  `json:"changes_done"`
	Database     string `json:"database"`
	PID          string `json:"pid"`
	Progress     int64  `json:"progress"`
	StartedOn    int64  `json:"started_on"`
	TotalChanges int64  `json:"total_changes"`
	Type         string `json:"type"`
	UpdatedOn    int64  `json:"updated_on"`
}
