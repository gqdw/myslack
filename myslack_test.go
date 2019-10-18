package myslack

import "testing"

func TestNewTokenFromArg(t *testing.T) {
	sm := NewTokenFromArg("myrancher", "xoxb-524118142161-682719279095-9ZS6fe7erNYh7gptAYDBQ2tY")
	sm.SendMessage("test1")
}
