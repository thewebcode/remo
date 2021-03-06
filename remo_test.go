package remo

import "testing"

var config Config

func init() {
	config = Config{
		Slack: SlackConfig{
			Token:      "dummy-token",
			MainPostChannel: "#random",
		},
	}
}

func TestNew(t *testing.T) {
	remo := New(&config)
	if remo == nil {
		t.Fatal("remo.New() should be returned pointer of Remo{}")
	}
}
