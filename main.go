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

				//Добавление элемента
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

				//Очистка списков
				if update.Message.CommandWithAt()[:5] == "clear" { //Определяем, что команда начинается с clear
					commandArgs := update.Message.CommandArguments()
					outputmsg := tgbotapi.NewMessage(update.Message.Chat.ID, msghandlers.ClearElement(update.Message.CommandWithAt(), commandArgs))
					tgbot.Send(outputmsg)
				}

			} else { //Если сообщение ЭТО НЕ команда
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду, но я могу выполнять следующие команды: \n\n"+strings.Join(msghandlers.PrintAllCommands(), ""))
				tgbot.Send(msg)
			} //Конец сообщение ЭТО команда
		}
	}
}
