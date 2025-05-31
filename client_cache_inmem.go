package discogs

import (
	"context"
	"strconv"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

type clientCacheInMem struct {
	lru    *expirable.LRU[string, any]
	client Client
}

func (c *clientCacheInMem) Release(ctx context.Context, id int) (Release, error) {
	return getClientCacheInMemEntity[Release](ctx, c.lru, c.client.Release, "release", id)
}

func (c *clientCacheInMem) Artist(ctx context.Context, id int) (Artist, error) {
	return getClientCacheInMemEntity[Artist](ctx, c.lru, c.client.Artist, "artist", id)
}

func (c *clientCacheInMem) Master(ctx context.Context, id int) (Master, error) {
	return getClientCacheInMemEntity[Master](ctx, c.lru, c.client.Master, "master", id)
}

func getClientCacheInMemEntity[T any](
	ctx context.Context,
	lru *expirable.LRU[string, any],
	fn func(ctx context.Context, id int) (T, error),
	entity string, id int,
) (result T, err error) {

	key := entity + "_" + strconv.Itoa(id)
	if rawResult, ok := lru.Get(key); ok {
		return rawResult.(T), nil
	}
	result, err = fn(ctx, id)
	if err != nil {
		return result, err
	}
	lru.Add(key, result)
	return result, nil
}
