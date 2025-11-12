package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type botCommands struct {
	command            string
	commandDescription string
}

type propertiesOfElement struct {
	nameDesc  string
	titleDesc string
}

// Берёт рандомные значения из категории
func randomizeSlice(sliceCategory []propertiesOfElement, displayRange int) []propertiesOfElement {
	newSlice := []propertiesOfElement{}
	if displayRange > len(sliceCategory) {
		displayRange = len(sliceCategory)
	}
	for len(newSlice) < displayRange {
		randomNumber := rand.Intn(len(sliceCategory))
		if !slices.Contains(newSlice, sliceCategory[randomNumber]) {
			newSlice = append(newSlice, sliceCategory[randomNumber])
		}
	}
	return newSlice
}

func main() {

	commands := []botCommands{
		{command: "/addGame", commandDescription: " - добавить новую игру в Ваш список игр в формате: \n\"Игра, Описание игры\"\nПример команды: /addGame Spyro, Игра про дракончика\n\n"},
		{command: "/addBook", commandDescription: " - добавить новую книгу в Вашу библиотеку в формате: \n\"Книга, Описание книги\"\nПример команды: /addBook 1984, Роман-антиутопия\n\n"},
		{command: "/addFilm", commandDescription: " - добавить новый фильм в Вашу фильмотеку в формате: \n\"Фильм, Описание фильма\"\nПример команды: /addFilm Джонни Мнемоник, фантастика\n\n"},
		{command: "/getGames", commandDescription: " - посоветовать рандомную игру из Вашего списка.\nПример команды: /getGames \nПример команды, если хотите вывести заданное количество игр: /getGames 3\n\n"},
		{command: "/getBooks", commandDescription: " - посоветовать рандомную книгу из Вашего списка.\nПример команды: /getBooks \nПример команды, если хотите вывести заданное количество игр: /getBooks 3\n\n"},
		{command: "/getFilms", commandDescription: " - посоветовать рандомный фильм из Вашего списка.\nПример команды: /getFilms \nПример команды, если хотите вывести заданное количество игр: /getFilms 3\n\n"},
	}

	commandsSlice := make([]string, 0, len(commands))
	for _, s := range commands {
		commandsSlice = append(commandsSlice, s.command, s.commandDescription)
	}

	//Список массивов, доступных к заполнению, хранению и выводу
	games := []propertiesOfElement{}
	books := []propertiesOfElement{}
	films := []propertiesOfElement{}

	tokenFile, errRead := os.ReadFile("token.txt")
	if errRead != nil {
		fmt.Println("Failed to open token.txt file")
		os.Exit(1)
	}

	fmt.Println("Токен Вашего бота:", string(tokenFile))

	tgbot, err := tgbotapi.NewBotAPI(string(tokenFile))
	if err != nil {
		log.Println("Указан неверный токен бота")
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

				if update.Message.CommandWithAt() == "start" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я твой помощник-напоминалка-запоминалка. Я могу выполнять следующие команды: \n\n"+strings.Join(commandsSlice, ""))
					tgbot.Send(msg)
				}

				//Добавление игры
				if update.Message.CommandWithAt() == "addGame" {
					if update.Message.CommandArguments() != "" {
						gameDescription := strings.SplitN(update.Message.CommandArguments(), ", ", 2)
						gameEntered := propertiesOfElement{nameDesc: gameDescription[0], titleDesc: gameDescription[1]}
						games = append(games, gameEntered)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Игра успешно добавлена")
						tgbot.Send(msg)
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите имя и описание игры после команды. Например:\n/addGame Spyro, Игра про дракончика")
						tgbot.Send(msg)
					}
				}

				//Добавление книги
				if update.Message.CommandWithAt() == "addBook" {
					if update.Message.CommandArguments() != "" {
						bookDescription := strings.SplitN(update.Message.CommandArguments(), ", ", 2)
						bookEntered := propertiesOfElement{nameDesc: bookDescription[0], titleDesc: bookDescription[1]}
						books = append(books, bookEntered)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Книга успешно добавлена")
						tgbot.Send(msg)
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите имя и описание книги после команды. Например:\n/addBook 1984, Роман-антиутопия")
						tgbot.Send(msg)
					}
				}

				//Добавление фильма
				if update.Message.CommandWithAt() == "addFilm" {
					if update.Message.CommandArguments() != "" {
						filmDescription := strings.SplitN(update.Message.CommandArguments(), ", ", 2)
						filmEntered := propertiesOfElement{nameDesc: filmDescription[0], titleDesc: filmDescription[1]}
						films = append(films, filmEntered)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фильм успешно добавлен")
						tgbot.Send(msg)
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите имя и описание фильма после команды. Например:\n/addFilm Джонни Мнемоник, фантастика")
						tgbot.Send(msg)
					}
				}

				//Чтение списка игр
				if update.Message.CommandWithAt() == "getGames" {
					if update.Message.CommandArguments() != "" {
						arguments := update.Message.CommandArguments()
						displayRange, err := strconv.Atoi(arguments)
						if err == nil {
							sliceOfEntireCategory := randomizeSlice(games, displayRange)
							sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
							for _, s := range sliceOfEntireCategory {
								sliceMsg = append(sliceMsg, s.nameDesc, s.titleDesc)
							}
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список игр по указанному количеству (или меньше, если элементов меньше указанного количества):\n\n"+strings.Join(sliceMsg, "\n"))
							tgbot.Send(msg)
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /getGames 3")
							tgbot.Send(msg)
						}
					} else {
						if len(games) == 0 {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы ещё не добавили ни одной игры")
							tgbot.Send(msg)
						} else {
							sliceGames := make([]string, 0, len(games))
							for _, s := range games {
								sliceGames = append(sliceGames, s.nameDesc, s.titleDesc)
							}
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список всех добавленных вами игр:\n\n"+strings.Join(sliceGames, "\n\n"))
							tgbot.Send(msg)
						}
					}
				}

				//Чтение списка книг
				if update.Message.CommandWithAt() == "getBooks" {
					if update.Message.CommandArguments() != "" {
						arguments := update.Message.CommandArguments()
						displayRange, err := strconv.Atoi(arguments)
						if err == nil {
							sliceOfEntireCategory := randomizeSlice(books, displayRange)
							sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
							for _, s := range sliceOfEntireCategory {
								sliceMsg = append(sliceMsg, s.nameDesc, s.titleDesc)
							}
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список книг по указанному количеству (или меньше, если элементов меньше указанного количества):\n"+strings.Join(sliceMsg, "\n\n"))
							tgbot.Send(msg)
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /getBooks 3")
							tgbot.Send(msg)
						}
					} else {
						if len(books) == 0 {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы ещё не добавили ни одной книги")
							tgbot.Send(msg)
						} else {
							sliceBooks := make([]string, 0, len(books))
							for _, s := range books {
								sliceBooks = append(sliceBooks, s.nameDesc, s.titleDesc)
							}
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список всех добавленных вами книг:\n\n"+strings.Join(sliceBooks, "\n\n"))
							tgbot.Send(msg)
						}
					}
				}

				//Чтение списка фильмов
				if update.Message.CommandWithAt() == "getFilms" {
					if update.Message.CommandArguments() != "" {
						arguments := update.Message.CommandArguments()
						displayRange, err := strconv.Atoi(arguments)
						if err == nil {
							sliceOfEntireCategory := randomizeSlice(films, displayRange)
							sliceMsg := make([]string, 0, len(sliceOfEntireCategory))
							for _, s := range sliceOfEntireCategory {
								sliceMsg = append(sliceMsg, s.nameDesc, s.titleDesc)
							}
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список фильмов по указанному количеству (или меньше, если элементов меньше указанного количества):\n"+strings.Join(sliceMsg, "\n\n"))
							tgbot.Send(msg)
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В качестве аргумента к команде должно быть число.\nНапример: /getFilms 3")
							tgbot.Send(msg)
						}
					} else {
						if len(films) == 0 {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы ещё не добавили ни одного фильма")
							tgbot.Send(msg)
						} else {
							sliceFilms := make([]string, 0, len(films))
							for _, s := range films {
								sliceFilms = append(sliceFilms, s.nameDesc, s.titleDesc)
							}
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список всех добавленных вами фильмов:\n\n"+strings.Join(sliceFilms, "\n\n"))
							tgbot.Send(msg)
						}
					}
				}
			} //Конец условия if message IsCommand

			if !update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду, но я могу выполнять следующие команды: \n\n"+strings.Join(commandsSlice, ""))
				tgbot.Send(msg)
			} //Конец условия if message !IsCommand
		}
	}
}
