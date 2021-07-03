package subject

import (
	"context"
	"src/domain"
)

type subjectService struct {
	entity           domain.SubjectEntity
	repository       domain.SubjectRepository
	fieldsEntity     domain.SubjectFieldEntity
	fieldsRepository domain.SubjectFieldRepository
}

func NewSubjectService(entity domain.SubjectEntity, repository domain.SubjectRepository, fieldsEntity domain.SubjectFieldEntity, fieldsRepository domain.SubjectFieldRepository) domain.SubjectService {
	return &subjectService{entity: entity, repository: repository, fieldsEntity: fieldsEntity, fieldsRepository: fieldsRepository}
}

func (service *subjectService) List(ctx context.Context) ([]*domain.Subject, error) {
	return service.repository.List(ctx)
}

func (service *subjectService) Create(ctx context.Context, DTOs []*domain.CreateSubjectDTO) ([]*domain.Subject, error) {
	subjects := make([]*domain.Subject, len(DTOs))
	fields := make([]*domain.SubjectField, createSubjectDTOs(DTOs).fieldsCount())
	fieldsCount := 0

	for i, DTO := range DTOs {
		subject, err := service.entity.NewSubject(DTO)
		if err != nil {
			return nil, err
		}

		for _, field := range DTO.Fields {
			field.SubjectID = subject.ID
		}

		_fields, err := service.fieldsEntity.NewFields(DTO.Fields)
		if err != nil {
			return nil, err
		}

		for _, field := range _fields {
			fields[fieldsCount] = field
			fieldsCount++
		}

		subject.Fields = _fields
		subjects[i] = subject
	}

	transaction, err := service.repository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = service.repository.InsertInTransaction(ctx, subjects, transaction)
	if err != nil {
		return nil, err
	}

	err = service.fieldsRepository.InsertInTransaction(ctx, fields, transaction)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
