package subclient

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type redisServer struct {
	redisHost string
	redisAuth string
}

func NewRedis(redisHost, redisAuth string) *redisServer {
	return &redisServer{redisHost, redisAuth}
}

func (rs *redisServer) NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {

			conn, err := redis.Dial("tcp", rs.redisHost)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", rs.redisAuth); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (rs *redisServer) Listen(pool *redis.Pool, key string) {
	conn := pool.Get()
	defer conn.Close()
	psc := redis.PubSubConn{conn}
	psc.Subscribe(key)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message: //有消息Publish到指定Channel时
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription: //Subscribe一个Channel时
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}

func (rs *redisServer) SendMsg(pool *redis.Pool, key, value string) {
	i := 0
	conn := pool.Get()
	defer conn.Close()
	for {
		pubtext := value + strconv.Itoa(i)
		res, err := redis.Values(conn.Do("publish", key, pubtext))
		if err != nil {

		}
		fmt.Println(res)
		time.Sleep(1000000 * time.Microsecond)
		i++

	}
}
