package cache

import (
	"context"
	"errors"
	"time"
)

type NoopCache struct{}

func (n *NoopCache) Get(ctx context.Context, key string, dest interface{}) error {
	return errors.New("reddis disable!")
}

func (n *NoopCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return errors.New("reddis disable!")
}

func (n *NoopCache) Delete(ctx context.Context, key string) error {
	return errors.New("reddis disable!")
}
