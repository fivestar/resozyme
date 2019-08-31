package resources

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fivestar/resozyme"
	"github.com/fivestar/resozyme/_examples/repos"
	"github.com/go-chi/chi"
)

func TestArticleResource_OnGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/articles/1", nil)
	w := httptest.NewRecorder()

	ctx := r.Context()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Keys = append(rctx.URLParams.Keys, "articleID")
	rctx.URLParams.Values = append(rctx.URLParams.Values, "1")

	r = r.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))

	resc := &ArticleResource{
		Base:        resozyme.NewBase(ctx),
		view:        &articleView{},
		articleRepo: repos.NewArticleRepository(),
	}

	wantView := &articleView{
		ID:      1,
		Title:   "Foo",
		PubDate: "2019-01-01",
	}

	resc.OnGet(w, r)

	if resc.Code() != http.StatusOK {
		t.Fatalf("Unexpected code: got=%d, want=%d", resc.Code(), http.StatusOK)
	}

	gotView := resc.View().(*articleView)
	if gotView.ID != wantView.ID || gotView.Title != wantView.Title || gotView.PubDate != wantView.PubDate {
		t.Fatalf("Unexpected view: got=%#v, want=%#v", gotView, wantView)
	}
}

func TestArticleResource_OnPatch(t *testing.T) {
	r := httptest.NewRequest(http.MethodPatch, "/articles/1", bytes.NewBufferString(`{"title":"Hello, World"}`))
	w := httptest.NewRecorder()

	ctx := r.Context()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Keys = append(rctx.URLParams.Keys, "articleID")
	rctx.URLParams.Values = append(rctx.URLParams.Values, "1")

	r = r.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))

	resc := &ArticleResource{
		Base:        resozyme.NewBase(ctx),
		view:        &articleView{},
		articleRepo: repos.NewArticleRepository(),
	}

	wantView := &articleView{
		ID:      1,
		Title:   "Hello, World",
		PubDate: "2019-01-01",
	}

	resc.OnPatch(w, r)

	if resc.Code() != http.StatusOK {
		t.Fatalf("Unexpected code: got=%d, want=%d", resc.Code(), http.StatusOK)
	}

	gotView := resc.View().(*articleView)
	if gotView.ID != wantView.ID || gotView.Title != wantView.Title || gotView.PubDate != wantView.PubDate {
		t.Fatalf("Unexpected view: got=%#v, want=%#v", gotView, wantView)
	}
}
