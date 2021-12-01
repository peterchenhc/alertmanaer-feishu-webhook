package main

import (
	"flag"
	"net/http"

	model "github.com/chenhuaicong/alertmanaer-feishu-webhook/model"
	"github.com/chenhuaicong/alertmanaer-feishu-webhook/notifier"
	"github.com/gin-gonic/gin"
)

var (
	h            bool
	defaultRobot string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global feishu robot webhook, you can overwrite by alert rule with annotations feishuRobot")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification
		//		data, _ := ioutil.ReadAll(c.Request.Body)
		//		fmt.Println(string(data))

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = notifier.Send(notification, defaultRobot)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to feishu successful!"})

	})
	router.Run()
}
