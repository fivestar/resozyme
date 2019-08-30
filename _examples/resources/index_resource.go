package resources

import (
	"context"
	"net/http"

	"github.com/fivestar/resozyme"
)

// NewIndexResource creates an IndexResource.
func NewIndexResource(ctx context.Context) resozyme.Resource {
	return &IndexResource{
		Base: resozyme.NewBase(ctx),
		view: &struct{}{},
	}
}

// IndexResource is a resource.
type IndexResource struct {
	*resozyme.Base
	view      *struct{}
	articleID int64
}

// View returns a resource view.
func (resc *IndexResource) View() interface{} {
	return resc.view
}

// Href returns a path of resource.
func (resc *IndexResource) Href() string {
	return "/"
}

// OnGet handles the GET request.
func (resc *IndexResource) OnGet(w http.ResponseWriter, r *http.Request) {
	resc.LinkResource("articles", resozyme.BindTo(NewArticlesResource(resc.Context())))
}
