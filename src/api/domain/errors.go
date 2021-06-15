package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type TicketAnswerValueTypeError struct {
	AnswerTitle string
	Value       interface{}
}

func (err *TicketAnswerValueTypeError) Error() string {
	return fmt.Sprintf("answer type %s is not valid for %s", err.Value, err.AnswerTitle)
}

type entityNotFound struct {
	Entity string
	ID     uuid.UUID
}

func (err *entityNotFound) Error() string {
	return fmt.Sprintf("%s with id %s not found", err.Entity, err.ID.String())
}

func EntityNotFound(entity string, id uuid.UUID) *entityNotFound {
	return &entityNotFound{Entity: entity, ID: id}
}

func SubjectFieldNotFound(id uuid.UUID) *entityNotFound {
	return EntityNotFound("subject field", id)
}

var (
	FileTypeNotFound = errors.New("file type not found")
	AnswerRequired   = errors.New("answer is required")
)

var (
	ticketAnswerValueTypeError = errors.New("ticket answer value type error")
)
