package resource

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestController_ServeHTTP(t *testing.T) {
	mx := chi.NewRouter()

	contr := NewController(mx, &NilLogger{}, false)
	contr.DefaultRenderer = &HALRenderer{}

	Route(mx, "/hello", newHelloResource)

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
			http.Header{},
		},
		{
			"/hello",
			http.MethodPatch,
			http.StatusMethodNotAllowed,
			[]byte(`{"message":"Method Not Allowed"}`),
			http.Header{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.method, func(t *testing.T) {
			t.Parallel()

			r := httptest.NewRequest(tt.method, "http://example.com"+tt.path, nil)
			w := httptest.NewRecorder()

			contr.ServeHTTP(w, r)

			resp := w.Result()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Unexpected error: got=%s", err)
			}

			if resp.StatusCode != tt.wantCode {
				t.Fatalf("Unexpected code: got=%d, want=%d", resp.StatusCode, tt.wantCode)
			}

			if bytes.Compare(body, tt.wantBody) != 0 {
				t.Fatalf("Unexpected body: got=%s, want=%s", string(body), string(tt.wantBody))
			}

			for key := range tt.wantHeader {
				if resp.Header.Get(key) != tt.wantHeader.Get(key) {
					t.Fatalf("Unexpected %s header: got=%s, want=%s", key, resp.Header.Get(key), tt.wantHeader.Get(key))
				}
			}
		})
	}
}
