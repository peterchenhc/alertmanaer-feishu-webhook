package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chenhuaicong/alertmanaer-feishu-webhook/model"
	"github.com/chenhuaicong/alertmanaer-feishu-webhook/transformer"
)

// Send send markdown message to feishu
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}
	//fmt.Println(markdown)
	data, err := json.Marshal(markdown)
	fmt.Println(string(data))
	if err != nil {
		return
	}

	var dingTalkRobotURL string

	if robotURL != "" {
		dingTalkRobotURL = robotURL
	} else {
		dingTalkRobotURL = defaultRobot
	}

	if len(dingTalkRobotURL) == 0 {
		return nil
	}

	req, err := http.NewRequest(
		"POST",
		dingTalkRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("feishu robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
