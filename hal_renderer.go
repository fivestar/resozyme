package resozyme

import (
	"encoding/json"
	"reflect"

	"github.com/nvellon/hal"
)

// NewHALRenderer creates a new HALRenderer.
func NewHALRenderer() *HALRenderer {
	return &HALRenderer{}
}

// HALRenderer represents the resource in hal+json format.
type HALRenderer struct {
}

// Render implements resource.Renderer.
func (renderer *HALRenderer) Render(resc Resource, pretty bool) []byte {
	if resc.Header().Get("Content-Type") == "" {
		resc.Header().Set("Content-Type", "application/hal+json")
	}

	h := renderer.buildHALResource(resc)

	var payload []byte
	if pretty {
		payload, _ = json.MarshalIndent(h, "", "  ")
	} else {
		payload, _ = json.Marshal(h)
	}

	return payload
}

func (renderer *HALRenderer) buildHALResource(resc Resource) *hal.Resource {
	view := resc.View()

	// WTF: nvellon/hal cannot handle the pointer-type value.
	// This code converts pointer-type to concrete-type.
	view = reflect.Indirect(reflect.ValueOf(view)).Interface()

	h := hal.NewResource(view, resc.Href())

	for rel, link := range resc.Links() {
		h.AddLink(hal.Relation(rel), hal.NewLink(link.Href))
	}

	for rel, links := range resc.LinksCollection() {
		l := hal.LinkCollection{}
		for _, link := range links {
			l = append(l, hal.NewLink(link.Href))
		}
		h.AddLinkCollection(hal.Relation(rel), l)
	}

	for rel, er := range resc.Embedded() {
		h.Embed(hal.Relation(rel), renderer.buildHALResource(er))
	}

	for rel, ers := range resc.EmbeddedCollection() {
		h.EmbedCollection(hal.Relation(rel), renderer.buildHALResourceCollection(ers))
	}

	return h
}

func (renderer *HALRenderer) buildHALResourceCollection(rc []Resource) hal.ResourceCollection {
	hc := hal.ResourceCollection{}

	for _, resc := range rc {
		hc = append(hc, renderer.buildHALResource(resc))
	}

	return hc
}
