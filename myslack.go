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
	slackToken = os.Getenv("SLACKTOKEN")
	// 没有token就退出。
	if slackToken == "" {
		log.Fatalln("cannot get SLACKTOKEN from env.!")
	}

	slackClient = slack.New(slackToken)
}

//func SetToken(token string){
//	slackToken = token
//}

func SetChannel(cid string) {
	channelId = cid
}

//func main() {
// func SendMessage ,成功返回true, 失败返回false
func SendMessage(message string) bool {
	// 判断SLACKTOKEN是否为空
	// 判断CHANNELID是否为空
	if channelId == "" {
		log.Println("cannot get CHANNELID from SetChannel().")
		return false
	}

	channelID, _, err := slackClient.PostMessage(channelId, slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("%s\n", err)
		return false
	}
	//i,_ := strconv.ParseInt(timestamp,10,64)
	//tm := time.Unix(i,0)
	log.Printf("Message successfully sent to channel %s.", channelID)
	return true

}
