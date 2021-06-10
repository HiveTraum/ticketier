package ticket

import (
	"encoding/json"
	"io"
	"net/http"
	"src/domain"
)

type ticketRestController struct {
	ticketService domain.TicketService
}

func NewTicketRESTController(service domain.TicketService) *ticketRestController {
	return &ticketRestController{ticketService: service}
}

func (controller *ticketRestController) Create(w http.ResponseWriter, r *http.Request) {
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

	ticket, err := controller.ticketService.Create(&createTicketDTO)
	if err != nil {
		responseBody, _ := json.Marshal(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(responseBody)
	} else {
		responseBody, _ := json.Marshal(ticket)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}
}
