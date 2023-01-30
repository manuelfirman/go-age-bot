package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Comand Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4570247641910-4573929313829-7l9UlUFC0MJgqNzvVDdH6hUM")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04GYMUS6LT-4600554739360-1948c14d349133df5d5a01b8ea5f4ab5c9938c2de876a59163c4b4b93be823c8")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("age-bot naci en <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		//Examples: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				response.Reply("Error. Intenta con 'age-bot naci en <año de nacimiento>'")
				return
			}
			if yob > 2016 {
				response.Reply("No me mientas")
				return
			}
			if yob < 1922 {
				response.Reply("Estas muerto!")
				return
			}

			age := 2022 - yob
			r := fmt.Sprintf("Tu edad es %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
