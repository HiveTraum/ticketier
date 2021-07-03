package subject_field

import (
	"github.com/google/uuid"
	"src/domain"
)

type subjectFieldEntity struct {
}

func NewSubjectFieldEntity() domain.SubjectFieldEntity {
	return &subjectFieldEntity{}
}

func (entity *subjectFieldEntity) NewField(DTO *domain.CreateSubjectFieldDTO) (*domain.SubjectField, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &domain.SubjectField{
		ID:             id,
		SubjectID:      DTO.SubjectID,
		Title:          DTO.Title,
		Required:       DTO.Required,
		ProgrammaticID: DTO.ProgrammaticID,
		Order:          DTO.Order,
		Type:           DTO.Type,
	}, nil
}

func (entity *subjectFieldEntity) NewFields(DTOs []*domain.CreateSubjectFieldDTO) ([]*domain.SubjectField, error) {
	fields := make([]*domain.SubjectField, len(DTOs))

	for i, DTO := range DTOs {
		field, err := entity.NewField(DTO)
		if err != nil {
			return nil, err
		}

		fields[i] = field
	}

	return fields, nil
}
