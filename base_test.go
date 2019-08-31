package resozyme

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBase(t *testing.T) {
	ctx := context.Background()
	base := NewBase(ctx)

	if base.Context() != ctx {
		t.Fatalf("Unexpected context")
	}

	if base.Code() != http.StatusOK {
		t.Fatalf("Unexpected code: got=%d, want=%d", base.Code(), http.StatusOK)
	}

	if base.Error() != nil {
		t.Fatalf("Unexpected error: got=%s", base.Error())
	}

	if base.Renderer() != nil {
		t.Fatalf("Unexpected renderer: got=%+v", base.Renderer())
	}

	base.Bind(nil)
}

func TestBase_SetError(t *testing.T) {
	ctx := context.Background()
	base := NewBase(ctx)

	err := errors.New("Error")
	base.SetCode(http.StatusInternalServerError)
	base.SetError(err)

	if base.Code() != http.StatusInternalServerError {
		t.Fatalf("Unexpected code: got=%d, want=%d", base.Code(), http.StatusInternalServerError)
	}
}

func TestBase_SetRenderer(t *testing.T) {
	ctx := context.Background()
	base := NewBase(ctx)

	renderer := NewJSONRenderer()
	base.SetRenderer(renderer)

	if base.Renderer() != renderer {
		t.Fatalf("Unexpected renderer: got=%+v, want=%+v", base.Renderer(), renderer)
	}
}

func TestBase_Handler(t *testing.T) {
	tests := []struct {
		method  string
		handler func(base *Base, w http.ResponseWriter, r *http.Request)
	}{
		{
			http.MethodGet,
			func(base *Base, w http.ResponseWriter, r *http.Request) {
				base.OnGet(w, r)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.method, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			base := NewBase(ctx)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			tt.handler(base, w, r)

			if base.Code() != http.StatusMethodNotAllowed {
				t.Fatalf("Unexpected code: got=%d, want=%d", base.Code(), http.StatusMethodNotAllowed)
			}
		})
	}
}
