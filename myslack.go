package main

import "github.com/slack-go/slack"

type MySlack struct {
	token string
	//chanId string
	client *slack.Client
}

func (ms *MySlack) Init() {
	ms.client = slack.New(ms.token)
}

func (ms MySlack) SendMsg(chanId string, msg string) {
	ms.client.PostMessage(chanId, slack.MsgOptionAsUser(true),
		slack.MsgOptionText(msg, false))
}
