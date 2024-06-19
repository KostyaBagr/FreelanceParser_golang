package main 

import (
	"fmt"
	"os"
	
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/joho/godotenv"
)


func botHandlers(bh *th.BotHandler){
	// Function keeps all bot handlers.
	keyboard, _ := StartKeyboard()
	
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels start event.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID),
			"Привет %s!", update.Message.From.FirstName, 
		).WithReplyMarkup(keyboard))
	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels info event.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), 
				`Этот бот разработан @kostya_IT на языке Golang.
			Бот парсит сайт habr freelance с заказами на тему бекенда, ботов, скриптов и т.п`,
		).WithReplyMarkup(keyboard))
	}, th.CommandEqual("info"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels parsing event.

		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), 
				"Ожидайте, пожалуйста",
		))

		data, _ := Scraper()
		for _, val := range data{
			message := fmt.Sprintf(
				"Название: %s\nЦена: %s\n%s\n%s",
				val.Title,
				val.Price,
				val.Link,
				val.CreatedAt,
			)
			bot.SendMessage(tu.Messagef(
				tu.ID(update.Message.Chat.ID),
				message,
			))
		}

		
	}, th.TextEqual("Начать парсинг"))
}



func main() {
	// main function for running bot. There is config.
	
	envFile, _ := godotenv.Read(".env")
	botToken := envFile["TOKEN"]

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	// Call handlers function 
	botHandlers(bh) 

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Start()
  
}