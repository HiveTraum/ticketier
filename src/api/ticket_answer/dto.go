package ticket_answer

import "src/domain"

type ticketAnswerValue struct {
	_type   domain.FieldType
	_bytes  []byte
	_string string
	number  float64
	_bool   bool
	isNil   bool
}
