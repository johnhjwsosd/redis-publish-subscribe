package main

import (
	"fmt"

	"./subclient"
)

func main() {
	redis := subclient.NewRedis("127.0.0.1:6379", "123")
	pool := redis.NewPool()
	go redis.Listen(pool, "sub1")

	fmt.Println("...Lintening ...")
	redis.SendMsg(pool, "sub1", "test")
	// for {
	// }

}
