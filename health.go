package main

import (
	"encoding/json"
)

type State int

func (s *State) UnmarshalJSON(data []byte) error {
	var state string
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}
	switch state {
	case "critical":
		*s = StateCritical
	case "passing":
		*s = StatePassing
	case "warning":
		*s = StateWarning
	default:
		*s = StateAny
	}
	return nil
}

const (
	StateAny State = iota
	StateCritical
	StatePassing
	StateWarning
)

type Check struct {
	ID          string `json:"CheckID"`
	Name        string `json:"Name"`
	Node        string `json:"Node"`
	Notes       string `json:"Notes"`
	Output      string `json:"Output"`
	ServiceID   string `json:"ServiceID"`
	ServiceName string `json:"ServiceName"`
	ServiceTags string `json:"ServiceTags"`
	Status      State  `json:"Status"`
}

type Checks []*Check
