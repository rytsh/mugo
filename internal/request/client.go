package request

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/worldline-go/klient"

	"github.com/rytsh/mugo/internal/config"
)

func New() (*Request, error) {
	client, err := klient.New(
		klient.WithInsecureSkipVerify(config.App.SkipVerify),
		klient.WithLogger(slog.Default()),
		klient.WithDisableRetry(config.App.DisableRetry),
		klient.WithDisableBaseURLCheck(true),
	)
	if err != nil {
		return nil, err
	}

	r := &Request{
		client: client,
	}

	return r, nil
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

	if err := r.client.Do(req, func(r *http.Response) error {
		if err := klient.UnexpectedResponse(r); err != nil {
			return err
		}

		response, _ = io.ReadAll(r.Body)

		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
