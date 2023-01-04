package api

import (
	"log"
	"net/http"

	"github.com/lekoller/dokumentar/core/persistence"
)

func StartRestAPI() {
	memory := &persistence.Memory{}
	factory := NewHandlerFactory(memory)

	accountRateAPIHandler := factory.mountDocumentationHandler()

	mg := NewMuxGuardian(http.NewServeMux())

	mg.registerEndpoint("/docs", accountRateAPIHandler)

	log.Println("API listening to :8080")

	err := http.ListenAndServe(":8080", mg.mux)
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}
}
