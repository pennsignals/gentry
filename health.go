package main

type State string

const (
	StateAny      State = "any"
	StateCritical State = "critical"
	StatePassing  State = "passing"
	StateWarning  State = "warning"
)

type Check struct {
	ID          string   `json:"CheckID"`
	Name        string   `json:"Name"`
	Node        string   `json:"Node"`
	Notes       string   `json:"Notes"`
	Output      string   `json:"Output"`
	ServiceID   string   `json:"ServiceID"`
	ServiceName string   `json:"ServiceName"`
	ServiceTags []string `json:"ServiceTags"`
	Status      State    `json:"Status"`
}

type Checks []*Check
