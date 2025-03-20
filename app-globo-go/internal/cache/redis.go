package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
    "log"
    "time"
)

var client *redis.Client

func init() {
    client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis rodando na mesma máquina
    })
}

func GetCache(key string) (string, error) {
    ctx := context.Background()
    val, err := client.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", nil
    } else if err != nil {
        log.Printf("Erro ao acessar o Redis: %v", err)
        return "", err
    }
    return val, nil
}

func SetCache(key string, value string, expiration time.Duration) error {
    ctx := context.Background()
    err := client.Set(ctx, key, value, expiration).Err()
    if err != nil {
        log.Printf("Erro ao definir valor no Redis: %v", err)
        return err
    }
    return nil
}
