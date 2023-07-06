package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
)

func SendDingDing(token, title, fileUrl string) string {
	// 钉钉通知
	dingClient := dingtalk.NewClient(token, "")
	text := fmt.Sprintf(`### %s
	---
	- 上传时间: %s
	- 应用: %s
	- podIP: %s
	- 文件路径: %s`, title, time.Now().Format("2006/01/02 15:04:05"), os.Getenv("APPNAME"), os.Getenv("KUBERNETES_POD_IP"), fmt.Sprintf("[点我下载](%s)", fileUrl))
	msg := dingtalk.NewMarkdownMessage().SetMarkdown(title, text)
	_, res, _ := dingClient.Send(msg)
	return res.ErrMsg
}
