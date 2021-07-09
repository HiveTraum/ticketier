package subject

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"src/domain"
)

type subjectRESTController struct {
	service domain.SubjectService
}

func NewSubjectRESTController(router *mux.Router, service domain.SubjectService) {
	controller := &subjectRESTController{service: service}
	router.HandleFunc("/", controller.Create).Methods(http.MethodPost)
	router.HandleFunc("/", controller.List).Methods(http.MethodGet)
}

func (controller *subjectRESTController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var createSubjectDTO domain.CreateSubjectDTO
	err = json.Unmarshal(body, &createSubjectDTO)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ticket, err := controller.service.Create(r.Context(), []*domain.CreateSubjectDTO{&createSubjectDTO})
	if err != nil {
		responseBody, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(responseBody)
	} else {
		responseBody, _ := json.Marshal(ticket)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}
}

func (controller *subjectRESTController) List(w http.ResponseWriter, r *http.Request) {
	ticket, err := controller.service.List(r.Context())
	if err != nil {
		responseBody, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(responseBody)
	} else {
		responseBody, _ := json.Marshal(ticket)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}
}
