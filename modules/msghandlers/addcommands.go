package msghandlers

import (
	"recommendtgbot/modules/commonvariables"
)

func AddElement(elementType string, params []string) (msgToSend string) {
	if elementType == "addGame" {
		gameEntered := commonvariables.PropertiesOfElement{NameDesc: params[0], TitleDesc: params[1]}
		commonvariables.Games = append(commonvariables.Games, gameEntered)
		msgToSend = ("Игра успешно добавлена")
	}

	if elementType == "addBook" {
		bookEntered := commonvariables.PropertiesOfElement{NameDesc: params[0], TitleDesc: params[1]}
		commonvariables.Books = append(commonvariables.Books, bookEntered)
		msgToSend = ("Книга успешно добавлена")
	}

	if elementType == "addFilm" {
		filmEntered := commonvariables.PropertiesOfElement{NameDesc: params[0], TitleDesc: params[1]}
		commonvariables.Films = append(commonvariables.Films, filmEntered)
		msgToSend = ("Фильм успешно добавлен")
	}
	return msgToSend
}
