package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/pelletier/go-toml/v2"
)

var discord *discordgo.Session
var role *discordgo.Role

var LocalConfig struct {
	BotToken string
	RoleID   string
	GuildID  string
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
	discord, err = discordgo.New("Bot " + LocalConfig.BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
		// initialize slash commands
		RefreshSlashCommands()

		// initialize the verification role
		role, err = discord.State.Role(LocalConfig.GuildID, LocalConfig.RoleID)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	discord.Open()

	//RefreshSlashCommandsThread()

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
