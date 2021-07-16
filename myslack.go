package myslack

import "github.com/slack-go/slack"

type MySlack struct {
	Token string
	//chanId string
	client *slack.Client
}

func (ms *MySlack) Init() {
	ms.client = slack.New(ms.Token)
}

func (ms MySlack) SendMsg(chanId string, msg string) error {
	_, _, ans := ms.client.PostMessage(chanId, slack.MsgOptionAsUser(true),
		slack.MsgOptionText(msg, false))
	return ans
}
