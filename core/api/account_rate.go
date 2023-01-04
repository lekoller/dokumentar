package api

import (
	"encoding/json"
	"net/http"

	p "github.com/lekoller/dokumentar/core/persistence"
)

type DocumentationAPI struct {
	res           http.ResponseWriter
	req           *http.Request
	slashedURL    string
	notSlashedURL string
}

func NewDocumentationAPI(w http.ResponseWriter, r *http.Request, memory *p.Memory) *DocumentationAPI {
	return &DocumentationAPI{
		res:           w,
		req:           r,
		slashedURL:    "/docs" + "/",
		notSlashedURL: "/docs",
	}
}

func (api *DocumentationAPI) Resolve() {
	api.res.Header().Set("Content=Type", "application/json")

	switch api.req.Method {
	case "POST":
		err := api.create()
		if err != nil {
			api.respondBadRequest(err.Error())
		}
	default:
		api.res.WriteHeader(405)
	}
}

func (api *DocumentationAPI) create() (err error) {
	return
}

func (api *DocumentationAPI) respondBadRequest(msg string) {
	api.res.WriteHeader(400)
	json.NewEncoder(api.res).Encode(&APIError{msg})
}
