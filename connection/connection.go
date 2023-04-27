package configs

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func Connection() *redis.Client {

	client := rClient()

	// check connection status
	err := ping(client)

	if err != nil {
		logrus.Fatal(err)
	}

	return client
}

func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)

	return nil
}

func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return client
}
