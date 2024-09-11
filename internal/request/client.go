package request

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"

	"github.com/rytsh/mugo/internal/config"
	"github.com/worldline-go/klient"
)

var (
	Client *Request
	once   sync.Once
)

func New() *Request {
	once.Do(func() {
		client, _ := klient.New(
			klient.WithInsecureSkipVerify(config.App.SkipVerify),
			klient.WithLogger(slog.Default()),
			klient.WithDisableRetry(config.App.DisableRetry),
			klient.WithDisableBaseURLCheck(true),
		)
		Client = &Request{
			client: client,
		}
	})

	return Client
}

type Request struct {
	client *klient.Client
}

func (r *Request) Get(ctx context.Context, url string) ([]byte, error) {
	var response []byte
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if err := r.client.Do(
		req,
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
		return nil, err
	}

	return response, nil
}
