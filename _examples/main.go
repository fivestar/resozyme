package main

import (
	"log"
	"net/http"

	"github.com/fivestar/resozyme"
	"github.com/fivestar/resozyme/_examples/resources"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	resozyme.Route(r, `/`, resources.NewIndexResource)
	r.Route(`/articles`, func(r chi.Router) {
		resozyme.Route(r, `/`, resources.NewArticlesResource)
		resozyme.Route(r, `/{articleID:\d+}`, resources.NewArticleResource)
	})

	logger := &resozyme.NilLogger{}
	debug := true

	contr := resozyme.NewController(r, logger, debug)
	contr.DefaultRenderer = resozyme.NewHALRenderer()

	log.Fatal(http.ListenAndServe(`:9393`, contr))
}
