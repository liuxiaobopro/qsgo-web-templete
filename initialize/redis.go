package initialize

import (
	"qsgo-web-templete/global"

	"github.com/go-redis/redis"
)

func Redis() {
	if global.Conf.Redis.Dns != "" {
		rd := redis.NewClient(&redis.Options{
			Addr:     global.Conf.Redis.Dns,
			Password: global.Conf.Redis.Password,
			DB:       global.Conf.Redis.DB,
		})

		_, err := rd.Ping().Result()
		if err != nil {
			panic(err)
		}

		global.Redis = rd
	}
}
