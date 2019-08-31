package resources

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fivestar/resozyme"
	"github.com/fivestar/resozyme/_examples/repos"
)

func TestArticlesResource_OnGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/articles", nil)
	w := httptest.NewRecorder()

	ctx := r.Context()

	resc := &ArticlesResource{
		Base:        resozyme.NewBase(ctx),
		view:        &articlesView{},
		articleRepo: repos.NewArticleRepository(),
	}

	wantView := &articlesView{}

	resc.OnGet(w, r)

	if resc.Code() != http.StatusOK {
		t.Fatalf("Unexpected code: got=%d, want=%d", resc.Code(), http.StatusOK)
	}

	gotView := resc.View().(*articlesView)
	if gotView != wantView {
		t.Fatalf("Unexpected view: got=%#v, want=%#v", gotView, wantView)
	}
}

func TestArticlesResource_OnPost(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBufferString(`{"title":"Hello","pubDate":"2019-09-01"}`))
	w := httptest.NewRecorder()

	ctx := r.Context()

	resc := &ArticlesResource{
		Base:        resozyme.NewBase(ctx),
		view:        &articlesView{},
		articleRepo: repos.NewArticleRepository(),
	}

	wantView := &articleView{
		ID:      3,
		Title:   "Hello",
		PubDate: "2019-09-01",
	}

	resc.OnPost(w, r)

	if resc.Code() != http.StatusCreated {
		t.Fatalf("Unexpected code: got=%d, want=%d", resc.Code(), http.StatusCreated)
	}

	gotView := resc.View().(*articleView)
	if gotView.ID != wantView.ID || gotView.Title != wantView.Title || gotView.PubDate != wantView.PubDate {
		t.Fatalf("Unexpected view: got=%#v, want=%#v", gotView, wantView)
	}
}
