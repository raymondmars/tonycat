package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"tonycat/config"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	bot, err := tgbotapi.NewBotAPI(config.BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	if gin.Mode() == "debug" {
		bot.Debug = true
	}

	catBot := &CatBot{Api: bot}

	r.POST("bot/webhook", func(c *gin.Context) {

		body, _ := ioutil.ReadAll(c.Request.Body)
		update := tgbotapi.Update{}

		json.Unmarshal(body, &update)
		log.Printf("chat content: %v", update.Message)
		if update.Message != nil {
			catBot.Chat(update.Message)
		}
		c.String(200, "ok")
	})

	fmt.Printf("%s", r.Run(":1300"))
}
