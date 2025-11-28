package msghandlers

import (
	"fmt"
	"recommendtgbot/internal/commonvariables"
	"slices"
	"strconv"
)

func ClearElement(elementType string, params string) (msgToSend string) {
	if elementType == "clearAll" {
		if params == "" {
			commonvariables.Games = []commonvariables.PropertiesOfElement{}
			commonvariables.Books = []commonvariables.PropertiesOfElement{}
			commonvariables.Films = []commonvariables.PropertiesOfElement{}
			msgToSend = "Все списки успешно очищены"
		} else {
			msgToSend = (fmt.Sprintf("У команды /%s не должно быть аргументов", elementType))
		}
	} // Конец условия /clearAll

	if elementType == "clearGames" {
		if params == "" {
			commonvariables.Games = []commonvariables.PropertiesOfElement{}
			msgToSend = ("Список игр полностью очищен")
		} else {
			displayRange, err := strconv.Atoi(params)
			if err == nil && displayRange > 0 {
				if displayRange <= len(commonvariables.Games) {
					commonvariables.Games = slices.Delete(commonvariables.Games, displayRange-1, displayRange)
					msgToSend = ("Игра успешно удалена")
				} else {
					if len(commonvariables.Games) == 0 {
						msgToSend = ("Вы ещё не добавили ни одной игры :(")
					} else {
						msgToSend = ("Вам надо ввести число, не превышающее текущее количество игр (" + strconv.Itoa(len(commonvariables.Games)) + ")")
					}
				}
			} else { // Если есть ошибка конвертации(типа) ИЛИ число-аргумент МЕНЬШЕ ИЛИ РАВНО 0
				msgToSend = (fmt.Sprintf("В качестве аргумента к команде должно быть число больше 0\nНапример: /%s 3", elementType))
			}
		}
	} // Конец условия /clearGames

	if elementType == "clearBooks" {
		if params == "" {
			commonvariables.Books = []commonvariables.PropertiesOfElement{}
			msgToSend = ("Список книг полностью очищен")
		} else {
			displayRange, err := strconv.Atoi(params)
			if err == nil && displayRange > 0 {
				if displayRange <= len(commonvariables.Books) {
					commonvariables.Books = slices.Delete(commonvariables.Books, displayRange-1, displayRange)
					msgToSend = ("Книга успешно удалена")
				} else {
					if len(commonvariables.Books) == 0 {
						msgToSend = ("Вы ещё не добавили ни одной книги :(")
					} else {
						msgToSend = ("Вам надо ввести число, не превышающее текущее количество книг (" + strconv.Itoa(len(commonvariables.Books)) + ")")
					}
				}
			} else { // Если есть ошибка конвертации(типа) ИЛИ число-аргумент МЕНЬШЕ ИЛИ РАВНО 0
				msgToSend = (fmt.Sprintf("В качестве аргумента к команде должно быть число больше 0\nНапример: /%s 3", elementType))
			}
		}
	} // Конец условия /clearBooks

	if elementType == "clearFilms" {
		if params == "" {
			commonvariables.Films = []commonvariables.PropertiesOfElement{}
			msgToSend = ("Список фильмов полностью очищен")
		} else {
			displayRange, err := strconv.Atoi(params)
			if err == nil && displayRange > 0 {
				if displayRange <= len(commonvariables.Films) {
					commonvariables.Films = slices.Delete(commonvariables.Films, displayRange-1, displayRange)
					msgToSend = ("Фильм успешно удалён")
				} else {
					if len(commonvariables.Films) == 0 {
						msgToSend = ("Вы ещё не добавили ни одного фильма :(")
					} else {
						msgToSend = ("Вам надо ввести число, не превышающее текущее количество фильмов (" + strconv.Itoa(len(commonvariables.Films)) + ")")
					}
				}
			} else { // Если есть ошибка конвертации(типа) ИЛИ число-аргумент МЕНЬШЕ ИЛИ РАВНО 0
				msgToSend = (fmt.Sprintf("В качестве аргумента к команде должно быть число больше 0\nНапример: /%s 3", elementType))
			}
		}
	} // Конец условия /clearFilms

	return msgToSend
}
