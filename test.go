package main

import (
	"fmt"

	"./subclient"
)

func main() {
	redis := subclient.NewRedis("192.168.1.41:6379", "123")
	pool := redis.NewPool()
	go redis.Listen(pool, "sub1")

	fmt.Println("...Lintening ...")
	go redis.SendMsg(pool, "sub1", "test")
	// for {
	// }
	select {}
}
