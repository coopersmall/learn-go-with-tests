package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	data := "hello, world"

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &StoreDouble{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &ResponseWriterDouble{}

		srv.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		store := &StoreDouble{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", wanted "%s"`, response.Body.String(), data)
		}
	})
}
