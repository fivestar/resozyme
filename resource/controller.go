package resource

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
)

// ControllerContextKey is a context key.
type ControllerContextKey struct{}

// ResourceFactory is a resource.Resource factory.
type ResourceFactory = func(context.Context) Resource

// NewController creates a new resource controller.
func NewController(r chi.Router, logger Logger, debug bool) *Controller {
	return &Controller{
		Router:          r,
		DefaultRenderer: NewJSONRenderer(),
		ErrorHandler: &ExposedErrorHandler{
			Renderer: NewJSONRenderer(),
		},
		Logger:             logger,
		Debug:              debug,
		PrettyRenderingKey: "pretty",
	}
}

// Controller handles HTTP request and response.
type Controller struct {
	Router             chi.Router
	DefaultRenderer    Renderer
	ErrorHandler       ErrorHandler
	Logger             Logger
	Debug              bool
	PrettyRenderingKey string
}

// ServeHTTP implements http.Handler.
func (contr *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = WithController(ctx, contr)
	ctx = WithActiveResourceContainer(ctx, NewResourceContainer())

	r = r.WithContext(ctx)

	contr.Router.ServeHTTP(w, r)
}

// dispatch dispatches the resource to the response.
func (contr *Controller) dispatch(resc Resource, w http.ResponseWriter, r *http.Request) {
	renderer := resc.Renderer()
	if renderer == nil {
		renderer = contr.DefaultRenderer
	}

	contr.Logger.Debugf(`Renderer: %T`, renderer)

	rr := resc
	if resc.HasSubstituteView() {
		rr = resc.SubstituteView()
	}

	// check nil
	var bytes []byte
	view := rr.View()
	vv := reflect.ValueOf(view)
	if view != nil && (vv.Kind() == reflect.Struct || !vv.IsNil()) {
		bytes = renderer.Render(rr, contr.isPrettyRendering(r))
	}

	for k, v := range resc.Header() {
		w.Header()[k] = v
	}
	w.WriteHeader(resc.Code())
	w.Write(bytes)
}

// isPrettyRendering checks whether the pretty-rendering mode is enabled or not.
func (contr *Controller) isPrettyRendering(r *http.Request) bool {
	enabled := contr.Debug

	if v := r.URL.Query().Get(contr.PrettyRenderingKey); v != "" {
		enabled = (v != "0")
	}

	return enabled
}

// Route binds a ResourceFactory to the router.
func Route(mx chi.Router, path string, factory ResourceFactory) {
	mx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		contr := GetController(ctx)
		logger := contr.Logger

		resc := factory(ctx)

		if err := ActivateResource(ctx, resc); err != nil {
			logger.Errorf("Failed to activate resource: %v", err)
		}

		logger.Debugf(`Matched to %T`, resc)

		switch r.Method {
		case http.MethodGet:
			resc.OnGet(w, r)
		case http.MethodPost:
			resc.OnPost(w, r)
		case http.MethodPut:
			resc.OnPut(w, r)
		case http.MethodPatch:
			resc.OnPatch(w, r)
		case http.MethodDelete:
			resc.OnDelete(w, r)
		default:
			resc.SetCode(http.StatusMethodNotAllowed)
		}

		// Prioritize to render a substitute view over an error view.
		if contr.ErrorHandler.IsError(resc) && !resc.HasSubstituteView() {
			errResc := contr.ErrorHandler.HandleError(resc, w, r)
			contr.dispatch(errResc, w, r)
			return
		}

		contr.dispatch(resc, w, r)
	})
}
