package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TELE_TOKEN берём из окружения
	TeleToken = os.Getenv("TELE_TOKEN")
)

// tbotCmd represents the tbot command
var tbotCmd = &cobra.Command{
	Use:     "tbot",
	Aliases: []string{"start"}, // ← исправлено: []string, без ".string"
	Short:   "Start Telegram bot",
	Long: `Starts the Telegram bot using TELE_TOKEN env variable.

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// appVersion должен быть объявлен в пакете cmd:
		//   var appVersion string
		// и задаётся через -ldflags (см. ниже команды сборки)
		fmt.Printf("tbot %s started\n", appVersion) // ← Printf, а не Println с %s

		if TeleToken == "" {
			return fmt.Errorf("TELE_TOKEN is empty; export your bot token")
		}

		b, err := telebot.NewBot(telebot.Settings{
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second}, // ← исправлено: &telebot.LongPoller{...}
		})
		if err != nil {
			return fmt.Errorf("telebot init failed: %w", err)
		}

		// Хэндлер на команду /start
		b.Handle("/start", func(c telebot.Context) error {
			// Аргументы после команды: /start hello -> ["hello"]
			args := c.Args()

			// Иногда payload может приезжать через c.Message().Payload (deep-link и т.п.)
			payload := ""
			if c.Message() != nil {
				payload = c.Message().Payload
			}

			// Если первым аргументом или payload == "hello" — отвечаем
			if (len(args) > 0 && args[0] == "hello") || payload == "hello" {
				return c.Send(fmt.Sprintf("Hello I'm Tbot %s!", appVersion))
			}

			// Иначе подскажем формат
			return c.Send("Usage: /start hello")
		})

		// Простейший обработчик всего текста (опционально)
		b.Handle(telebot.OnText, func(c telebot.Context) error {
			log.Printf("message from %s: %q", c.Sender().Username, c.Text())
			return c.Send("pong")
		})

		b.Start()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tbotCmd)
}
