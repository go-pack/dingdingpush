# dingdingpush
package main

import "github.com/go-pack/dingdingpush"

const url = "https://oapi.dingtalk.com/robot/send?access_token=xxx"

func main() {
	msg := dingdingpush.DingdingMessage{
		Msgtype: "text",
		At:      dingdingpush.AtPeople{AtMobiles: []string{""}},
		Text:    dingdingpush.DingdingText{Content: "通告: 测试"},
	}
	dingdingpush.Dingding(url, msg)
}
