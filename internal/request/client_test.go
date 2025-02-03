package request

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRequest_Get(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello, world!"))
	}))
	defer httpServer.Close()

	client, err := New()
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	want := []byte("Hello, world!")

	got, err := client.Get(context.Background(), httpServer.URL)
	if err != nil {
		t.Errorf("Request.Get() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Request.Get() = %s, want %s", got, want)
	}
}
