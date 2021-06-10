package main

import "src/ticket"

func main() {
	ticketRepository := ticket.
	ticketService := ticket.NewTicketService()
	ticketRestController := ticket.NewTicketRESTController(ticketService)
}