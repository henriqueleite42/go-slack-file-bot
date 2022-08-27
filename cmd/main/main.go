package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	godotenv.Load("../../.env")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"../../example.txt"}

	for _, fileName := range(fileArr) {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileName,
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf("Name: %s, URL: %s", file.Name, file.URL)
	}
}