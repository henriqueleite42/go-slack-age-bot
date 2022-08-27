package main

import (
	"context"
	"go-slack-age-bot/pkg/commands"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load("../../.env")

	bot := slacker.NewClient(
		os.Getenv("SLACK_BOT_TOKEN"),
		os.Getenv("SLACK_APP_TOKEN"),
	)

	commands.RegisterCommands(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}