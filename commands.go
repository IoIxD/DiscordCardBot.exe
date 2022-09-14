package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// File for handling the Discord's CommandData shit

var RollCommandData = discordgo.ApplicationCommand{
	Name:        "roll",
	Description: "Use up one token to roll a card",
}

var AddCommandData = discordgo.ApplicationCommand{
	Name:        "add",
	Description: "(moderator only) Add a card",
}

var DelCommandData = discordgo.ApplicationCommand{
	Name:        "del",
	Description: "(moderator only) Delete a card",
}

var Commands = map[string]discordgo.ApplicationCommand{
	"roll": RollCommandData,
	"add":  AddCommandData,
	"del":  DelCommandData,
}

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"roll": RandomCard,
	"add":  AddCard,
	"del":  DelCard,
}

// Remove the slash commands
func RemoveSlashCommands() {
	for _, v := range discord.State.Guilds {
		registeredCommands, err := discord.ApplicationCommands(discord.State.User.ID, v.ID)
		if err != nil {
			fmt.Printf("Could not fetch registered commands: %v\n", err)
		}

		for _, n := range registeredCommands {
			err := discord.ApplicationCommandDelete(discord.State.User.ID, v.ID, n.ID)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// Refresh the slash commands
func RefreshSlashCommands() {
	for _, w := range Commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, LocalConfig.GuildID, &w)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Thread for refreshing the slash commands every minute.
func RefreshSlashCommandsThread() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			RefreshSlashCommands()
		}
	}
}
