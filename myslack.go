package myslack

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

var SLACKTOKEN string
var CHANNELID string
var SlackClient *slack.Client

func init() {
	SLACKTOKEN = os.Getenv("SLACKTOKEN")
	// 没有token就退出。
	if SLACKTOKEN == "" {
		log.Fatalln("cannot get SLACKTOKEN from env.!")
	}

	SlackClient = slack.New(SLACKTOKEN)
}

func SetChannel(channelId string) {
	CHANNELID = channelId
}

//func main() {
// func SendMessage ,成功返回true, 失败返回false
func SendMessage(message string) bool {
	// 判断SLACKTOKEN是否为空
	// 判断CHANNELID是否为空
	if CHANNELID == "" {
		log.Println("cannot get CHANNELID from SetChannel().")
		return false
	}

	channelID, timestamp, err := SlackClient.PostMessage(CHANNELID, slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("%s\n", err)
		return false
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return true

}
