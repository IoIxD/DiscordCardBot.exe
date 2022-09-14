package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// File for commands for actually getting a card, and registering it as a card
// that the user owns.

// Give the user a random card from the database
func RandomCard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Get the user's roles.
	roles := i.Member.Roles
	valid := false
	for _, v := range roles {
		if v == LocalConfig.RoleID {
			valid = true
		}
	}

	if !valid {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You are not authorized to perform this command",
			},
		})
		return
	}

	rand.Intn(int(time.Now().Unix()))
	choose := rand.Intn(256)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprint(choose),
		},
	})
}

// Add a card
func AddCard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO
}

// Add a card
func DelCard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO
}
