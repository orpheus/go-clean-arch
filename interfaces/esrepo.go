package interfaces

import (
	"clean-arch/domain"
	"database/sql"
)

type EsRepo struct {
	conn sql.DB
}

func (p *EsRepo) FindAll() []domain.FieldRecord {
	_, _ = p.conn.Exec("select * from field_record")
	return []domain.FieldRecord{}
}

func (p *EsRepo) FindById(id int) domain.FieldRecord {
	_, _ = p.conn.Exec("select * from field_record where id = %s", id)
	return domain.FieldRecord{}
}

func (p *EsRepo) UpdateTitleEligibility(recordId int, titleEligibility string) domain.FieldRecord {
	_, _ = p.conn.Exec("update field_record set title_eligibility = 'ELIGIBLE'")
	return domain.FieldRecord{}
}
