package usecases

import "clean-arch/domain"

type FieldRecordService struct {
	RecordRepository domain.FieldRecordRepository

	// Option 1:
	EsRecordRepository domain.FieldRecordRepository

	// Option 2:
	RecordSearchRepo RecordSearchRepo
}

type RecordSearchRepo interface {
	FindAll() []domain.FieldRecord
}

func (f *FieldRecordService) FindAll() []domain.FieldRecord {
	return f.RecordRepository.FindAll()

	// option 2:
	// return f.RecordSearchRepo.FindAll()
}

func (f *FieldRecordService) FindById(id int) domain.FieldRecord {
	return f.RecordRepository.FindById(id)
}

// Option 1:
func (f *FieldRecordService) EsFindAll() []domain.FieldRecord {
	return f.EsRecordRepository.FindAll()
}
