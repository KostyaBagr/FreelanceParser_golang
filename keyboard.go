package main
import (

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

)

func StartKeyboard() (*telego.ReplyKeyboardMarkup, error){
	keyboard := tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton("Начать парсинг"), // Column 1
		),

	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
	return keyboard, nil
}