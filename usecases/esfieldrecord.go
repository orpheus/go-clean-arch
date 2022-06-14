package usecases

import "clean-arch/domain"

type ESFieldRecordService struct {
	EsRepo domain.FieldRecordRepository
}

func (f *ESFieldRecordService) FindAll() []domain.FieldRecord {
	return f.EsRepo.FindAll()
}

func (f *ESFieldRecordService) FindById(id int) domain.FieldRecord {
	return f.EsRepo.FindById(id)
}
