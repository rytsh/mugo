package request

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/liz/loader/httpx"
	"github.com/rytsh/mugo/internal/config"
	"github.com/worldline-go/logz"
)

var (
	Client *Request
	once   sync.Once
)

func New() *Request {
	once.Do(func() {
		Client = &Request{
			client: httpx.New(
				httpx.WithSkipVerify(config.App.SkipVerify),
				httpx.WithLog(logz.AdapterKV{Log: log.With().Str("source", "http").Logger()}),
			)}
	})

	return Client
}

type Request struct {
	client *httpx.Client
}

func (r *Request) Get(ctx context.Context, url string) ([]byte, error) {
	response, err := r.client.Send(
		ctx,
		url,
		http.MethodGet,
		nil, nil,
		&httpx.Retry{DisableRetry: config.App.DisableRetry},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to download url %s: %w", url, err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get: %d", response.StatusCode)
	}

	return response.Body, nil
}
