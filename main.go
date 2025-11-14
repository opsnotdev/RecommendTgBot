package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"recommendtgbot/modules/msghandlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	tokenFile, errRead := os.ReadFile("token.txt")
	if errRead != nil {
		fmt.Println("Failed to open token.txt file")
		os.Exit(1)
	}

	fmt.Println("Токен Вашего бота:", string(tokenFile))

	tgbot, err := tgbotapi.NewBotAPI(string(tokenFile))
	if err != nil {
		log.Println("Указан неверный токен бота")
		os.Exit(1)
	}

	log.Printf("Authorized on bot-account %s", tgbot.Self.UserName)

	updateMessage := tgbotapi.NewUpdate(0)
	updateMessage.Timeout = 10

	updates := tgbot.GetUpdatesChan(updateMessage)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message != nil { // Если бот получил новое сообщение
			if update.Message.IsCommand() {

				//Отправление сообщения на команду /start
				if update.Message.CommandWithAt() == "start" {
					outputmsg := tgbotapi.NewMessage(update.Message.Chat.ID, msghandlers.MsgStart(update.Message.CommandWithAt()))
					tgbot.Send(outputmsg)
				}

				//Добавление игры
				if update.Message.CommandWithAt()[:3] == "add" { //Определяем, что команда начинается с add
					if update.Message.CommandArguments() != "" && len(strings.SplitN(update.Message.CommandArguments(), ", ", 2)) == 2 {
						commandArgs := strings.SplitN(update.Message.CommandArguments(), ", ", 2) //Парсим аргументы к команде add
						outputmsg := tgbotapi.NewMessage(update.Message.Chat.ID, msghandlers.AddElement(update.Message.CommandWithAt(), commandArgs))
						tgbot.Send(outputmsg)
					} else {
						outputmsg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Введите имя и описание нового элемента Вашей коллекции после команды /%s\nНапример /%s Имя, Описание\n", update.Message.CommandWithAt(), update.Message.CommandWithAt()))
						tgbot.Send(outputmsg)
					}
				}

				// Чтение списков
				if update.Message.CommandWithAt()[:3] == "get" { //Определяем, что команда начинается с get
					commandArgs := update.Message.CommandArguments()
					outputmsg := tgbotapi.NewMessage(update.Message.Chat.ID, msghandlers.GetElement(update.Message.CommandWithAt(), commandArgs))
					tgbot.Send(outputmsg)
				}

				// //Чтение списка книг
				// if update.Message.CommandWithAt() == "getBooks" {
				// 	if update.Message.CommandArguments() != "" {
				// 		arguments := update.Message.CommandArguments()
				// 		displayRange, err := strconv.Atoi(arguments)
				// 		if err == nil {
				// 			sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(books, displayRange)
				// 			sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
				// 			for _, s := range sliceOfEntireCategory {
				// 				sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
				// 			}
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список книг по указанному количеству (или меньше, если элементов меньше указанного количества):\n"+strings.Join(sliceMsg, "\n\n"))
				// 			tgbot.Send(msg)
				// 		} else {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /getBooks 3")
				// 			tgbot.Send(msg)
				// 		}
				// 	} else {
				// 		if len(books) == 0 {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы ещё не добавили ни одной книги")
				// 			tgbot.Send(msg)
				// 		} else {
				// 			sliceBooks := make([]string, 0, len(books))
				// 			for _, s := range books {
				// 				sliceBooks = append(sliceBooks, s.NameDesc, s.TitleDesc)
				// 			}
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список всех добавленных вами книг:\n\n"+strings.Join(sliceBooks, "\n\n"))
				// 			tgbot.Send(msg)
				// 		}
				// 	}
				// }

				// //Чтение списка фильмов
				// if update.Message.CommandWithAt() == "getFilms" {
				// 	if update.Message.CommandArguments() != "" {
				// 		arguments := update.Message.CommandArguments()
				// 		displayRange, err := strconv.Atoi(arguments)
				// 		if err == nil {
				// 			sliceOfEntireCategory := botmiscfunctions.RandomizeSlice(films, displayRange)
				// 			sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
				// 			for _, s := range sliceOfEntireCategory {
				// 				sliceMsg = append(sliceMsg, s.NameDesc, s.TitleDesc)
				// 			}
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список фильмов по указанному количеству (или меньше, если элементов меньше указанного количества):\n"+strings.Join(sliceMsg, "\n\n"))
				// 			tgbot.Send(msg)
				// 		} else {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /getFilms 3")
				// 			tgbot.Send(msg)
				// 		}
				// 	} else {
				// 		if len(films) == 0 {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы ещё не добавили ни одного фильма")
				// 			tgbot.Send(msg)
				// 		} else {
				// 			sliceFilms := make([]string, 0, len(films))
				// 			for _, s := range films {
				// 				sliceFilms = append(sliceFilms, s.NameDesc, s.TitleDesc)
				// 			}
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список всех добавленных вами фильмов:\n\n"+strings.Join(sliceFilms, "\n\n"))
				// 			tgbot.Send(msg)
				// 		}
				// 	}
				// }

				// //Очистка всех списков /clearAll
				// if update.Message.CommandWithAt() == "clearAll" {
				// 	games = []botmiscfunctions.PropertiesOfElement{}
				// 	books = []botmiscfunctions.PropertiesOfElement{}
				// 	films = []botmiscfunctions.PropertiesOfElement{}
				// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Списки успешно очищены")
				// 	tgbot.Send(msg)
				// }

				// //Очистка списка /clearGames
				// if update.Message.CommandWithAt() == "clearGames" {
				// 	if update.Message.CommandArguments() != "" {
				// 		arguments := update.Message.CommandArguments()
				// 		displayRange, err := strconv.Atoi(arguments)
				// 		if err == nil {
				// 			if displayRange > 0 && displayRange <= len(games) {
				// 				games = slices.Delete(games, displayRange-1, displayRange)
				// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Игра успешно удалена")
				// 				tgbot.Send(msg)
				// 			} else {
				// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вам надо ввести число, не превышающее текущее количество игр ("+strconv.Itoa(len(games))+")")
				// 				tgbot.Send(msg)
				// 			}

				// 		} else {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /clearGames 3")
				// 			tgbot.Send(msg)
				// 		}
				// 	} else {
				// 		games = []botmiscfunctions.PropertiesOfElement{}
				// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Весь список игр успешно очищен")
				// 		tgbot.Send(msg)
				// 	}
				// }

				// //Очистка списка /clearBooks
				// if update.Message.CommandWithAt() == "clearBooks" {
				// 	if update.Message.CommandArguments() != "" {
				// 		arguments := update.Message.CommandArguments()
				// 		displayRange, err := strconv.Atoi(arguments)
				// 		if err == nil {
				// 			if displayRange > 0 && displayRange <= len(books) {
				// 				books = slices.Delete(books, displayRange-1, displayRange)
				// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Книга успешно удалена")
				// 				tgbot.Send(msg)
				// 			} else {
				// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вам надо ввести число, не превышающее текущее количество книг ("+strconv.Itoa(len(books))+")")
				// 				tgbot.Send(msg)
				// 			}

				// 		} else {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /clearBooks 3")
				// 			tgbot.Send(msg)
				// 		}
				// 	} else {
				// 		books = []botmiscfunctions.PropertiesOfElement{}
				// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Весь список книг успешно очищен")
				// 		tgbot.Send(msg)
				// 	}
				// }

				// //Очистка списка /clearFilms
				// if update.Message.CommandWithAt() == "clearFilms" {
				// 	if update.Message.CommandArguments() != "" {
				// 		arguments := update.Message.CommandArguments()
				// 		displayRange, err := strconv.Atoi(arguments)
				// 		if err == nil {
				// 			if displayRange > 0 && displayRange <= len(films) {
				// 				films = slices.Delete(films, displayRange-1, displayRange)
				// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фильм успешно удалён")
				// 				tgbot.Send(msg)
				// 			} else {
				// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вам надо ввести число, не превышающее текущее количество фильмов ("+strconv.Itoa(len(films))+")")
				// 				tgbot.Send(msg)
				// 			}

				// 		} else {
				// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /clearFilms 3")
				// 			tgbot.Send(msg)
				// 		}
				// 	} else {
				// 		games = []botmiscfunctions.PropertiesOfElement{}
				// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Весь список фильмов успешно очищен")
				// 		tgbot.Send(msg)
				// 	}
				// }
			} else { //Если сообщение ЭТО НЕ команда
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду, но я могу выполнять следующие команды: \n\n"+strings.Join(msghandlers.PrintAllCommands(), ""))
				tgbot.Send(msg)
			} //Конец сообщение ЭТО команда
		}
	}
}
