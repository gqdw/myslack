package myslack

import (
	"github.com/go-kit/kit/log"
	"github.com/nlopes/slack"
	"os"
)

type MySlackClient struct {
	slackClient *slack.Client
	slackToken  string
	channelId   string
	Logger      log.Logger
}

// New == NewDefaultTokenFromEnv
func New(channelId string) MySlackClient {
	return NewDefaultTokenFromEnv(channelId)
}

// new func NewWithLogger for pass logger into func
func NewWithLogger(logger log.Logger, channelId string) MySlackClient {
	return NewDefaultTokenFromEnvWithLogger(logger, channelId)
}

func NewDefaultTokenFromEnv(channelId string) MySlackClient {
	slackToken := os.Getenv("SLACKTOKEN")
	ret := MySlackClient{slackToken: slackToken, channelId: channelId, Logger: log.NewLogfmtLogger(os.Stdout)}
	ret.slackClient = slack.New(ret.slackToken)
	ret.Logger = log.With(ret.Logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
	ret.Logger.Log("slackToken", slackToken)
	return ret
}

func NewDefaultTokenFromEnvWithLogger(logger log.Logger, channelId string) MySlackClient {
	slackToken := os.Getenv("SLACKTOKEN")
	ret := MySlackClient{slackToken: slackToken, channelId: channelId, Logger: logger}
	ret.slackClient = slack.New(ret.slackToken)
	ret.Logger.Log("slackToken", slackToken)
	return ret
}

func NewTokenFromArg(channelId string, token string) MySlackClient {
	ret := MySlackClient{slackToken: token, channelId: channelId}
	ret.slackClient = slack.New(ret.slackToken)
	return ret
}

// func SendMessage ,成功返回true, 失败返回false
func (client *MySlackClient) SendMessage(message string) bool {

	channelID, _, err := client.slackClient.PostMessage(client.channelId, slack.MsgOptionText(message, false))
	if err != nil {
		client.Logger.Log("SendMessage() failed", err)
		return false
	}
	client.Logger.Log("Message successfully sent to channel", channelID, "Message", message)
	return true

}
