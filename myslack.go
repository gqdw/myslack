package myslack

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
)

var SLACKTOKEN string

func init() {
	SLACKTOKEN = os.Getenv("SLACKTOKEN")

}

//func main() {
// func SendMessage ,成功返回true, 失败返回false
func SendMessage(message string) bool {
	if SLACKTOKEN == "" {
		log.Println("cannot get SLACKTOKEN from env. exit !")
		return false
	}
	api := slack.New(SLACKTOKEN)

	channelID, timestamp, err := api.PostMessage("nodenotify", slack.MsgOptionText(message, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return true

}
