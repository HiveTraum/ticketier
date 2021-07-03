package subject

import (
	"github.com/google/uuid"
	"src/domain"
)

type subjectEntity struct {
}

func NewSubjectEntity() domain.SubjectEntity {
	return &subjectEntity{}
}

func (entity *subjectEntity) NewSubject(DTO *domain.CreateSubjectDTO) (*domain.Subject, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &domain.Subject{
		ID:        id,
		Title:     DTO.Title,
		ParentID:  DTO.ParentID,
		CreatedBy: DTO.CreatedBy,
	}, nil
}

func (entity *subjectEntity) NewSubjects(DTOs []*domain.CreateSubjectDTO) ([]*domain.Subject, error) {
	subjects := make([]*domain.Subject, len(DTOs))

	for i, DTO := range DTOs {
		subject, err := entity.NewSubject(DTO)
		if err != nil {
			return nil, err
		}

		subjects[i] = subject
	}

	return subjects, nil
}
