package redis

import (
	"context"
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
	"time"
)

func getConfig() *Config {
	return &Config{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	}
}
func TestNewRedisCache(t *testing.T) {
	config := getConfig()
	cache, err := NewCache(*config)
	if err != nil {
		t.Fatalf("new redis cache error: %s", err.Error())
	}
	// set a key
	err = cache.Store(context.Background(), "test", "test")
	if err != nil {
		t.Fatalf("store error: %s", err.Error())
	}
	// get a key
	_, val, err := cache.Load(context.Background(), "test")
	if err != nil {
		t.Fatalf("load error: %s", err.Error())
	}
	t.Logf("val: %s", val)
}

func TestNewRedisCacheWithPrefix(t *testing.T) {
	config := getConfig()
	config.Prefix = "temp"
	cache, err := NewCache(*config)
	if err != nil {
		t.Fatalf("new redis cache error: %s", err.Error())
	}
	// set a key
	err = cache.Store(context.Background(), "test", "test-with-prefix")
	if err != nil {
		t.Fatalf("store error: %s", err.Error())
	}
	// get a key
	_, val, err := cache.Load(context.Background(), "test")
	if err != nil {
		t.Fatalf("load error: %s", err.Error())
	}
	t.Logf("val: %s", val)

	assert.Equal(t, "test-with-prefix", val)
}

func TestNewRedisCacheWithTTL(t *testing.T) {
	config := getConfig()
	config.Prefix = "temp-"
	cache, err := NewCache(*config)
	if err != nil {
		t.Fatalf("new redis cache error: %s", err.Error())
	}
	// set a key
	err = cache.StoreEX(context.Background(), "test", "test-with-prefix", time.Second*10)
	if err != nil {
		t.Fatalf("store error: %s", err.Error())
	}
	// get a key
	_, tt, val, err := cache.LoadWithEX(context.Background(), "test")
	if err != nil {
		t.Fatalf("load error: %s", err.Error())
	}
	t.Logf("val: %s, ExpTime: %v", val, tt)

	assert.Equal(t, "test-with-prefix", val)
}
