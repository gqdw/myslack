package myslack

//
//import (
//	"github.com/go-kit/kit/log"
//	"os"
//	"testing"
//	"time"
//)
//
//func TestSend(t *testing.T) {
//	ms := NewDefaultTokenFromEnv("project1")
//	ms.SendMessage(time.Now().String())
//}
//
//func TestNewWithLogger(t *testing.T) {
//	logger := log.NewLogfmtLogger(os.Stdout)
//	logger = log.With(logger, "ts", log.DefaultTimestamp)
//	ms := NewWithLogger(logger, "project1")
//	ms.SendMessage(time.Now().String())
//}
