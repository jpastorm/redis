package bootstrap

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

func newDiscord(token string) (*discordgo.Session, error){
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return dg, errors.New("error creating Discord session, "+ err.Error())
	}

	return dg, nil
}
