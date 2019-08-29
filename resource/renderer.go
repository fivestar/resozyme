package resource

// Renderer is a resource renderer.
type Renderer interface {
	// Render renders resource view.
	Render(resc Resource, pretty bool) []byte
}
