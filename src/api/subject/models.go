package subject

import "src/domain"

type createSubjectDTOs []*domain.CreateSubjectDTO

func (DTOs createSubjectDTOs) fieldsCount() int {
	i := 0
	for _, subject := range DTOs {
		i = i + len(subject.Fields)
	}

	return i
}
