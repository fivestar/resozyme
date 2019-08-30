package resozyme

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// NewBase creates a new Base.
func NewBase(ctx context.Context) *Base {
	return &Base{
		ctx:                ctx,
		code:               http.StatusOK,
		header:             http.Header{},
		links:              map[string]Link{},
		linksCollection:    map[string][]Link{},
		embedded:           map[string]Resource{},
		embeddedCollection: map[string][]Resource{},
	}
}

// Base provides the core behavior as resource.Resource.
// In principle, this struct should be embedded when defining a new resource.
//
// And then to complete to implement resource.Resource,
// the concrete resource need to be implemented `Href` method.
//
// Also, as the resource almost certainly has own view struct,
// needs to be implemented `View` method.
type Base struct {
	ctx context.Context

	code   int
	header http.Header
	error  error

	substituteView Resource

	links              map[string]Link
	linksCollection    map[string][]Link
	embedded           map[string]Resource
	embeddedCollection map[string][]Resource

	renderer Renderer
}

// Context implements resource.Resource.
func (resc *Base) Context() context.Context {
	return resc.ctx
}

// Code implements resource.Resource.
func (resc *Base) Code() int {
	return resc.code
}

// Header implements resource.Resource.
func (resc *Base) Header() http.Header {
	return resc.header
}

// View implements resource.Resource.
func (resc *Base) View() interface{} {
	return nil
}

// Error implements resource.Resource.
func (resc *Base) Error() error {
	return resc.error
}

// SetCode implements resource.Resource.
func (resc *Base) SetCode(code int) {
	resc.code = code
}

// SetError implements resource.Resource.
func (resc *Base) SetError(err error) {
	// Save the stack trace at this point
	resc.error = errors.WithStack(err)

	if resc.code == 0 {
		resc.SetCode(http.StatusInternalServerError)
	}
}

// SetSubstituteView implements resource.Resource.
func (resc *Base) SetSubstituteView(sr Resource) {
	resc.substituteView = sr
}

// HasSubstituteView implements resource.Resource.
func (resc *Base) HasSubstituteView() bool {
	return resc.substituteView != nil
}

// SubstituteView implements resource.Resource.
func (resc *Base) SubstituteView() Resource {
	return resc.substituteView
}

// Links implements resource.Resource.
func (resc *Base) Links() map[string]Link {
	return resc.links
}

// AddLink implements resource.Resource.
func (resc *Base) AddLink(rel string, l Link) {
	resc.links[rel] = l
}

// LinkResource implements resource.Resource.
func (resc *Base) LinkResource(rel string, aResc Resource) {
	resc.AddLink(rel, Link{
		Href: aResc.Href(),
	})
}

// LinkResourceCollection implements resource.Resource.
func (resc *Base) LinkResourceCollection(rel string, aResc []Resource) {
	var l []Link
	for _, ar := range aResc {
		l = append(l, Link{
			Href: ar.Href(),
		})
	}
	resc.AddLinkCollection(rel, l)
}

// LinksCollection implements resource.Resource.
func (resc *Base) LinksCollection() map[string][]Link {
	return resc.linksCollection
}

// AddLinkCollection implements resource.Resource.
func (resc *Base) AddLinkCollection(rel string, l []Link) {
	resc.linksCollection[rel] = l
}

// Embedded implements resource.Resource.
func (resc *Base) Embedded() map[string]Resource {
	return resc.embedded
}

// Embed implements resource.Resource.
func (resc *Base) Embed(rel string, er Resource) {
	resc.embedded[rel] = er
}

// EmbeddedCollection implements resource.Resource.
func (resc *Base) EmbeddedCollection() map[string][]Resource {
	return resc.embeddedCollection
}

// EmbedCollection implements resource.Resource.
func (resc *Base) EmbedCollection(rel string, er []Resource) {
	resc.embeddedCollection[rel] = er
}

// Renderer implements resource.Resource.
func (resc *Base) Renderer() Renderer {
	return resc.renderer
}

// SetRenderer implements resource.Resource.
func (resc *Base) SetRenderer(renderer Renderer) {
	resc.renderer = renderer
}

// Bind implements resource.Resource.
func (resc *Base) Bind(i interface{}) {
}

// OnGet implements resource.Resource.
// If the resource needs to handle the GET request, please override this.
// Unless defined, it responds the status 405 Method Not Allowed.
func (resc *Base) OnGet(w http.ResponseWriter, r *http.Request) {
	resc.code = http.StatusMethodNotAllowed
}

// OnPost implements resource.Resource.
// If the resource needs to handle the POST request, please override this.
// Unless defined, it responds the status 405 Method Not Allowed.
func (resc *Base) OnPost(w http.ResponseWriter, r *http.Request) {
	resc.code = http.StatusMethodNotAllowed
}

// OnPut implements resource.Resource.
// If the resource needs to handle the PUT request, please override this.
// Unless defined, it responds the status 405 Method Not Allowed.
func (resc *Base) OnPut(w http.ResponseWriter, r *http.Request) {
	resc.code = http.StatusMethodNotAllowed
}

// OnPatch implements resource.Resource.
// If the resource needs to handle the PATCH request, please override this.
// Unless defined, it responds the status 405 Method Not Allowed.
func (resc *Base) OnPatch(w http.ResponseWriter, r *http.Request) {
	resc.code = http.StatusMethodNotAllowed
}

// OnDelete implements resource.Resource.
// If the resource needs to handle the DELETE request, please override this.
// Unless defined, it responds the status 405 Method Not Allowed.
func (resc *Base) OnDelete(w http.ResponseWriter, r *http.Request) {
	resc.code = http.StatusMethodNotAllowed
}
