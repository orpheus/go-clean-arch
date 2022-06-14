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
		},
	}
}
