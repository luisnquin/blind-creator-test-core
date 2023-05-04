package models

//Imports
import (
	"database/sql"
)

type BaseModelIDInt struct {
	ID        uint          `gorm:"primary_key;" sql:"serial; int4" json:"id"`
	CreatedAt sql.NullTime  `gorm:"created_at;" json:"created_at"`
	UpdatedAt sql.NullTime  `gorm:"updated_at;" json:"updated_at"`
	DeletedAt *sql.NullTime `gorm:"deleted_at;" sql:"index" json:"deleted_at"`
}
