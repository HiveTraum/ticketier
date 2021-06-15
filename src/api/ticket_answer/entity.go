package ticket_answer

import (
	"github.com/google/uuid"
	"reflect"
	"src/domain"
)

func New(dto *domain.CreateTicketAnswerDTO) (*domain.TicketAnswer, error) {
	if dto.SubjectField == nil {
		return nil, domain.SubjectFieldNotFound(dto.SubjectFieldID)
	}

	subjectField := dto.SubjectField

	_answer := &domain.TicketAnswer{
		ID:             uuid.New(),
		TicketID:       dto.TicketID,
		Title:          subjectField.Title,
		ProgrammaticID: subjectField.ProgrammaticID,
	}

	parsedTicketAnswerValue, err := parseAnswerValue(subjectField.Type, dto.Value, subjectField.Required)
	if err != nil {
		if err == ticketAnswerValueTypeError {
			return nil, &domain.TicketAnswerValueTypeError{AnswerTitle: subjectField.Title, Value: dto.Value}
		}

		return nil, err
	}

	answerValue, err := createAnswerValue(parsedTicketAnswerValue)
	if err != nil {
		return nil, err
	}

	_answer.Value = answerValue
	return _answer, nil
}

func createAnswerValue(parsedAnswerValue *ticketAnswerValue) (*domain.TicketAnswerValue, error) {
	answerValue := &domain.TicketAnswerValue{Type: parsedAnswerValue._type}
	if parsedAnswerValue.isNil {
		return answerValue, nil
	}

	switch parsedAnswerValue._type {
	case domain.Flag:
		answerValue.Flag = &parsedAnswerValue._bool
	case domain.Number:
		answerValue.Number = &parsedAnswerValue.number
	case domain.String:
		answerValue.String = &parsedAnswerValue._string
	}

	return answerValue, nil
}

func parseAnswerValue(fieldType domain.FieldType, value interface{}, isRequired bool) (*ticketAnswerValue, error) {
	isNil := value == nil || reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()
	if isNil {
		if isRequired {
			return nil, domain.AnswerRequired
		} else {
			return &ticketAnswerValue{_type: fieldType, isNil: isNil}, nil
		}
	}

	switch fieldType {
	case domain.Number:
		return parseAnswerValueNumber(value)
	case domain.String:
		return parseAnswerValueString(value)
	case domain.Flag:
		return parseAnswerValueBoolean(value)
	default:
		return nil, domain.FileTypeNotFound
	}
}

func parseAnswerValueNumber(value interface{}) (*ticketAnswerValue, error) {
	var _value float64
	switch v := value.(type) {
	case float64:
		_value = v
	case float32:
		_value = float64(v)
	case int64:
		_value = float64(v)
	case int32:
		_value = float64(v)
	case int16:
		_value = float64(v)
	case int8:
		_value = float64(v)
	case int:
		_value = float64(v)
	default:
		return nil, ticketAnswerValueTypeError
	}

	return &ticketAnswerValue{_type: domain.Number, number: _value}, nil
}

func parseAnswerValueString(value interface{}) (*ticketAnswerValue, error) {
	_value, ok := value.(string)
	if !ok {
		return nil, ticketAnswerValueTypeError
	}

	return &ticketAnswerValue{_type: domain.String, _string: _value}, nil
}

func parseAnswerValueBoolean(value interface{}) (*ticketAnswerValue, error) {
	_value, ok := value.(bool)
	if !ok {
		return nil, ticketAnswerValueTypeError
	}

	return &ticketAnswerValue{_type: domain.Flag, _bool: _value}, nil
}
