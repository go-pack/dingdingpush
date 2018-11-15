package dingdingpush_test

import (
	"github.com/go-pack/dingdingpush"
	"testing"
)

func TestDingding(t *testing.T)  {
	msg := dingdingpush.DingdingMessage{}
	msg.Msgtype = "text"
	msg.At = dingdingpush.AtPeople{AtMobiles: []string{"13107700239"},IsAtAll:false}
	msg.Text.Content = "xx" + "\r\n" + "yy"
	dingdingpush.Dingding("https://oapi.dingtalk.com/robot/send?access_token=82137b9f44d286cb92960eea2fc6f7a25d2ed634244b2bb96f01a85b7610e787",msg)
}
