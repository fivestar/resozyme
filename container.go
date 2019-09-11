package resozyme

import (
	"context"
	"errors"
)

// ActiveResourceContainerContextKey is a context key.
type ActiveResourceContainerContextKey struct{}

// NewResourceContainer creates a new ResourceContainer.
func NewResourceContainer() *ResourceContainer {
	return &ResourceContainer{}
}

// WithController sets a controller to the context.
func WithController(ctx context.Context, rcont Controller) context.Context {
	return context.WithValue(ctx, ControllerContextKey{}, rcont)
}

// GetController gets a controller from the context.
func GetController(ctx context.Context) Controller {
	v := ctx.Value(ControllerContextKey{})
	if v == nil {
		return nil
	}
	return v.(Controller)
}

// WithActiveResourceContainer sets a resource container.
func WithActiveResourceContainer(ctx context.Context, rcont *ResourceContainer) context.Context {
	return context.WithValue(ctx, ActiveResourceContainerContextKey{}, rcont)
}

// GetActiveResourceContainer gets a resource container.
func GetActiveResourceContainer(ctx context.Context) *ResourceContainer {
	v := ctx.Value(ActiveResourceContainerContextKey{})
	if v == nil {
		return nil
	}
	return v.(*ResourceContainer)
}

// ActivateResource sets a resource to ActiveResourceContainer.
func ActivateResource(ctx context.Context, resc Resource) error {
	rcont := GetActiveResourceContainer(ctx)
	if rcont == nil {
		return errors.New("no resource container activated")
	}

	rcont.Set(resc)

	return nil
}

// GetActiveResource gets the activated resource from ActiveResourceContainer.
func GetActiveResource(ctx context.Context) Resource {
	rcont := GetActiveResourceContainer(ctx)
	if rcont == nil {
		return nil
	}

	if !rcont.Exists() {
		return nil
	}

	return rcont.Get()
}

// ResourceContainer is a container.
type ResourceContainer struct {
	resc Resource
}

// Exists checks whether the container has a resource.
func (rcont *ResourceContainer) Exists() bool {
	return rcont.resc != nil
}

// Set sets a resource to the container.
func (rcont *ResourceContainer) Set(resc Resource) {
	rcont.resc = resc
}

// Get gets a resource from the container.
func (rcont *ResourceContainer) Get() Resource {
	return rcont.resc
}
