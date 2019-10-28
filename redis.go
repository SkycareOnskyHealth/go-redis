package redis

import (
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"time"
)

func NewRedis(options interface{}) *Redis {
	opts := new(redis.Options)
	Merge(opts, options)
	client := redis.NewClient(opts)
	return &Redis{
		db: client,
	}
}

type Redis struct {
	db       *redis.Client
	pipeline *redis.Pipeline
}

func (r *Redis) DB() *redis.Client {
	return r.db
}

func (r *Redis) Close() error {
	return r.db.Close()
}
func (r *Redis) Ping() *redis.StatusCmd {
	return r.db.Ping()
}
func (r *Redis) Del(keys ...string) error {
	_, err := r.db.Del(keys...).Result()
	return err
}
func (r *Redis) HDel(key string, fields ...string) error {
	_, err := r.db.HDel(key, fields...).Result()
	return err
}
func (r *Redis) Get(key string) (string, error) {
	return r.db.Get(key).Result()
}
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.db.Set(key, value, expiration).Result()
}
func (r *Redis) SetObject(objectKey string, field string, value interface{}) (bool, error) {
	bytes, err := msgpack.Marshal(value)
	if err != nil {
		return false, err
	}
	return r.db.HSet(objectKey, field, bytes).Result()
}
func (r *Redis) GetObject(objectKey string, field string, result interface{}) error {
	temp, err := r.db.HGet(objectKey, field).Bytes()
	if err != nil {
		return err
	}
	err = msgpack.Unmarshal(temp, result)
	return err
}

// CheckExistedObject will return true if the object is existed.
func (r *Redis) CheckExistedObject(objectKey string, field string) (bool, error) {
	existed, err := r.db.HExists(objectKey, field).Result()
	if err != nil {
		return false, err
	}
	if existed == true {
		return true, nil
	}else {
		return false, nil
	}
}
