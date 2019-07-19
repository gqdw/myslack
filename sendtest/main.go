package main

import "github.com/gqdw/myslack"

func main() {

	myslack.SetChannel("project1")
	myslack.SendMessage("test1")
}
