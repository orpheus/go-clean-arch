package usecases

import "clean-arch/domain"

type FieldRecordService struct {
	RecordRepository domain.FieldRecordRepository

	// Option 1-0:
	EsRecordRepository domain.FieldRecordRepository

	// Option 2: Create a search specific interface
	RecordSearchRepo RecordSearchRepo
}

// RecordSearchRepo Option 2: Create a search specific interface
type RecordSearchRepo interface {
	FindAll() []domain.FieldRecord
}

func (f *FieldRecordService) FindAll() []domain.FieldRecord {
	return f.RecordRepository.FindAll()

	// Option 1-1: Swap out the existing search for a full new Repo implementation
	// return f.EsRecordRepository.FindAll()

	// Option 2: Swap out the existing search for a lightweight new Repo implementation
	// return f.RecordSearchRepo.FindAll()
}

func (f *FieldRecordService) FindById(id int) domain.FieldRecord {
	return f.RecordRepository.FindById(id)
}

// EsFindAll Option 1-0: Create a separate FindAll function
func (f *FieldRecordService) EsFindAll() []domain.FieldRecord {
	return f.EsRecordRepository.FindAll()
}
