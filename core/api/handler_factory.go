package api

import (
	"net/http"

	p "github.com/lekoller/dokumentar/core/persistence"
)

type HandlerFactory struct {
	memory *p.Memory
}

func NewHandlerFactory(memory *p.Memory) *HandlerFactory {
	return &HandlerFactory{memory}
}

func (hf *HandlerFactory) mountDocumentationHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateAccess(w, r) {
			return
		}
		api := NewDocumentationAPI(w, r, hf.memory)
		api.Resolve()
	}
}
