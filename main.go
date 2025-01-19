package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
)

type User struct {
	Age  int64
	Name string
}

func main() {
	ctx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:7000"},
		ReplicaOnly: true,
	})
	if err != nil {
		panic(err)
	}
	adapter := rueidiscompat.NewAdapter(client)

	user := &User{Name: "k-jun", Age: 27}
	x, err := adapter.Cache(5*time.Second).Get(ctx, "user").Result()
	fmt.Println("x:", x, "err:", err)
	adapter.Set(ctx, "user", user, 60*time.Second).Result()
	time.Sleep(1 * time.Millisecond)
	x, err = adapter.Cache(5*time.Second).Get(ctx, "user").Result()
	fmt.Println("x:", x, "err:", err)
	adapter.Del(ctx, "user")
}
