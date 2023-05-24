package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/mugo/internal/config"
	"github.com/worldline-go/utility/httpx"
)

var (
	Client *Request
	once   sync.Once
)

func New() *Request {
	once.Do(func() {
		client, _ := httpx.NewClient(
			httpx.WithInsecureSkipVerify(config.App.SkipVerify),
			httpx.WithZerologLogger(log.With().Str("source", "http").Logger()),
			httpx.WithDisableRetry(config.App.DisableRetry),
			httpx.WithDisableBaseURLCheck(true),
		)
		Client = &Request{
			client: client,
		}
	})

	return Client
}

type Request struct {
	client *httpx.Client
}

func (r *Request) Get(ctx context.Context, url string) ([]byte, error) {
	var response []byte
	if err := r.client.RequestWithURL(
		ctx,
		http.MethodGet,
		url,
		nil, nil,
		func(r *http.Response) error {
			if r.StatusCode < 200 || r.StatusCode >= 300 {
				return fmt.Errorf("failed to download url %s: %s", url, r.Status)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}

			response = body

			return nil
		},
	); err != nil {
		return nil, fmt.Errorf("failed to download url %s: %w", url, err)
	}

	return response, nil
}
