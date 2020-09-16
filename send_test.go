package myslack

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	ms := NewDefaultTokenFromEnv("project1")
	ms.SendMessage(time.Now().String())
}
