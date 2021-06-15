package ticket_answer

import (
	"github.com/google/uuid"
	"src/domain"
)

type ticketAnswerService struct {
	repository             domain.TicketAnswerRepository
	subjectFieldRepository domain.SubjectFieldRepository
}

func NewTicketAnswerService(repository domain.TicketAnswerRepository, subjectFieldRepository domain.SubjectFieldRepository) domain.TicketAnswerService {
	return &ticketAnswerService{repository: repository, subjectFieldRepository: subjectFieldRepository}
}

func (service *ticketAnswerService) New(answers []*domain.CreateTicketAnswerDTO) ([]*domain.TicketAnswer, error) {
	_answers := make([]*domain.TicketAnswer, len(answers))

	subjectFields, err := service.getSubjectFieldsByAnswers(answers)
	if err != nil {
		return nil, err
	}

	for i, answer := range answers {
		answer.SubjectField = subjectFields[answer.SubjectFieldID]
		_answer, err := New(answer)
		if err != nil {
			return nil, err
		}

		_answers[i] = _answer
	}

	return _answers, nil
}

func (service *ticketAnswerService) Create(answers []*domain.CreateTicketAnswerDTO) ([]*domain.TicketAnswer, error) {
	_answers, err := service.New(answers)
	if err != nil {
		return nil, err
	}

	err = service.repository.Save(_answers)
	if err != nil {
		return nil, err
	}

	return _answers, nil
}

func (service *ticketAnswerService) getSubjectFieldsByAnswers(answers []*domain.CreateTicketAnswerDTO) (map[uuid.UUID]*domain.SubjectField, error) {
	subjectFieldIdentifiers := make([]uuid.UUID, len(answers))
	for i, answer := range answers {
		subjectFieldIdentifiers[i] = answer.SubjectFieldID
	}

	return service.subjectFieldRepository.FetchByIdentifiers(subjectFieldIdentifiers)
}
