package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.dbnumber"),
		PoolSize: viper.GetInt("redis.conn_pool"),
	})

	//_, err = rdb.Ping().Result()
	//if err != nil {
	//	return err
	//}
	fmt.Print(rdb)
	return nil
}

func Close()  {
	_ = rdb.Close()
}
