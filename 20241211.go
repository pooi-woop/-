package main

import (
"context"
"fmt"
"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}

	user, err := rdb.Get(ctx, name).Result() // 使用上下文和检查错误
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key not found")
		} else {
			fmt.Printf("Error getting value: %v\n", err)
		}
	} else {
		fmt.Println(user)
	}
}
