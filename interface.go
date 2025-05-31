package discogs

import (
	"context"
)

type Client interface {
	Artist(ctx context.Context, id int) (Artist, error)
	Master(ctx context.Context, id int) (Master, error)
	Release(ctx context.Context, id int) (Release, error)
}
