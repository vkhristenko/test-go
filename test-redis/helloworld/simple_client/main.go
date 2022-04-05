package main

import (
    "fmt"
    "context"

    redis "github.com/go-redis/redis/v8"
)

func Run() {
    ctx := context.Background()

    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:9999",
        Password: "",
        DB: 0,
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Print("key: ", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2: ", val2)
    }
}

func main() {
    fmt.Println("hello world!")

    Run()
}
