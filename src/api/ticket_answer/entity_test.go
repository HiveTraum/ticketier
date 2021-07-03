package ticket_answer

import (
	"github.com/bxcodec/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"src/domain"
	"testing"
)

func init() {
	_ = faker.AddProvider("uuid", func(v reflect.Value) (interface{}, error) {
		return uuid.NewRandom()
	})
}

func TestParseAnswerBooleanTrue(t *testing.T) {
	t.Parallel()

	var value interface{}
	value = true
	parsedValue, err := parseAnswerValueBoolean(value)
	assert.Nil(t, err)
	assert.Equal(t, true, parsedValue._bool)
}

func TestParseAnswerBooleanFalse(t *testing.T) {
	t.Parallel()

	var value interface{}
	value = false
	parsedValue, err := parseAnswerValueBoolean(value)
	assert.Nil(t, err)
	assert.Equal(t, false, parsedValue._bool)
}

func TestTicketAnswerService_createNumberValue(t *testing.T) {
	t.Parallel()

	createTicketAnswerDTO := &domain.CreateTicketAnswerDTO{SubjectFieldID: uuid.New(), TicketID: uuid.New(), Value: 1}
	subjectField := &domain.SubjectField{
		ID:        createTicketAnswerDTO.SubjectFieldID,
		SubjectID: uuid.New(),
		Type:      domain.Number,
	}

	entity := NewTicketAnswerEntity()

	answer, err := entity.NewAnswer(createTicketAnswerDTO, subjectField)
	value := 1.0
	assert.Nil(t, err)
	assert.Equal(t, &domain.TicketAnswer{
		ID:       answer.ID,
		TicketID: createTicketAnswerDTO.TicketID,
		Value: &domain.TicketAnswerValue{
			Type:   domain.Number,
			Number: value,
		},
	}, answer)
}

func TestTicketAnswerService_createNilValue(t *testing.T) {
	t.Parallel()

	createTicketAnswerDTO := &domain.CreateTicketAnswerDTO{SubjectFieldID: uuid.New(), TicketID: uuid.New(), Value: nil}
	subjectField := &domain.SubjectField{
		ID:        createTicketAnswerDTO.SubjectFieldID,
		SubjectID: uuid.New(),
		Type:      domain.Number,
	}

	entity := NewTicketAnswerEntity()

	answer, err := entity.NewAnswer(createTicketAnswerDTO, subjectField)
	assert.Nil(t, err)
	assert.Equal(t, &domain.TicketAnswer{
		ID:       answer.ID,
		TicketID: createTicketAnswerDTO.TicketID,
		Value: &domain.TicketAnswerValue{
			Type: domain.Number,
		},
	}, answer)
}

func TestTicketAnswerService_createNilRequiredValue(t *testing.T) {
	t.Parallel()

	createTicketAnswerDTO := &domain.CreateTicketAnswerDTO{SubjectFieldID: uuid.New(), TicketID: uuid.New(), Value: nil}
	subjectField := &domain.SubjectField{
		ID:        createTicketAnswerDTO.SubjectFieldID,
		SubjectID: uuid.New(),
		Type:      domain.Number,
		Required:  true,
	}

	entity := NewTicketAnswerEntity()

	answer, err := entity.NewAnswer(createTicketAnswerDTO, subjectField)
	assert.Equal(t, domain.AnswerRequired, err)
	assert.Nil(t, answer)
}
