package creator_social_networks

import (
	"go-backend-challenge/core-models-private-library/models"
)

type CreatorSocialNetworkAccount struct {
	models.BaseModelIDInt
	SocialNetwork             string `gorm:"column:social_network" sql:"type:varchar(250);" json:"social_network"`
	Username                  string `gorm:"column:username" sql:"type:varchar(250);" json:"username"`
	ProfileURL                string `gorm:"column:profile_url" sql:"type:varchar(250);" json:"profile_url"`
	BlindCreatorAccountId     uint   `gorm:"column:blind_creator_account_id" json:"blind_creator_account_id"`
	BlindCreatorAccountPageId uint   `gorm:"column:blind_creator_account_page_id" json:"blind_creator_account_page_id"`
	IsVerified                bool   `gorm:"column:is_verified" json:"is_verified"`
	StatsMediaUrl             string `gorm:"column:stats_media_url" json:"stats_media_url"`
	CreatedBy                 string `gorm:"column:created_by" sql:"type:varchar(100);" json:"created_by"`
	// RELATIONS
	CreatorId uint `gorm:"column:creator_id" sql:"type:integer;" json:"creator_id"`
}

func (model CreatorSocialNetworkAccount) TableName() string {
	return "creator_social_networks"
}
