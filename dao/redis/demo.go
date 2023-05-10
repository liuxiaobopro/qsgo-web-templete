package redis

import (
	"qsgo-web-templete/global"
)

/*
redis参考链接:
	http://blogs.liuxiaobo.net.cn/2023/05/10/go%e7%ae%80%e5%8d%95%e6%93%8d%e4%bd%9credis/
	https://github.com/redis/go-redis
*/

func SetString() error {
	// set操作：第三个参数是过期时间，如果是0表示不会过期。
	err := global.Redis.Set("k1", "v1", 0).Err()
	if err != nil {
		global.ZapS.Error(err)
		return err
	}
	return nil
}

func GetString() (string, error) {
	// get操作
	val, err := global.Redis.Get("k1").Result()
	if err != nil {
		global.ZapS.Error(err)
		return "", err
	}
	return val, nil
}

func GetStringNotFound() (string, error) {
	// get获取一个不存在的key，err会返回redis.Nil，因此要注意判断err
	val2, err := global.Redis.Get("k2").Result()
	if err != nil {
		global.ZapS.Error(err)
		return "", err
	}
	return val2, nil
}

func SetHash() error {
	// hset操作
	err := global.Redis.HSet("hash1", "k1", "v1").Err()
	if err != nil {
		global.ZapS.Error(err)
		return err
	}
	return nil
}
