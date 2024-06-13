package main
import (

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

)

func StartKeyboard() (*telego.ReplyKeyboardMarkup, error){
	keyboard := tu.Keyboard(
		tu.KeyboardRow( 
			tu.KeyboardButton("Начать парсинг"), 
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
	return keyboard, nil
}