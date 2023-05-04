package companies

import (
	"database/sql"
	"go-backend-challenge/core-models-private-library/models"
)

type Company struct {
	models.BaseModelIDInt
	Name                   string         `gorm:"column:name" sql:"type:varchar(250);" json:"name"`
	Domain                 sql.NullString `gorm:"column:domain" sql:"type:varchar(250);" json:"domain"`
	Logo                   sql.NullString `gorm:"column:logo" sql:"type:varchar(250);" json:"logo"`
	Email                  sql.NullString `gorm:"column:email" sql:"type:varchar(250);" json:"email"`
	PhoneNumber            sql.NullString `gorm:"column:phone_number" sql:"type:varchar(250);" json:"phone_number"`
	PhoneNumberCountryCode sql.NullString `gorm:"column:phone_number_country_code" sql:"type:varchar(250);" json:"phone_number_country_code"`
	ContactName            sql.NullString `gorm:"column:contact_name" sql:"type:varchar(250);" json:"contact_name"`

	// RELATIONS
	AgencyId  uint `gorm:"column:agency_id" sql:"type:integer;" json:"agency_id"`
	ManagerId uint `gorm:"column:manager_id" sql:"type:integer;" json:"manager_id"`
}

func (model Company) TableName() string {
	return "companies"
}
