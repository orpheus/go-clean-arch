package infrastructure

import (
	"clean-arch/interfaces"
	"clean-arch/usecases"
)

type HttpServer struct {
}

func BuildServer() {
	_ = &HttpServer{}
	_ = &interfaces.FieldRecordController{
		Interactor: &usecases.FieldRecordService{
			RecordRepository: &interfaces.PgRepo{},
			// Option 1-0: implement a full ES Repository
			EsRecordRepository: &interfaces.EsRepo{},
			// Option 2: Implement a Specialized ES Repository
			RecordSearchRepo: &interfaces.EsRepo{},
		},
	}
}
