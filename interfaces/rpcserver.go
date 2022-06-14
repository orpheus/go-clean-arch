package interfaces

import "clean-arch/domain"

type FieldRecordRPC struct {
	Interactor FieldRecordInteractor
}

type RecordRPCInteractor interface {
	FindAll() []domain.FieldRecord
}

func (c FieldRecordController) RegisterRPC() map[string]map[string]interface{} {
	routes := make(map[string]map[string]interface{})

	routes["/api/fieldrecord"] = map[string]interface{}{
		"GET": FieldRecordController.FindAll,
	}

	return routes
}

func (c *FieldRecordRPC) FindAll() []domain.FieldRecord {
	return c.Interactor.FindAll()
}
