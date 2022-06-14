package usecases

import "clean-arch/domain"

type FieldRecordService struct {
	RecordRepository domain.FieldRecordRepository
}

func (f *FieldRecordService) FindAll() []domain.FieldRecord {
	return f.RecordRepository.FindAll()
}

func (f *FieldRecordService) FindById(id int) domain.FieldRecord {
	return f.RecordRepository.FindById(id)
}
