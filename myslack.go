package myslack

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

type MySlackClient struct {
	slackClient *slack.Client
	slackToken  string
	channelId   string
}

func NewDefaultTokenFromEnv(channelId string) MySlackClient {
	log.SetPrefix("[NewDefaultTokenFromEnv()]")
	defer log.SetPrefix("")
	slackToken := os.Getenv("SLACKTOKEN")
	log.Println("get key from env:", slackToken)
	ret := MySlackClient{slackToken: slackToken, channelId: channelId}
	ret.slackClient = slack.New(ret.slackToken)
	return ret
}

// func SendMessage ,成功返回true, 失败返回false
func (client *MySlackClient) SendMessage(message string) bool {
	log.SetPrefix("[SendMessage()]")
	defer log.SetPrefix("")
	//for debug
	//log.Println("slackToken:", slackToken)
	// 判断SLACKTOKEN是否为空
	// 判断CHANNELID是否为空
	//if channelId == "" {
	//	log.Println("cannot get CHANNELID from SetChannel().")
	//	return false
	//}

	channelID, _, err := client.slackClient.PostMessage(client.channelId, slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("SendMessage() failed: %s\n", err)
		return false
	}
	log.Printf("Message successfully sent to channel: %s. Message: %s\n", channelID, message)
	return true

}
