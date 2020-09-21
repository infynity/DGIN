package model

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// RedisConn RedisConn
var RedisConn redis.Conn
var num int
// Init Init
func Init() {
	c, err := redis.Dial("tcp", "localhost:10100")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	RedisConn = c
	//defer c.Close()
}

// SetTest SetTest
func SetTest(c redis.Conn) {
	_, err := c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

// GetTest1 GetTest1
func GetTest1() {
	c, err := redis.Dial("tcp", "127.0.0.1:10100")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		if num<10{
			fmt.Printf("Get mykey: %v \n", username)

		}
		num++
	}
}

// GetTest2 GetTest2
func GetTest2() {
	username, err := redis.String(RedisConn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		if num<10{
			fmt.Printf("Get mykey: %v \n", username)

		}
		num++
	}
}
