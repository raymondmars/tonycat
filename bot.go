package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"tonycat/config"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CatBot struct {
	Api *tgbotapi.BotAPI
}

func init() {
	if gin.Mode() == "release" {
		webhookUrl := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", config.BOT_TOKEN, config.WEBHOOK_URL)
		res, _ := http.Post(webhookUrl, "application/json", nil)
		body, _ := ioutil.ReadAll(res.Body)
		log.Println(webhookUrl)
		log.Println("webhook return:", string(body))
	}

}
func (c *CatBot) Chat(message *tgbotapi.Message) {
	//群聊
	atBotName := fmt.Sprintf("@%s", config.BOT_USER_NAME)
	if strings.Contains(message.Text, atBotName) {
		realTxt := strings.ReplaceAll(message.Text, atBotName, "")
		txt := (new(TuringBot)).Chat(realTxt)
		msg := tgbotapi.NewMessage(message.Chat.ID, txt)
		msg.ReplyToMessageID = message.MessageID
		// bot.GetMe()
		c.Api.Send(msg)
	} else {
		//私聊
		if message.Chat.ID == int64(message.From.ID) {
			txt := (new(TuringBot)).Chat(message.Text)
			msg := tgbotapi.NewMessage(message.Chat.ID, txt)
			c.Api.Send(msg)
		}
	}

}
