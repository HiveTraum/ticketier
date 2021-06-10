package domain

import "github.com/google/uuid"

type SubTicket struct {
	ID       uuid.UUID
	TicketID uuid.UUID
	GroupID  uuid.UUID
}
