package resource

import (
	"context"
	"fmt"
	"net/http"
)

type helloUser struct {
	Name string
}

func newHelloResource(ctx context.Context) Resource {
	return &helloResource{
		Base: NewBase(ctx),
		view: &helloView{},
	}
}

// helloResource is a resource.
type helloResource struct {
	*Base
	view *helloView
}

type helloView struct {
	Text string `json:"text"`
}

func (resc *helloResource) View() interface{} {
	return resc.view
}

func (resc *helloResource) Href() string {
	return "/hello"
}

func (resc *helloResource) Bind(i interface{}) {
	switch v := i.(type) {
	case *helloUser:
		resc.view.Text = fmt.Sprintf("Hello, %s", v.Name)
	}
}

func (resc *helloResource) OnGet(w http.ResponseWriter, r *http.Request) {
	user := &helloUser{Name: "World"}
	BindTo(resc, user)
}
