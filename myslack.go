package myslack

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

//var SLACKTOKEN string
var slackToken string

//var CHANNELID string
var channelId string
var slackClient *slack.Client

func init() {

}
func DefaultTokenFromEnv() {
	log.SetPrefix("[DefaultTokenFromEnv()]")
	slackToken := os.Getenv("SLACKTOKEN")
	slackClient = slack.New(slackToken)

}

func SetTokenFromEnv(envstr string) bool {
	log.SetPrefix("[SetTokenFromEnv()]")
	slackToken = os.Getenv(envstr)
	// 没有token就退出。
	if slackToken == "" {
		return false
	}
	slackClient = slack.New(slackToken)
	return true
}

func SetChannel(cid string) {
	channelId = cid
}

//func main() {
// func SendMessage ,成功返回true, 失败返回false
func SendMessage(message string) bool {
	log.SetPrefix("[SendMessage()]")
	//for debug
	log.Println("slackToken:", slackToken)
	// 判断SLACKTOKEN是否为空
	// 判断CHANNELID是否为空
	if channelId == "" {
		log.Println("cannot get CHANNELID from SetChannel().")
		return false
	}

	channelID, _, err := slackClient.PostMessage(channelId, slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("SendMessage() failed: %s\n", err)
		return false
	}
	//i,_ := strconv.ParseInt(timestamp,10,64)
	//tm := time.Unix(i,0)
	log.Printf("Message successfully sent to channel %s.", channelID)
	return true

}
