package bootstrap

import (
	"github.com/jpastorm/redis/domain/user"
	discordSession "github.com/jpastorm/redis/infraestructure/discord"
	userStorage "github.com/jpastorm/redis/infraestructure/redis/user"
)

func Run() error{
	token, api, err := loadConfiguration()
	if err != nil {
		return err
	}
	pool := newPoolRedis()
	db := pool.Get()
	//defer redis.Close()
	dg, err := newDiscord(token)
	if err != nil {
		return err
	}
	userUsecase := user.New(userStorage.New(db))
	discord := discordSession.NewDiscord(api,dg, userUsecase)
	discord.Run()

	return nil
}
