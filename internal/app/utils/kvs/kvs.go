package kvs

import (
    "context"
    //"fmt"
    //"log"
    "time"

    "github.com/go-redis/redis/v8"

	"modak_ratelimit/internal/app/utils/logger"
    "modak_ratelimit/internal/app/entity"
)


type KvsClient interface {
    SetWithTTL(key string, expiration time.Duration, value interface{}) error
    Get(key string) (interface{}, error)
    Delete(key string) error
    Close() error
    Exists(key string) (bool, error) 
    GetExp(key string) ( time.Duration,  error)
    SetExp(key string, ttl time.Duration) error 
}

type KvsClientImpl struct {
    client *redis.Client
}


func NewClient(kvs entity.RedisDB) (KvsClient, error) {
    
    client := redis.NewClient(&redis.Options{
        Addr:     kvs.Addr,
        Password: kvs.Password, 
        DB:       kvs.Db,
    })

    ctx := context.Background()
    _, err := client.Ping(ctx).Result()
    if err != nil {
		logger.Error("Error connection to redis %v", err)
    }

    return &KvsClientImpl{client}, err
}

func (r *KvsClientImpl) SetWithTTL(key string, expiration time.Duration , value interface{}) error {
    ctx := context.Background()
    err := r.client.Set(ctx, key, value, expiration).Err()
    return err
}


func (r *KvsClientImpl) Get(key string) (interface{}, error) {
    ctx := context.Background()

    val, err := r.client.Get(ctx, key).Result()
    if err != nil {
        return "", err
    }
    
    return val, nil  
}


func (r *KvsClientImpl) GetExp(key string) ( time.Duration,  error) {

    ctx := context.Background()

    totalSeg, err := r.client.TTL(ctx, key).Result() 
    if err != nil {
        return 0, err
    }

    return totalSeg, nil  
}


func (r *KvsClientImpl) SetExp(key string, ttl time.Duration) error {

    ctx := context.Background()

    err := r.client.Expire(ctx, key, ttl).Err()
    if err != nil {
        return err
    }

    return nil  
}

func (r *KvsClientImpl) Exists(key string) (bool, error) {
    ctx := context.Background()

    status, err := r.client.Exists(ctx, key).Result()
    if err != nil {
        return false, err
    }
    if status == 0 {
        return false, nil
    }

    return true, nil  
}

func (r *KvsClientImpl) Delete(key string) error {
    ctx := context.Background()
    err := r.client.Del(ctx, key).Err()
    return err
}

func (r *KvsClientImpl) Close() error {
    return r.client.Close()
}
