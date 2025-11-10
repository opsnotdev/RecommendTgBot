package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Берёт три рандомных значения из категории
func randomizeSlice(sliceCategory []string) []string {
	newSlice := []string{}
	for i := 0; i < 3; i++ {
		randomNumber := rand.Intn(len(sliceCategory))
		newSlice = append(newSlice, sliceCategory[randomNumber])
	}
	return newSlice
}

func main() {

	games := []string{"Far Cry 3", "Cyberpunk 2077", "Ведьмак 3: Дикая Охота", "Stronghold 2", "Tetris", "Эщкер фром Трахов"}
	films := []string{}
	books := []string{}

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
		if update.Message != nil { // Если бот получил новое сообщение
			if update.Message.From == nil {
				log.Println("Message has no sender")
				continue
			}
			if update.Message.Text == "" {
				log.Printf("[%s] sent a non-text message", update.Message.From.UserName)
				continue
			}

			if update.Message.Text == "/start" || update.Message.Text == "/menu" {
				//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добрый день\nЯ Ваш персональный бот\nЯ могу выполнять следующие команды:\n/addNewBook - для добавления новой книги,\n/addNewGame - для добавления новой игры,\n/addNewFilm - для добавления нового фильма\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Добрый день\nЯ Ваш персональный бот\nЯ могу выполнять следующие команды:\n/getBooks - посоветовать 3 рандомных книги,\n/getGames - посоветовать 3 рандомных игры,\n/getFilms - посоветовать 3 рандомных фильма\n")
				tgbot.Send(msg)
				continue
			}

			if update.Message.Text == "/exit" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "До встречи")
				tgbot.Send(msg)
			}

			if update.Message.Text == "/getGames" {
				if len(games) >= 3 {
					sliceCategory := randomizeSlice(games)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(sliceCategory, ", "))
					tgbot.Send(msg)
				}
			}

			if update.Message.Text == "/getBooks" {
				if len(books) >= 3 {
					sliceCategory := randomizeSlice(books)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(sliceCategory, ", "))
					tgbot.Send(msg)
				}
			}

			if update.Message.Text == "/getFilms" {
				if len(films) >= 3 {
					sliceCategory := randomizeSlice(films)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(sliceCategory, ", "))
					tgbot.Send(msg)
				}
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}
