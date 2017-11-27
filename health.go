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
		*s = Critical
	case "passing":
		*s = Passing
	case "warning":
		*s = Warning
	default:
		*s = Any
	}
	return nil
}

const (
	Any State = iota
	Critical
	Passing
	Warning
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
