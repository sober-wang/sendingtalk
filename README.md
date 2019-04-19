# sendingtalk
The package use Golang program language send message to DingTalk
# 发送消息到钉钉
这个包用 Golang 实现了发送信息到 钉钉机器人
```shell
go get github.com/sober-wang/dingtalk
```

```go
// 发送钉钉消息
// 电话号码列表
phList := []string{"12345678901","12345678901"}
//钉钉机器人连接
robotLink := "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxx"
msgBody := sendingtalk.CreatMsgBody("Hello Wold", phList)
//发送消息的函数
// 传入 机器人连接，和消息体
sendingtalk.SendMsg(robotLink, msgBody)
```
