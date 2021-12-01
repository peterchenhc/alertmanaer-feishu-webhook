package model

// type DingTalkMessage struct {
// }

// type At struct {
// 	AtMobiles []string `json:"atMobiles"`
// 	IsAtAll   bool     `json:"isAtAll"`
// }

// type DingTalkMarkdown struct {
// 	MsgType  string    `json:"msgtype"`
// 	At       *At       `json:at`
// 	Markdown *Markdown `json:"markdown"`
// }

// type Markdown struct {
// 	Title string `json:"title"`
// 	Text  string `json:"text"`
// }

type DingTalkMarkdown struct {
	//	ChatID string `json:"chat_id"`
	MsgType string `json:"msg_type"`
	Card    Card   `json:"card"`
}
type Config struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}
type Title struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}
type Header struct {
	Title    *Title  `json:"title"`
	Template string `json:"template"`
}
type Elements struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}
type Card struct {
	Config   *Config     `json:"config"`
	Header   *Header     `json:"header"`
	Elements []Elements `json:"elements"`
}
