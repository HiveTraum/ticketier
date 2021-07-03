package subject_field

import (
	"context"
	"src/domain"
)

type subjectFieldService struct {
	entity     domain.SubjectFieldEntity
	repository domain.SubjectFieldRepository
}

func NewSubjectFieldService(entity domain.SubjectFieldEntity, repository domain.SubjectFieldRepository) domain.SubjectFieldService {
	return &subjectFieldService{entity: entity, repository: repository}
}

func (service *subjectFieldService) Create(ctx context.Context, DTOs []*domain.CreateSubjectFieldDTO) ([]*domain.SubjectField, error) {
	fields, err := service.entity.NewFields(DTOs)
	if err != nil {
		return nil, err
	}

	err = service.repository.Insert(ctx, fields)
	if err != nil {
		return nil, err
	}

	return fields, nil
}
