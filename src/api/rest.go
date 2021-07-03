package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"src/domain"
	"src/subject"
	"src/ticket"
)

func InitREST(ticketService domain.TicketService, subjectService domain.SubjectService) error {
	router := mux.NewRouter()

	apiV1Path := router.PathPrefix("/api/v1").Subrouter()
	ticket.NewTicketRESTController(apiV1Path.PathPrefix("/tickets").Subrouter(), ticketService)
	subject.NewSubjectRESTController(apiV1Path.PathPrefix("/subjects").Subrouter(), subjectService)

	http.Handle("/", router)
	return http.ListenAndServe("0.0.0.0:8000", nil)
}
