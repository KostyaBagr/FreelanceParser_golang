package main 

import (
	"context"
	"fmt"
	"os"


	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/joho/godotenv"
)


var activeContexts = make(map[int64]context.CancelFunc)

func CustomSendMessage(bot *telego.Bot, chatId int64, amount int) {
	// function to send messages
	data, _ := Scraper(amount)
	for _, val := range data{
		message := fmt.Sprintf(
			"Название: %s\nЦена: %s\n%s\n%s",
			val.Title,
			val.Price,
			val.Link,
			val.CreatedAt,
		)
		bot.SendMessage(tu.Messagef(
			tu.ID(chatId),
			message,
		))
}
}


func botHandlers(bh *th.BotHandler){
	// Function keeps all bot handlers.
	start_keyboard, _ := StartKeyboard()
	choose_pages, _ :=  ChoosePagesKeyboard()
	
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels start event.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID),
			"Привет %s!", update.Message.From.FirstName, 
		).WithReplyMarkup(start_keyboard))
	}, th.CommandEqual("start"))


	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handles start parsing enent.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), 
			"Давай начнем парсить!",
		).WithReplyMarkup(choose_pages))
	}, th.TextEqual("Начать парсинг"))


	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels info event.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID),
			"Бот парсит freelance.habr для получения актуальных заказов.\nНаписан на Golang, Telego разработчик - @kostya_IT",
		).WithReplyMarkup(start_keyboard))
	}, th.CommandEqual("info"))


	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels parsing 1 page.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), 
				"Ожидайте, пожалуйста",
		))
		CustomSendMessage(bot, update.Message.Chat.ID, 1)
	}, th.TextEqual("Самое свежее"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels parsing 3 pages.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), 
				"Ожидайте, пожалуйста",
		))
		CustomSendMessage(bot, update.Message.Chat.ID, 3)
	}, th.TextEqual("Последние 3 страницы"))
	

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Handels parsing 5 pages.
		bot.SendMessage(tu.Messagef(
			tu.ID(update.Message.Chat.ID), 
				"Ожидайте, пожалуйста",
		))
		CustomSendMessage(bot, update.Message.Chat.ID, 5)
	}, th.TextEqual("Последние 5 страниц"))
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