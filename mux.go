package resozyme

import (
	"net/http"
)

// Mux is an HTTP request multiplexer. This interface is designed to
// be compatible with chi.Mux.
type Mux interface {
	http.Handler

	// HandleFunc registers the handler function for the given pattern.
	HandleFunc(pattern string, handler http.HandlerFunc)
}
