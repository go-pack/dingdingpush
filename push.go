package dingdingpush

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AtPeople struct {
	IsAtAll   bool     `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles"`
}
type DingdingText struct {
	Content string `json:"content"`
}
type DingdingMessage struct {
	Msgtype string       `json:"msgtype"`
	At      AtPeople     `json:"at"`
	Text    DingdingText `json:"text"`
}

func Dingding(url string, msg DingdingMessage) (string, error) {

	marshal, _ := json.Marshal(msg)
	println(string(marshal))
	buf := bytes.NewBuffer(marshal)
	respon, err := http.Post(url, "application/json;charset=utf-8", buf)
	if err != nil {
		return "", err
	}
	jsonData, _ := ioutil.ReadAll(respon.Body)
	println(string(jsonData))
	return "", nil
}
