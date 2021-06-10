package domain

import (
	"errors"
	"fmt"
)

type TicketAnswerValueTypeError struct {
	AnswerTitle string
	Value       interface{}
}

func (err *TicketAnswerValueTypeError) Error() string {
	return fmt.Sprintf("answer type %s is not valid for %s", err.Value, err.AnswerTitle)
}

var (
	FileTypeNotFound = errors.New("file type not found")
	AnswerRequired   = errors.New("answer is required")
)
