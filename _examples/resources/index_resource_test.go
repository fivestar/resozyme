package resources

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fivestar/resozyme"
)

func TestIndexResource_OnGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	ctx := r.Context()

	resc := &IndexResource{
		Base: resozyme.NewBase(ctx),
		view: &indexView{},
	}

	wantView := &indexView{}

	resc.OnGet(w, r)

	if resc.Code() != http.StatusOK {
		t.Fatalf("Unexpected code: got=%d, want=%d", resc.Code(), http.StatusOK)
	}

	gotView := resc.View().(*indexView)
	if gotView != wantView {
		t.Fatalf("Unexpected view: got=%#v, want=%#v", gotView, wantView)
	}
}
