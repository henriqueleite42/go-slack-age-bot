package commands

import "github.com/shomali11/slacker"

func RegisterCommands(bot *slacker.Slacker) {
	CalcAge(bot)
}