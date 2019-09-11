package resozyme

import (
	"context"
	"testing"
)

func TestResourceContainer(t *testing.T) {
	rcont := &ResourceContainer{}

	resc := newHelloResource(context.Background())

	rcont.Set(resc)

	if !rcont.Exists() {
		t.Fatal("unexpected existence: want=true, got=false")
	}

	got := rcont.Get()
	if got != resc {
		t.Fatalf("unexpected gotten: want=%v, got=%v", resc, got)
	}
}

func TestResourceContainer_ByDefault(t *testing.T) {
	rcont := &ResourceContainer{}

	if rcont.Exists() {
		t.Fatal("unexpected existence: want=false, got=true")
	}

	got := rcont.Get()
	if got != nil {
		t.Fatalf("unexpected gotten: want=nil, got=%v", got)
	}
}

func TestResourceContainer_SetNil(t *testing.T) {
	rcont := &ResourceContainer{}

	rcont.Set(nil)

	if rcont.Exists() {
		t.Fatal("unexpected existence: want=false, got=true")
	}

	got := rcont.Get()
	if got != nil {
		t.Fatalf("unexpected gotten: want=nil, got=%v", got)
	}
}
