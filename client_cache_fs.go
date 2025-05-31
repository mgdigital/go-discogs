package discogs

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path"
	"strconv"
)

type clientCacheFS struct {
	client  Client
	baseDir string
}

func (c *clientCacheFS) Release(ctx context.Context, id int) (Release, error) {
	return getClientCacheFSEntity[Release](ctx, c.baseDir, "release", id, c.client.Release)
}

func (c *clientCacheFS) Artist(ctx context.Context, id int) (Artist, error) {
	return getClientCacheFSEntity[Artist](ctx, c.baseDir, "artist", id, c.client.Artist)
}

func (c *clientCacheFS) Master(ctx context.Context, id int) (Master, error) {
	return getClientCacheFSEntity[Master](ctx, c.baseDir, "master", id, c.client.Master)
}

func getClientCacheFSEntity[T any](
	ctx context.Context,
	baseDir, name string, id int,
	fn func(ctx context.Context, id int) (T, error),
) (T, error) {
	var (
		bytes  []byte
		record T
		err    error
	)

	filePath := path.Join(baseDir, name, strconv.Itoa(id)+".json")
	bytes, err = os.ReadFile(filePath)

	if errors.Is(err, os.ErrNotExist) {
		if record, err = fn(ctx, id); err == nil {
			if bytes, err = json.MarshalIndent(record, "", "  "); err == nil {
				if err = os.MkdirAll(path.Dir(filePath), fs.ModeDir|fs.ModePerm); err == nil {
					err = os.WriteFile(filePath, bytes, 0777)
				}
			}
		}
	} else if err == nil {
		err = json.Unmarshal(bytes, &record)
	}

	return record, err
}
