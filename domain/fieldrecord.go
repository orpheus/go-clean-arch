package domain

// FieldRecordRepository describes how my entity is going to be persisted and managed.
type FieldRecordRepository interface {
	FindAll() []FieldRecord
	FindById(id int) FieldRecord
	UpdateTitleEligibility(recordId int, titleEligibility string) FieldRecord
}

type TitleEligibility int

const (
	TITLE_READY  TitleEligibility = 0
	ELIGIBLE                      = 1
	NOT_ELIGIBLE                  = 2
	PENDING                       = 3
)

// FieldRecord Entity
type FieldRecord struct {
	Id               int
	TitleEligibility TitleEligibility
	Owner            string
	AmountPaid       int
}

type FieldRecordPolicy struct{}

func (f FieldRecordPolicy) IsValidOwner(fr FieldRecord) bool {
	return fr.Owner != ""
}

func (f FieldRecordPolicy) HasPaidMinAmount(fr FieldRecord) bool {
	return fr.AmountPaid > 248924
}

func (f *FieldRecord) UpdateTitleEligibility() {
	policy := &FieldRecordPolicy{}

	if policy.IsValidOwner(*f) {
		f.TitleEligibility = ELIGIBLE
	} else if policy.HasPaidMinAmount(*f) {
		f.TitleEligibility = TITLE_READY
	} else if !policy.HasPaidMinAmount(*f) && !policy.IsValidOwner(*f) {
		f.TitleEligibility = NOT_ELIGIBLE
	} else {
		f.TitleEligibility = PENDING
	}
}
