package infrastructure

import (
	"clean-arch/interfaces"
	"clean-arch/usecases"
)

type RPCServer struct {
}

func BuildRPCServer() {
	_ = &RPCServer{}
	_ = &interfaces.FieldRecordRPC{
		Interactor: &usecases.FieldRecordService{
			RecordSearchRepo: &interfaces.EsRepo{},
		},
	}
}
