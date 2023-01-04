package api

import "net/http"

type MuxGuardian struct {
	mux *http.ServeMux
}

func NewMuxGuardian(mux *http.ServeMux) *MuxGuardian {
	return &MuxGuardian{mux}
}

func (mg *MuxGuardian) registerEndpoint(url string, callback func(w http.ResponseWriter, r *http.Request)) {
	mg.mux.HandleFunc(url, callback)
	mg.mux.HandleFunc(url+"/", callback)
}
