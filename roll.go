package main

import (
	"fmt"
	"math/rand"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

// File for commands for actually getting a card, and registering it as a card
// that the user owns.

func RandomCard() *api.InteractionResponse {
	choose := rand.Intn(256)
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content: option.NewNullableString(fmt.Sprint(choose)),
		},
	}
}
