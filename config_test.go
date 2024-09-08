package config

import "testing"

func TestGetString(t *testing.T) {
	table := []struct {
		Name          string
		EnvValue      string
		DefaultValues []string
		Result        string
	}{{
		Name:          "Name0",
		EnvValue:      "Env",
		DefaultValues: []string{"A", "B"},
		Result:        "Env",
	}, {
		Name:          "Name1",
		DefaultValues: []string{"A", "B"},
		Result:        "A",
	}, {
		Name:          "Name2",
		DefaultValues: []string{"", "B"},
		Result:        "B",
	}, {
		Name:          "Name3",
		DefaultValues: []string{},
		Result:        "",
	}}
	for _, item := range table {
		if item.EnvValue != "" {
			t.Setenv(item.Name, item.EnvValue)
		}
		result := GetEnvString(item.Name, item.DefaultValues...)
		if result != item.Result {
			t.Errorf("Received: %v, Expected %v", result, item.Result)
		}
	}
}
