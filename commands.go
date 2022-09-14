package main

import (
	"fmt"
	"os"
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
)

// File for handling the Discord's CommandData shit

var RollCommandData = api.CreateCommandData{
	Name:        "roll",
	Description: "Use up one token to roll a card",
}

var Commands = []api.CreateCommandData{
	RollCommandData,
}

var CommandMap map[string]func() *api.InteractionResponse

func init() {
	// Sometimes the compiler isn't smart enough to know that
	// this function should go after main.go. If this is the case
	// then start a new thread while we wait for it to client to be a thing.
	if c == nil {
		go func() {
			for c == nil {
			}
			InitFunction()
		}()
	} else {
		InitFunction()
	}

}

func InitFunction() {
	CommandMap = make(map[string]func() *api.InteractionResponse)
	CommandMap["roll"] = RandomCard
	RefreshCommands()
	go RefreshCommandThread()
}
func RefreshCommandThread() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			RefreshCommands()
		}
	}
}

func RefreshCommands() {
	_, err := c.BulkOverwriteCommands(AppID, Commands)
	if err != nil {
		fmt.Printf("Error while refreshing commands: %v\n", err.Error())
		os.Exit(1)
	}

}

type handler struct {
	*state.State
}

func (h handler) HandleInteraction(ev *discord.InteractionEvent) *api.InteractionResponse {
	switch data := ev.Data.(type) {
	case *discord.CommandInteraction:
		cmd, ok := CommandMap[data.Name]
		if ok {
			return cmd()
		}
	}
	return nil
}
