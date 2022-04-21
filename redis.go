package syservice

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func NewRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     yamlConfig.Redis.Host,
		Password: yamlConfig.Redis.Auth,
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pong, "redis success.")
}

// 向redis中添加消息
func AddMessage(data string) {
	if rdb == nil {
		NewRedis()
	}

	value, err := json.Marshal(data)
	if err != nil {
		log.Println("生成消息出错：", err)
		return
	}
	log.Println("生成消息：", string(value))

	// 发布消息
	err = rdb.LPush(yamlConfig.Redis.TWMsgMQKey, string(value)).Err()
	if err != nil {
		log.Println("发布消息出错：", err)
		return
	}
}
