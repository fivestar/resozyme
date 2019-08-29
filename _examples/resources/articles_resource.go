package resources

import (
	"context"
	"net/http"

	"github.com/fivestar/resozyme/_examples/model"
	"github.com/fivestar/resozyme/_examples/repos"
	"github.com/fivestar/resozyme/resource"
)

// NewArticlesResource creates an ArticlesResource.
func NewArticlesResource(ctx context.Context) resource.Resource {
	return &ArticlesResource{
		Base:        resource.NewBase(ctx),
		view:        &articlesView{},
		articleRepo: repos.NewArticleRepository(),
	}
}

// ArticlesResource is a resource.
type ArticlesResource struct {
	*resource.Base
	view        *articlesView
	articleRepo repos.ArticleRepository
}

// articlesView is a view for the ArticlesResource.
type articlesView struct{}

// View returns a resource view.
func (resc *ArticlesResource) View() interface{} {
	return resc.view
}

// Href returns a path of resource.
func (resc *ArticlesResource) Href() string {
	return "/articles"
}

// Bind binds specified values to the resource.
func (resc *ArticlesResource) Bind(i interface{}) {
	switch v := i.(type) {
	case []*model.Article:
		var rescs []resource.Resource
		for _, article := range v {
			rescs = append(rescs, resource.BindTo(NewArticleResource(resc.Context()), article))
		}
		resc.EmbedCollection("articles", rescs)
	}
}

// OnGet handles the GET request.
func (resc *ArticlesResource) OnGet(w http.ResponseWriter, r *http.Request) {
	articles, err := resc.articleRepo.FindLatest(5)
	if err != nil {
		resc.SetCode(http.StatusInternalServerError)
		resc.SetError(err)
		return
	}

	resource.BindTo(resc, articles)
}

// OnPost handles the POST request.
func (resc *ArticlesResource) OnPost(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{}
	if err := submitArticle(r, &articleView{}, article); err != nil {
		resc.SetCode(http.StatusBadRequest)
		resc.SetError(err)
		return
	}

	if err := resc.articleRepo.Add(article); err != nil {
		resc.SetCode(http.StatusInternalServerError)
		resc.SetError(err)
		return
	}

	aResc := resource.BindTo(NewArticleResource(resc.Context()), article)
	resc.SetSubstituteView(aResc)
	resc.SetCode(http.StatusCreated)
	resc.Header().Add("Location", aResc.Href())
}
