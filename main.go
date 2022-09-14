package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/pelletier/go-toml/v2"
)

var c *api.Client
var g *gateway.Gateway
var AppID discord.AppID

var LocalConfig struct {
	BotToken string
}

func main() {
	// Load the config
	f, err := os.Open("config.toml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = toml.NewDecoder(f).Decode(&LocalConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	// initialize the Discord client
	c = api.NewClient("Bot " + LocalConfig.BotToken)
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	g, err = gateway.NewWithIntents(ctx, LocalConfig.BotToken, gateway.IntentGuilds|gateway.IntentGuildMessages|gateway.IntentGuildIntegrations)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Signals for terminating the program
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// wait for termination
	for {
		select {
		case <-sigs:
			return
		}
	}
}

func RefreshGuildCommands() {

}
