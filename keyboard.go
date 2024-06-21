package main
import (

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

)

func StartKeyboard() (*telego.ReplyKeyboardMarkup, error){
	// Keyboard for star.
	keyboard := tu.Keyboard(
		tu.KeyboardRow( 
			tu.KeyboardButton("Начать парсинг"), 
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
	return keyboard, nil
}


func ChoosePagesKeyboard() (*telego.ReplyKeyboardMarkup, error){
	// Keyboard for choosing pages amount to parsing.
	keyboard := tu.Keyboard(
		tu.KeyboardRow( 
			tu.KeyboardButton("Самое свежее"), 
			tu.KeyboardButton("Последние 3 страницы"),
			tu.KeyboardButton("Последние 5 страниц"),  
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
	return keyboard, nil
}
