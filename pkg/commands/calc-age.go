package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func CalcAge(bot *slacker.Slacker) {
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := time.Now().UTC().Year() - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})
}