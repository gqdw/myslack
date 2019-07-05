package myslack

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

var SLACKTOKEN string
var CHANNELID string

func init() {
	SLACKTOKEN = os.Getenv("SLACKTOKEN")
}

func SetChannel(channelId string) {
	CHANNELID = channelId
}

//func main() {
// func SendMessage ,成功返回true, 失败返回false
func SendMessage(message string) bool {
	if SLACKTOKEN == "" {
		log.Println("cannot get SLACKTOKEN from env. exit !")
		return false
	}
	api := slack.New(SLACKTOKEN)

	channelID, timestamp, err := api.PostMessage(CHANNELID, slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("%s\n", err)
		return false
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return true

}
