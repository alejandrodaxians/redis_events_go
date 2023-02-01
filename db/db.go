package db

import (
	"fmt"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)


func getRedisClient() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{	
		Addr: "localhost:6379",
	})
	if err := redisDB.Ping().Err(); err != nil {
		log.Fatalf("Couldn't connect to database. %s", err)
	}
	return redisDB
}

func getDBChannel() <-chan *redis.Message {
	redisDB := getRedisClient()
	redisDB.Do("config", "set", "notify-keyspace-events", "AKE")
	log.Info("Awaiting for changes on the database.")

	sub := redisDB.PSubscribe("__keyspace@*__:*")
	return sub.Channel()
}

func CheckForEvents() {
	channel := getDBChannel()
	for message := range channel {
		// Your logic on event goes here
		fmt.Printf("An event in the database has happened!, %v\n", message.String())
	}
}