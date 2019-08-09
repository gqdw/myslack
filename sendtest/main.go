package main

import "github.com/gqdw/myslack"

func main() {

	ms := myslack.NewDefaultTokenFromEnv("project1")
	//myslack.SetChannel("project1")
	ms.SendMessage("test1")
}
