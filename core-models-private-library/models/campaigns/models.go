package campaigns

import (
	"database/sql"
	"go-backend-challenge/core-models-private-library/models"
)

type Campaign struct {
	models.BaseModelIDInt
	Name        string       `gorm:"column:name" sql:"type:varchar(250);" json:"name"`
	InitialDate sql.NullTime `gorm:"column:initial_date" sql:"type:date;" json:"initial_date"`
	FinalDate   sql.NullTime `gorm:"column:final_date" sql:"type:date;"  json:"final_date"`
	Budget      float64      `gorm:"column:budget" json:"budget"`
	Currency    string       `gorm:"column:currency" sql:"type:varchar(250);" json:"currency"`

	// RELATIONS
	AgencyId  uint          `gorm:"column:agency_id" sql:"type:integer;" json:"agency_id"`
	ManagerId uint          `gorm:"column:manager_id" sql:"type:integer;" json:"manager_id"`
	CompanyId uint          `gorm:"column:company_id" sql:"type:integer;" json:"company_id"`
	BundleId  sql.NullInt64 `gorm:"column:bundle_id" sql:"type:integer;" json:"bundle_id"`
}

func (model Campaign) TableName() string {
	return "campaigns"
}
