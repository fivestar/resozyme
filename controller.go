package resozyme

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
)

// ControllerContextKey is a context key.
type ControllerContextKey struct{}

// ResourceFactory is a resource factory.
type ResourceFactory = func(context.Context) Resource

// NewController creates a controller.
func NewController(r chi.Router, logger Logger, debug bool) *Dispatcher {
	return &Dispatcher{
		router:          r,
		defaultRenderer: NewJSONRenderer(),
		errorHandler:    &ExposedErrorHandler{Renderer: NewJSONRenderer()},
		logger:          logger,
		debug:           debug,
		prettyKey:       "pretty",
	}
}

// Route binds a ResourceFactory to the router.
func Route(mx chi.Router, path string, fac ResourceFactory) {
	mx.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		contr := GetController(ctx)
		logger := contr.Logger()

		resc := fac(ctx)

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

		// Render error.
		if contr.IsError(resc) {
			errResc := contr.HandleError(resc, w, r)
			contr.Dispatch(errResc, w, r)
			return
		}

		contr.Dispatch(resc, w, r)
	})
}

// Controller provides a resource controller interface.
type Controller interface {
	http.Handler

	// Dispatch dispatches a resource.
	Dispatch(resc Resource, w http.ResponseWriter, r *http.Request)

	// SetDefaultRenderer sets a default renderer.
	SetDefaultRenderer(renderer Renderer)

	// IsError checks whether a resource has an error or not.
	IsError(resc Resource) bool

	// HandleError handles an error resource.
	HandleError(resc Resource, w http.ResponseWriter, r *http.Request) Resource

	// Logger returns a logger.
	Logger() Logger
}

// Dispatcher handles HTTP request and response.
type Dispatcher struct {
	router          chi.Router
	defaultRenderer Renderer
	errorHandler    ErrorHandler
	logger          Logger
	debug           bool
	prettyKey       string
}

// ServeHTTP implements http.Handler.
func (dispatcher *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = WithController(ctx, dispatcher)
	ctx = WithActiveResourceContainer(ctx, NewResourceContainer())

	r = r.WithContext(ctx)

	dispatcher.router.ServeHTTP(w, r)
}

// SetDefaultRenderer implements Controller.
func (dispatcher *Dispatcher) SetDefaultRenderer(renderer Renderer) {
	dispatcher.defaultRenderer = renderer
}

// Dispatch implements Controller.
func (dispatcher *Dispatcher) Dispatch(resc Resource, w http.ResponseWriter, r *http.Request) {
	logger := dispatcher.Logger()

	renderer := resc.Renderer()
	if renderer == nil {
		renderer = dispatcher.defaultRenderer
	}

	logger.Debugf("Renderer: %T", renderer)

	// check nil
	var bytes []byte
	view := resc.View()
	vv := reflect.ValueOf(view)
	if view != nil && (vv.Kind() == reflect.Struct || !vv.IsNil()) {
		bytes = renderer.Render(resc, dispatcher.prettyEnabled(r))
	}

	for k, v := range resc.Header() {
		w.Header()[k] = v
	}
	w.WriteHeader(resc.Code())
	w.Write(bytes)
}

// IsError implements Controller.
func (dispatcher *Dispatcher) IsError(resc Resource) bool {
	return dispatcher.errorHandler.IsError(resc)
}

// HandleError implements Controller.
func (dispatcher *Dispatcher) HandleError(resc Resource, w http.ResponseWriter, r *http.Request) Resource {
	return dispatcher.errorHandler.HandleError(resc, w, r)
}

// Logger implements Controller.
func (dispatcher *Dispatcher) Logger() Logger {
	return dispatcher.logger
}

// prettyEnabled checks whether the pretty-rendering mode is enabled or not.
func (dispatcher *Dispatcher) prettyEnabled(r *http.Request) bool {
	enabled := dispatcher.debug

	if v := r.URL.Query().Get(dispatcher.prettyKey); v != "" {
		enabled = (v != "0")
	}

	return enabled
}
