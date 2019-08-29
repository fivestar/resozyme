package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fivestar/resozyme/_examples/model"
	"github.com/fivestar/resozyme/_examples/repos"
	"github.com/fivestar/resozyme/resource"
	"github.com/go-chi/chi"
)

// NewArticleResource creates an ArticleResource.
func NewArticleResource(ctx context.Context) resource.Resource {
	return &ArticleResource{
		Base:        resource.NewBase(ctx),
		view:        &articleView{},
		articleRepo: repos.NewArticleRepository(),
	}
}

// ArticleResource is a resource.
type ArticleResource struct {
	*resource.Base
	view        *articleView
	articleID   int64
	articleRepo repos.ArticleRepository
}

// articleView is a view for the ArticleResource.
type articleView struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	PubDate string `json:"pubDate"`
}

// View returns a resource view.
func (resc *ArticleResource) View() interface{} {
	return resc.view
}

// Href returns a path of resource.
func (resc *ArticleResource) Href() string {
	return fmt.Sprintf("/articles/%d", resc.articleID)
}

// Bind binds specified values to the resource.
func (resc *ArticleResource) Bind(i interface{}) {
	switch v := i.(type) {
	case *model.Article:
		resc.articleID = v.ID
		resc.view.ID = v.ID
		resc.view.Title = v.Title
		resc.view.PubDate = v.PubDate.Format("2006-01-02")
	}
}

// OnGet handles the GET request.
func (resc *ArticleResource) OnGet(w http.ResponseWriter, r *http.Request) {
	articleID, _ := strconv.ParseInt(chi.URLParam(r, "articleID"), 10, 64)

	article, err := resc.articleRepo.Find(articleID)
	if err != nil {
		resc.SetCode(http.StatusInternalServerError)
		resc.SetError(err)
		return
	}
	if article == nil {
		resc.SetCode(http.StatusNotFound)
		return
	}

	resource.BindTo(resc, article)
}

// OnPatch handles the PATCH request.
func (resc *ArticleResource) OnPatch(w http.ResponseWriter, r *http.Request) {
	articleID, _ := strconv.ParseInt(chi.URLParam(r, "articleID"), 10, 64)

	article, err := resc.articleRepo.Find(articleID)
	if err != nil {
		resc.SetCode(http.StatusInternalServerError)
		resc.SetError(err)
		return
	}
	if article == nil {
		resc.SetCode(http.StatusNotFound)
		return
	}

	//
	resource.BindTo(resc, article)

	if err := submitArticle(r, resc.view, article); err != nil {
		resc.SetCode(http.StatusBadRequest)
		resc.SetError(err)
		return
	}

	resource.BindTo(resc, article)
}

// submitArticle submits request body to an article.
func submitArticle(r *http.Request, view *articleView, article *model.Article) error {
	var err error

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(view); err != nil {
		return err
	}

	if view.Title == "" {
		return fmt.Errorf("title is required")
	}

	if view.PubDate == "" {
		return fmt.Errorf("pubDate is required")
	}

	article.Title = view.Title
	article.PubDate, err = time.Parse("2006-01-02", view.PubDate)
	if err != nil {
		return err
	}

	return nil
}
