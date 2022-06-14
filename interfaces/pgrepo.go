package interfaces

import (
	"clean-arch/domain"
	"database/sql"
)

type PgRepo struct {
	conn sql.DB
}

func (p *PgRepo) FindAll() []domain.FieldRecord {
	_, _ = p.conn.Exec("select * from field_record")
	return []domain.FieldRecord{}
}

func (p *PgRepo) FindById(id int) domain.FieldRecord {
	_, _ = p.conn.Exec("select * from field_record where id = %s", id)
	return domain.FieldRecord{}
}

func (p *PgRepo) UpdateTitleEligibility(recordId int, titleEligibility string) domain.FieldRecord {
	_, _ = p.conn.Exec("update field_record set title_eligibility = 'ELIGIBLE'")
	return domain.FieldRecord{}
}
