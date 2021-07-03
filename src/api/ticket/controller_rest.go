package ticket

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"src/domain"
)

type ticketRESTController struct {
	ticketService domain.TicketService
}

func NewTicketRESTController(router *mux.Router, service domain.TicketService) {
	controller := &ticketRESTController{ticketService: service}
	router.HandleFunc("/", controller.Create).Methods(http.MethodPost)
	router.HandleFunc("/", controller.Get).Methods(http.MethodGet)
}

func (controller *ticketRESTController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var createTicketDTO domain.CreateTicketDTO
	err = json.Unmarshal(body, &createTicketDTO)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ticket, err := controller.ticketService.Create(r.Context(), &createTicketDTO)
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

func (controller *ticketRESTController) Get(w http.ResponseWriter, r *http.Request) {

}
