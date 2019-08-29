package main

import (
	"log"
	"net/http"

	"github.com/fivestar/resozyme/_examples/resources"
	"github.com/fivestar/resozyme/resource"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	resource.Route(r, `/`, resources.NewIndexResource)
	r.Route(`/articles`, func(r chi.Router) {
		resource.Route(r, `/`, resources.NewArticlesResource)
		resource.Route(r, `/{articleID:\d+}`, resources.NewArticleResource)
	})

	logger := &resource.NilLogger{}
	debug := true

	contr := resource.NewController(r, logger, debug)
	contr.DefaultRenderer = resource.NewHALRenderer()

	log.Fatal(http.ListenAndServe(`:9393`, contr))
}
