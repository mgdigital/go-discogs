package discogs

import (
	"context"
	"strconv"
)

type client struct {
	requester requester
}

func (c *client) Artist(ctx context.Context, id int) (Artist, error) {
	return clientAPIByID[Artist](
		ctx,
		c.requester,
		"artists",
		id,
	)
}

func (c *client) Release(ctx context.Context, id int) (Release, error) {
	return clientAPIByID[Release](
		ctx,
		c.requester,
		"releases",
		id,
	)
}

func (c *client) Master(ctx context.Context, id int) (Master, error) {
	return clientAPIByID[Master](
		ctx,
		c.requester,
		"masters",
		id,
	)
}

func clientAPIByID[T any](ctx context.Context, requester requester, entity string, id int) (T, error) {
	var data T
	_, err := requester.request(ctx, "/"+entity+"/"+strconv.Itoa(id), nil, &data)
	return data, err
}
