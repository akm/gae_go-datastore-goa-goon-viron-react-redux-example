package app

import(
	"io"
	"net/http"
)

type NewRequestFunc func(string, string, io.Reader) (*http.Request, error)
