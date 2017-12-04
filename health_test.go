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
		[{"Status": "any"}, {"Status": "critical"}, {"Status": "passing"}, {"Status": "warning"}]
	`)
	var checks Checks
	if err := json.Unmarshal(data, &checks); err != nil {
		t.Error(err)
	}
	var actual State
	for i, expected := range []State{StateAny, StateCritical, StatePassing, StateWarning} {
		actual = checks[i].Status
		if expected != actual {
			t.Errorf("main: expected %q, got %q instead", expected, actual)
		}
	}
}

func TestTypeStateDecodeMalformed(t *testing.T) {
	// State expects type `string` not type `number`.
	data := []byte(`[{"Status": 0}]`)
	var checks Checks
	if err := json.Unmarshal(data, &checks); err == nil {
		t.Error("main: expected type string, got number instead")
	}
}

func TestTypeStateFlagParsing(t *testing.T) {
	var state State
	actual, expected := "any", StateAny.String()
	if err := state.Set(actual); err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("main: expected %q, got %q instead", expected, actual)
	}
}
