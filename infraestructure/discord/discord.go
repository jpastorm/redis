package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jpastorm/redis/domain/user"
	"github.com/jpastorm/redis/infraestructure/request"
	"github.com/jpastorm/redis/model"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Discord struct {
	api string
	dg *discordgo.Session
	userUsecase user.UseCase
}

func NewDiscord (api string , dg *discordgo.Session, userUsecase user.UseCase) *Discord{
	return &Discord{api: api, dg: dg, userUsecase: userUsecase}
}

func (d Discord) Run() {
	// Register the messageCreate func as a callback for MessageCreate events.
	d.dg.AddHandler(d.messageCreate)

	// In this example, we only care about receiving message events.
	d.dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := d.dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	d.dg.Close()
}

func (d Discord) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "~ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "~pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "~menu" {
		message := ` 
		>>> addserver <<API>>
adduser <<email,password>>
loop <<time>>
request
		`
		s.ChannelMessageSend(m.ChannelID, message)
	}

	if m.Content == "~request" {
		data, status ,err := request.Execute(d.api)
		if err != nil {
			log.Println(err)
			s.ChannelMessageSend(m.ChannelID, status)
			return
		}
		token := fmt.Sprintf("```%s```", data["token"])
		message := `
		**` + data["email"].(string) + `**
		Token : ` + token + `
				`
		s.ChannelMessageSend(m.ChannelID, message)
	}
	//~adduser
	if strings.HasPrefix(m.Content, "~adduser") {
		userModel := new(model.User)
		userModel.GetPasswordAndName(m.Content)
		err := d.userUsecase.Create(userModel)
		if err != nil {
			log.Println(err)
		}
		s.ChannelMessageSend(m.ChannelID, "Probando")
	}
}