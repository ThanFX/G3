package main

import (
	"fmt"
	"log"
	"time"

	tbot "github.com/Syfaro/telegram-bot-api"
)

type Player struct {
	UserName    string
	ChatID      int64
	LastMesTime time.Time
}

var Players map[string]*Player

func getInlineKBMainScreen() tbot.InlineKeyboardMarkup {
	ikb := tbot.NewInlineKeyboardMarkup(
		tbot.NewInlineKeyboardRow(
			tbot.NewInlineKeyboardButtonData("О персонаже", "about_player"),
			tbot.NewInlineKeyboardButtonData("Карта", "map")),
		tbot.NewInlineKeyboardRow(
			tbot.NewInlineKeyboardButtonData("Навыки", "skills"),
			tbot.NewInlineKeyboardButtonData("Заметки", "keynotes")))
	return ikb
}

func startBot() {

	bot, err := tbot.NewBotAPI("839806396:AAGKWntZYsh4z1ippHIcDVWKVRy0P_ECr2o")
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %s", err)
	}
	bot.Debug = true

	Players = make(map[string]*Player)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tbot.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for upd := range updates {
		if upd.Message == nil { // ignore any non-Message Updates
			continue
		}
		if upd.Message.From.IsBot {
			fmt.Println("Боты атакуют!!")
			continue
		}

		if _, ok := Players[upd.Message.From.UserName]; ok {
			Players[upd.Message.From.UserName].LastMesTime = time.Now()
		} else {
			Players[upd.Message.From.UserName] = &Player{
				UserName:    upd.Message.From.UserName,
				ChatID:      upd.Message.Chat.ID,
				LastMesTime: time.Now()}
		}

		switch upd.Message.Command() {
		case "start":
			m := "Привет, " + upd.Message.From.FirstName
			msg := tbot.NewMessage(upd.Message.Chat.ID, m)
			msg.ReplyMarkup = getInlineKBMainScreen()
			bot.Send(msg)
		}

		log.Printf("[%s] %s", upd.Message.From.UserName, upd.Message.Text)
	}
}
