package agencies

import (
	"database/sql"
	"go-backend-challenge/core-models-private-library/models"
)

type Agency struct {
	models.BaseModelIDInt
	Name                   string         `gorm:"column:name" sql:"type:varchar(250);" json:"name"`
	Domain                 sql.NullString `gorm:"column:domain" sql:"type:varchar(250);" json:"domain"`
	Logo                   sql.NullString `gorm:"column:logo" sql:"type:varchar(250);" json:"logo"`
	Email                  sql.NullString `gorm:"column:email" sql:"type:varchar(250);" json:"email"`
	PhoneNumber            sql.NullString `gorm:"column:phone_number" sql:"type:varchar(250);" json:"phone_number"`
	PhoneNumberCountryCode sql.NullString `gorm:"column:phone_number_country_code" sql:"type:varchar(250);" json:"phone_number_country_code"`
}

func (model Agency) TableName() string {
	return "agencies"
}
