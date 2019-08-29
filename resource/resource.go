package resource

import (
	"context"
	"net/http"
)

// Resource is an interface represents the resource.
type Resource interface {
	// Href returns the pathinfo.
	Href() string
	// Context returns the context.
	Context() context.Context

	// Code returns the status code.
	Code() int
	// Header returns the HTTP header object.
	Header() http.Header
	// View returns the view part.
	View() interface{}
	// Error returns an error if occurred.
	Error() error

	// SetCode sets the code.
	SetCode(code int)
	// SetError sets an error.
	SetError(error error)

	// SetSubstituteView sets a substitute view.
	// When the substitute view is set, the view will be overridden by that.
	SetSubstituteView(sr Resource)
	// HasSubstituteView returns whether the substitute view is set.
	HasSubstituteView() bool
	// SubstituteView returns a substitute view.
	SubstituteView() Resource

	// AddLink adds a link.
	AddLink(rel string, l Link)
	// AddLinkCollection adds the link collection.
	AddLinkCollection(rel string, l []Link)
	// LinkResource adds a link to the resource.
	LinkResource(rel string, aResc Resource)
	// LinkResourceCollection adds link to the resource collection.
	LinkResourceCollection(rel string, aResc []Resource)
	// Links returns the links.
	Links() map[string]Link
	// LinksCollection returns link collection.
	LinksCollection() map[string][]Link

	// Embed embeds a resource.
	Embed(rel string, er Resource)
	// Embed embeds the resource collection.
	EmbedCollection(rel string, er []Resource)
	// Embedded returns embedded resources.
	Embedded() map[string]Resource
	// EmbeddedCollection returns embedded the resource collection.
	EmbeddedCollection() map[string][]Resource

	// Renderer returns the renderer.
	Renderer() Renderer
	// SetRenderer sets a renderer to the resource.
	SetRenderer(renderer Renderer)

	// Bind binds passed value to the resource.
	Bind(i interface{})
	// OnGet handles the GET request.
	OnGet(w http.ResponseWriter, r *http.Request)
	// OnGet handles the POST request.
	OnPost(w http.ResponseWriter, r *http.Request)
	// OnGet handles the PUT request.
	OnPut(w http.ResponseWriter, r *http.Request)
	// OnGet handles the PATCH request.
	OnPatch(w http.ResponseWriter, r *http.Request)
	// OnGet handles the DELETE request.
	OnDelete(w http.ResponseWriter, r *http.Request)
}

// Link is a link.
type Link struct {
	Href string
}
