package resozyme

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestDispatcher_ServeHTTP(t *testing.T) {
	r := chi.NewRouter()

	dispatcher := &Dispatcher{
		mux:             r,
		defaultRenderer: NewJSONRenderer(),
		errorHandler:    &ExposedErrorHandler{Renderer: NewJSONRenderer()},
		logger:          &NilLogger{},
		debug:           false,
		prettyKey:       "pretty",
	}
	dispatcher.SetDefaultRenderer(NewHALRenderer())

	Route(r, "/hello", newHelloResource)

	tests := []struct {
		path       string
		method     string
		wantCode   int
		wantBody   []byte
		wantHeader http.Header
	}{
		{
			"/hello",
			http.MethodGet,
			http.StatusOK,
			[]byte(`{"_links":{"self":{"href":"/hello"}},"text":"Hello, World"}`),
			http.Header{
				"Content-Type": []string{"application/hal+json"},
			},
		},
		{
			"/hello",
			http.MethodPost,
			http.StatusCreated,
			nil,
			http.Header{
				"Location": []string{"https://example.com/loc"},
			},
		},
		{
			"/hello",
			http.MethodPut,
			http.StatusMethodNotAllowed,
			[]byte(`{"message":"Method Not Allowed"}`),
			http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		{
			"/hello",
			http.MethodPatch,
			http.StatusMethodNotAllowed,
			[]byte(`{"message":"Method Not Allowed"}`),
			http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		{
			"/hello",
			http.MethodDelete,
			http.StatusMethodNotAllowed,
			[]byte(`{"message":"Method Not Allowed"}`),
			http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.method, func(t *testing.T) {
			t.Parallel()

			r := httptest.NewRequest(tt.method, "http://example.com"+tt.path, nil)
			w := httptest.NewRecorder()

			dispatcher.ServeHTTP(w, r)

			resp := w.Result()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Unexpected error: got=%s", err)
			}

			if resp.StatusCode != tt.wantCode {
				t.Fatalf("Unexpected code: got=%d, want=%d", resp.StatusCode, tt.wantCode)
			}

			if !bytes.Equal(body, tt.wantBody) {
				t.Fatalf("Unexpected body: got=%s, want=%s", string(body), string(tt.wantBody))
			}

			if len(resp.Header) != len(tt.wantHeader) {
				t.Fatalf("Unexpected header len: got=%d, want=%d", len(resp.Header), len(tt.wantHeader))
			}

			for key := range tt.wantHeader {
				if resp.Header.Get(key) != tt.wantHeader.Get(key) {
					t.Fatalf("Unexpected %s header: got=%s, want=%s", key, resp.Header.Get(key), tt.wantHeader.Get(key))
				}
			}
		})
	}
}
