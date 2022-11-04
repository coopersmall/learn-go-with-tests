package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"
	"time"
)

type StoreDouble struct {
	response string
	t        *testing.T
}

func (s *StoreDouble) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type ResponseWriterDouble struct {
	written bool
}

func (s *ResponseWriterDouble) Header() http.Header {
	s.written = true
	return nil
}

func (s *ResponseWriterDouble) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *ResponseWriterDouble) WriteHeader(statusCode int) {
	s.written = true
}
