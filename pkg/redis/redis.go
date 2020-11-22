package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zhangxuesong/josephblog/pkg/config"
	"log"
	"time"
)

var Redis *redis.Client

func init() {
	log.Println("连接Redis。。。")
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host + ":" + config.Config.Redis.Port,
		Password: config.Config.Redis.Auth,
		DB:       config.Config.Redis.Db,
		PoolSize: 5,
	})

	pong, err := Redis.Ping().Result()
	if err != nil {
		log.Println(pong)
		log.Panic("failed to connect redis：%v", err)
	}
	log.Println("连接Redis成功。。。")
}

// 批量向key的hash添加对应元素field的值
func BatchHashSet(client *redis.Client, key string, fields map[string]interface{}) (string, error) {
	val, err := client.HMSet(key, fields).Result()
	client.Expire(key, config.Config.Jwt.TimeOut*time.Hour)
	if err != nil {
		log.Println("Redis HMSet Error:", err)
	}
	return val, err
}

// 批量获取key的hash中对应多元素值
func BatchHashGet(client *redis.Client, key string, fields ...string) map[string]interface{} {
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := client.HGet(key, fmt.Sprintf("%s", field)).Result()
		if err == redis.Nil {
			log.Println("Key Doesn't Exists:", field)
			resMap[field] = result
		} else if err != nil {
			log.Println("Redis HMGet Error:", err)
			resMap[field] = result
		}
		if val != "" {
			resMap[field] = val
		} else {
			resMap[field] = result
		}
	}
	return resMap
}
