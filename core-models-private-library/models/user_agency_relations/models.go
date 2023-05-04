package user_agency_relations

import (
	"github.com/lib/pq"
)

const (
	FullAccessPermission = "FULL_ACCESS"
)

type UserAgencyRelation struct {
	// RELATIONS
	UserId      uint           `gorm:"primaryKey;autoIncrement:false;column:user_id;type:integer;unique_index:idx_user_agency_relation" json:"user_id"`
	AgencyId    uint           `gorm:"primaryKey;autoIncrement:false;column:agency_id;type:integer;unique_index:idx_user_agency_relation" json:"agency_id"`
	Permissions pq.StringArray `gorm:"column:permissions" sql:"type:TEXT[];" json:"permissions"`
}

func (model UserAgencyRelation) TableName() string {
	return "user_agency_relations"
}
