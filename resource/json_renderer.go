package resource

import (
	"encoding/json"
)

// NewJSONRenderer creates a new JSONRenderer.
func NewJSONRenderer() *JSONRenderer {
	return &JSONRenderer{}
}

// JSONRenderer represents the resource in json format.
type JSONRenderer struct {
}

// Render implements resource.Renderer.
func (renderer *JSONRenderer) Render(resc Resource, pretty bool) []byte {
	if resc.Header().Get("Content-Type") == "" {
		resc.Header().Set("Content-Type", "application/json")
	}

	view := resc.View()

	var payload []byte
	if pretty {
		payload, _ = json.MarshalIndent(view, "", "  ")
	} else {
		payload, _ = json.Marshal(view)
	}

	return payload
}
