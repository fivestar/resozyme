package resozyme

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrorHandler is an error handler interface.
type ErrorHandler interface {
	// IsError checks whether the given resource has the error.
	IsError(resc Resource) bool

	// HandleError transforms the resource that has the error to another resource
	// that represents the error.
	HandleError(resc Resource, w http.ResponseWriter, r *http.Request) Resource
}

// ErrorResource is an resource to represent the error.
type ErrorResource struct {
	*Base
	view *ErrorView
}

// ErrorView is a view for the error.
type ErrorView struct {
	Message string `json:"message"`
}

// View implements resource.Resource.
func (resc *ErrorResource) View() interface{} {
	return resc.view
}

// Href implements resource.Resource.
// This method returns the dummy URL contains the status code.
func (resc *ErrorResource) Href() string {
	return fmt.Sprintf("/errors/%d", resc.Code())
}

// ExposedErrorHandler is an ErrorHandler.
// [NOTICE] This handler exposes the raw error message to the response view.
type ExposedErrorHandler struct {
	Renderer Renderer
}

// IsError implements resource.ErrorHandler.
func (eh *ExposedErrorHandler) IsError(resc Resource) bool {
	if resc.Code() >= 400 {
		return true
	}

	return false
}

// HandleError implements resource.ErrorHandler.
func (eh *ExposedErrorHandler) HandleError(resc Resource, w http.ResponseWriter, r *http.Request) Resource {
	code := resc.Code()
	err := resc.Error()

	// Set general status text if no error set.
	if err == nil {
		err = errors.New(http.StatusText(code))
	}

	errResc := &ErrorResource{
		Base: NewBase(r.Context()),
		view: &ErrorView{},
	}

	errResc.SetCode(code)
	errResc.view.Message = err.Error()
	errResc.SetRenderer(eh.Renderer)

	for k, v := range resc.Header() {
		errResc.Header()[k] = v
	}

	return errResc
}
