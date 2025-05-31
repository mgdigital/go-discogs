package discogs

import (
	"errors"
	"fmt"
)

var (
	ErrDiscogsClient = errors.New("discogs client")
	ErrNotFound      = fmt.Errorf("%w: not found", ErrDiscogsClient)
)
