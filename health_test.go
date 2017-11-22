package main

import (
	"encoding/json"
	"testing"
)

func TestTypeStateDecode(t *testing.T) {
	// For the sake of brevity, the full response returned from the
	// 'List Checks in State' endpoint is omitted.
	// https://www.consul.io/api/health.html#sample-response-3
	data := []byte(`
		[
			{
				"Status": "any"
			},
			{
				"Status": "critical"
			},
			{
				"Status": "passing"
			},
			{
				"Status": "warning"
			}
		]
	`)
	var checks Checks
	if err := json.Unmarshal(data, &checks); err != nil {
		t.Error(err)
	}
	for i, expected := range []State{Any, Critical, Passing, Warning} {
		actual := checks[i].Status
		if expected != actual {
			t.Errorf("main: expected type %d, got %d instead", expected, actual)
		}
	}
}
