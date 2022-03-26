package sthttp

import "net/http"

type TestableClient interface {
	Do(*http.Request) (*http.Response, error)
}
