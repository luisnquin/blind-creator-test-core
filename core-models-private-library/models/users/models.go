package users

import (
	"database/sql"
	"go-backend-challenge/core-models-private-library/models"
)

type User struct {
	models.BaseModelIDInt
	Email                  string         `gorm:"column:email" sql:"type:varchar(250); not null; unique_index" json:"email"`
	FirstName              sql.NullString `gorm:"column:first_name" sql:"type:varchar(250);" json:"first_name"`
	LastName               sql.NullString `gorm:"column:last_name" sql:"type:varchar(250);" json:"last_name"`
	PhoneNumber            sql.NullString `gorm:"column:phone_number" sql:"type:varchar(250);" json:"phone_number"`
	PhoneNumberCountryCode sql.NullString `gorm:"column:phone_number_country_code" sql:"type:varchar(250);" json:"phone_number_country_code"`
	OnBoardingCompleted    bool           `gorm:"column:on_boarding_completed" sql:"type:boolean" json:"on_boarding_completed"`
	// RELATIONS
	CognitoId          string         `gorm:"column:cognito_id" sql:"type:varchar(250); not null; unique_index" json:"cognito_id"`
	DefaultProfileId   uint           `gorm:"column:default_profile_id" sql:"type:integer;" json:"default_profile_id"`
	DefaultProfileName sql.NullString `gorm:"column:default_profile_name" sql:"type:varchar(250); " json:"default_profile_name"`
}

func (model User) TableName() string {
	return "users"
}
