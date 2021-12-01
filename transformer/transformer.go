package transformer

import (
	"bytes"
	"fmt"
	"time"

	"github.com/chenhuaicong/alertmanaer-feishu-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to feishu markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {
	groupname := notification.GroupLabels
	//groupKey := notification.GroupKey
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["feishuRobot"]

	var buffer bytes.Buffer

	//buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))
	buffer.WriteString(fmt.Sprintf("**%s(当前状态:%s)** \n", groupname["alertname"], status))

	buffer.WriteString(fmt.Sprintf("**告警项:** \n"))
	var bg string
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		fmt.Println(alert.StartsAt.In(loc).Format("15:04:05"))
		starttime := alert.StartsAt.In(loc).Format("15:04:05")
		endtime := alert.EndsAt.In(loc).Format("15:04:05")
		lable := alert.Labels
		if status == "firing" {
			endtime = "--------"
			if lable["severity"] == "warning" {
				bg = "orange"
			} else if lable["severity"] == "critical" {
				bg = "red"
			} else {
				bg = "green"
			}
		} else {
			bg = "green"
		}
		buffer.WriteString(fmt.Sprintf("    **开始时间：**%s\n", starttime))
		buffer.WriteString(fmt.Sprintf("    **结束时间：**%s\n", endtime))
		buffer.WriteString(fmt.Sprintf("    **告警主题：**%s\n", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("    **告警内容：**%s\n --------\n", annotations["description"]))
		//buffer.WriteString(fmt.Sprintf("\n    开始时间：%s\n", alert.StartsAt.Format("15:04:05"))

	}
	config := &model.Config{
		WideScreenMode: true,
	}
	title := &model.Title{
		Tag:     "plain_text",
		Content: fmt.Sprintf("%s当前状态:%s ", groupname["alertname"], status),
	}

	header := &model.Header{
		Title:    title,
		Template: bg,
	}
	elements := []model.Elements{
		{
			Tag:     "markdown",
			Content: buffer.String(),
		},
	}
	card := model.Card{
		Config:   config,
		Header:   header,
		Elements: elements,
	}

	markdown = &model.DingTalkMarkdown{
		MsgType: "interactive",
		Card:    card,
	}

	return
}
