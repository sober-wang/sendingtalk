package sendingtalk

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type AtSet struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// 钉钉消息头
type RobotHead struct {
	MsgType string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
	At      AtSet             `json:"at"`
}

// 钉钉机器人发送的数据
type TestResult struct {
	HostName   string   `json:"HostName"`
	ProcesName []string `json:"procesName"`
	Status     bool     `json:"Status"`
}

func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

// 构建 钉钉机器人 请求体
func CreatMsgBody(msg string, phoneList []string) RobotHead {
	var robotHead RobotHead
	robotHead.MsgType = "text"
	robotHead.Text = make(map[string]string)
	robotHead.Text["content"] = msg
	robotHead.At.AtMobiles = phoneList
	robotHead.At.IsAtAll = false
	log.Printf("Message body is ==> %v", robotHead.Text["content"])

	return robotHead
}

// 发送消息的函数
func SendMsg(url string, data RobotHead) {
	//log.Println(data.ProcesName)
	d, err := json.Marshal(data)
	dropErr(err)
	// 构建一个新的请求，bytes.NewBuffer()传入[]byte 数据
	resq, err := http.NewRequest("POST", url, bytes.NewBuffer(d))
	dropErr(err)
	// 设置请求头
	resq.Header.Set("Content-Type", "application/json")

	// 定义客户端接收返回数据，如果不接受则请求不会请求成功
	client := &http.Client{}
	resp, err := client.Do(resq)
	dropErr(err)
	err = resp.Body.Close()
	dropErr(err)

}
