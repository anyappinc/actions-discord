package services

import (
	"github.com/anyappinc/actions-discord/app/config"
	"github.com/anyappinc/actions-discord/app/services/components"
)

func SendMessage() error {
	// messageAppを特定
	messageApp := config.GlobalConfig.MessageApp
	if messageApp == "discord" {
		var discordComponent components.DiscordComponent
		// Make discord request
		discordComponent.MakeRequest()
		err := discordComponent.SendRequest()
		if err != nil {
			return err
		}
	}
	if messageApp == "slack" {
		// Slackにメッセージを送信
		// components.SlackClient.SendMessage()
	}

	return nil
}
