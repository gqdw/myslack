package main

import (
	"github.com/gqdw/myslack"
	"time"
)

func main() {

	ms := myslack.NewDefaultTokenFromEnv("project1")
	ms.SendMessage(time.Now().String())
}
