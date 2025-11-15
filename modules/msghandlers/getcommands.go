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
				msgToSend = ("Список всех игр в Вашей коллекции:\n\n" + strings.Join(allSlice, "\n\n"))
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
	} //Конец условаия /getGames

	if elementType == "getBooks" {
		if params == "" {
			if len(commonvariables.Books) > 0 {
				allSlice := make([]string, 0, len(commonvariables.Books))
				for _, s := range commonvariables.Books {
					allSlice = append(allSlice, s.NameDesc, s.TitleDesc)
				}
				msgToSend = ("Список всех книг в Вашей коллекции:\n\n" + strings.Join(allSlice, "\n\n"))
			} else {
				msgToSend = ("Вы ещё не добавили ни одной книги :(")
			}
		} else {
			displayRange, err := strconv.Atoi(params)
			if err != nil || displayRange <= 0 { // Если есть ошибка конвертации(типа) ИЛИ число-аргумент МЕНЬШЕ ИЛИ РАВНО 0
				msgToSend = (fmt.Sprintf("В качестве аргумента к команде должно быть число больше 0\nНапример: /%s 3", elementType))
			} else {
				sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(commonvariables.Books, displayRange)
				sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
				for _, s := range sliceOfEntireCategory {
					sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
				}
				msgToSend = ("Список книг по указанному количеству (или меньше, если элементов в коллекции меньше указанного количества):\n\n" + strings.Join(sliceMsg, "\n"))
			}
		}
	} //Конец условаия /getBooks

	if elementType == "getFilms" {
		if params == "" {
			if len(commonvariables.Films) > 0 {
				allSlice := make([]string, 0, len(commonvariables.Films))
				for _, s := range commonvariables.Films {
					allSlice = append(allSlice, s.NameDesc, s.TitleDesc)
				}
				msgToSend = ("Список всех фильмов в Вашей коллекции:\n\n" + strings.Join(allSlice, "\n\n"))
			} else {
				msgToSend = ("Вы ещё не добавили ни одного фильма :(")
			}
		} else {
			displayRange, err := strconv.Atoi(params)
			if err != nil || displayRange <= 0 { // Если есть ошибка конвертации(типа) ИЛИ число-аргумент МЕНЬШЕ ИЛИ РАВНО 0
				msgToSend = (fmt.Sprintf("В качестве аргумента к команде должно быть число больше 0\nНапример: /%s 3", elementType))
			} else {
				sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(commonvariables.Films, displayRange)
				sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
				for _, s := range sliceOfEntireCategory {
					sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
				}
				msgToSend = ("Список фильмов по указанному количеству (или меньше, если элементов в коллекции меньше указанного количества):\n\n" + strings.Join(sliceMsg, "\n"))
			}
		}
	} //Конец условаия /getBooks

	return msgToSend
}
