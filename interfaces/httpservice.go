package interfaces

import "clean-arch/domain"

type FieldRecordController struct {
	Interactor FieldRecordInteractor
}

type FieldRecordInteractor interface {
	FindAll() []domain.FieldRecord
	FindById(id int) domain.FieldRecord
}

func (c FieldRecordController) RegisterRoutes() map[string]map[string]interface{} {
	routes := make(map[string]map[string]interface{})

	routes["/api/fieldrecord"] = map[string]interface{}{
		"GET": FieldRecordController.FindAll,
	}

	routes["/api/fieldrecord/{id}"] = map[string]interface{}{
		"GET": FieldRecordController.FindById,
	}

	return routes
}

func (c *FieldRecordController) FindAll() []domain.FieldRecord {
	return c.Interactor.FindAll()
}

func (c *FieldRecordController) FindById(id int) domain.FieldRecord {
	return c.Interactor.FindById(id)
}
