package discogs

import (
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"golang.org/x/time/rate"
)

func NewClient(config Config) Client {
	var cl Client

	cl = &client{
		requester: requesterLimiter{
			requester: requesterResty{
				resty: resty.New().
					SetBaseURL(config.BaseURL).
					SetTimeout(config.Timeout).
					SetHeader("User-Agent", config.UserAgent).
					SetRetryCount(config.RetryCount).
					SetRetryWaitTime(config.RetryWaitTime).
					SetRetryMaxWaitTime(config.RetryMaxWaitTime),
			},
			limiter: rate.NewLimiter(config.RateLimit, config.RateBurst),
		},
	}

	if config.FSCacheConfig.BaseDir != "" {
		cl = &clientCacheFS{
			baseDir: config.FSCacheConfig.BaseDir,
			client:  cl,
		}
	}

	if config.LRUCacheConfig.Size > 0 {
		cl = &clientCacheInMem{
			client: cl,
			lru: expirable.NewLRU[string, any](
				config.LRUCacheConfig.Size,
				nil,
				config.LRUCacheConfig.TTL,
			),
		}
	}

	return cl
}
