package component

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	sctx "github.com/linhhonblade/service-context"
	"log"
)

type sRedis struct {
	id       string
	url      string
	password string
	db       int
	client   *redis.Client
}

func NewSRedisComponent(id string) *sRedis {
	return &sRedis{
		id:  id,
		url: "",
	}
}

func (r *sRedis) ID() string {
	return r.id
}

func (r *sRedis) InitFlags() {
	flag.StringVar(
		&r.url,
		"redis-url",
		"localhost:6379",
		"URL of Redis service",
	)
	flag.StringVar(
		&r.password,
		"session-redis-password",
		"",
		"Password of session Redis service")
	flag.IntVar(
		&r.db,
		"session-redis-db",
		0,
		"Database of session Redis service")
}

func (r *sRedis) Activate(sctx sctx.ServiceContext) error {
	log.Println("...Connecting to Redis service...")
	client := redis.NewClient(r.setupConnOpts())
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis service: %v", err)
	}
	r.client = client
	log.Println("Connected to Redis service: %v", client.Options().Addr)
	return nil
}

func (r *sRedis) Stop() error {
	if r.client != nil {
		if err := r.client.Close(); err != nil {
			log.Printf("Failed to close Redis client: %v", err)
			return err
		}
		log.Println("Redis client closed successfully")
	}
	r.client = nil
	return nil
}

func (r *sRedis) setupConnOpts() *redis.Options {
	return &redis.Options{
		Addr:     r.url,
		Password: r.password,
		DB:       r.db,
	}
}

func (s *sRedis) GetRedisClient() *redis.Client {
	if s.client == nil {
		log.Fatal("Redis client is not initialized")
	}
	return s.client
}

type RedisClient interface {
	GetRedisClient() *redis.Client
}
