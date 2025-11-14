package msghandlers

import (
	"fmt"
	"recommendtgbot/modules/botmiscfunctions"
	"recommendtgbot/modules/commonvariables"
	"strconv"
	"strings"
)

func GetElement(elementType string, params string) (msgToSend string) {
	if elementType == "getGames" {
		if params == "" {
			if len(commonvariables.Games) > 0 {
				allSlice := make([]string, 0, len(commonvariables.Games))
				for _, s := range commonvariables.Games {
					allSlice = append(allSlice, s.NameDesc, s.TitleDesc)
				}
				msgToSend = ("Список всех игр в Вашей коллекции:\n\n" + strings.Join(allSlice, "\n"))
			} else {
				msgToSend = ("Вы ещё не добавили ни одной игры :(")
			}
		} else {
			displayRange, err := strconv.Atoi(params)
			if err != nil || displayRange <= 0 { // Если есть ошибка конвертации(типа) ИЛИ число-аргумент МЕНЬШЕ ИЛИ РАВНО 0
				msgToSend = (fmt.Sprintf("В качестве аргумента к команде должно быть число больше 0\nНапример: /%s 3", elementType))
			} else {
				sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(commonvariables.Games, displayRange)
				sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
				for _, s := range sliceOfEntireCategory {
					sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
				}
				msgToSend = ("Список игр по указанному количеству (или меньше, если элементов в коллекции меньше указанного количества):\n\n" + strings.Join(sliceMsg, "\n"))
			}
		}
	}

	// if elementType == "getBooks" {
	// 	sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(commonvariables.Books, params)
	// 	sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
	// 	for _, s := range sliceOfEntireCategory {
	// 		sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
	// 	}
	// 	msgToSend = ("Список книг по указанному количеству (или меньше, если элементов меньше указанного количества):\n\n" + strings.Join(sliceMsg, "\n"))
	// 	return msgToSend
	// }
	// if elementType == "getFilms" {
	// 	sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(commonvariables.Films, params)
	// 	sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
	// 	for _, s := range sliceOfEntireCategory {
	// 		sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
	// 	}
	// 	msgToSend = ("Список фильмов по указанному количеству (или меньше, если элементов меньше указанного количества):\n\n" + strings.Join(sliceMsg, "\n"))
	// 	return msgToSend
	// }
	// return
	return msgToSend
}
