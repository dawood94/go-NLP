package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Edw590/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warnung: Konnte .env Datei nicht laden: %v", err)
	}

	//create a new bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	// creat WitAi client : for NLP to understand human language ( user input ).
	// can extract structured dara fron text or voice input.
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	// create a new wolfram client
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	// go routine to print command events
	go printCommandEvents(bot.CommandEvents())

	// add a command : welche Aufgabe macht ein Bot

	bot.Command("My question is - <message>", &slacker.CommandDefinition{
		Description: "send any question to wolfram",
		Examples:    []string{"what is the capital of germany?"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")
			// send the query to witai
			msg, err := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			if err != nil {
				response.Reply(fmt.Sprintf("Error parsing message: %v", err))
				return
			}

			log.Printf("Received message: %v", msg)

			// we need to convert the meg to json so that Wolfram can understand it .
			// the Idea is to catch the Value of the intent and send it to Wolfram.
			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])

			//catch the value by gjson package
			value := gjson.Get(rough, "entities.wolfram_search_query.0.value") // .0.value means the first value of the array then the value .
			answer := value.String()
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				response.Reply(fmt.Sprintf("Error getting answer from wolfram: %v", err))
				return
			}
			fmt.Println(value)

			//response.Reply("received your question")
			if msg != nil && msg.Text != "" {
				response.Reply(fmt.Sprintf("Wolfram antwortet: %s", res))
			} else {
				response.Reply("Ich konnte deine Frage leider nicht verstehen.")
			}
		},
	})

	// cancel the Bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
