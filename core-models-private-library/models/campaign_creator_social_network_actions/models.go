package campaign_creator_social_network_actions

import (
	"database/sql"
	"go-backend-challenge/core-models-private-library/models"
)

const (
	ContentStatusPending   = "PENDING"
	ContentStatusCompleted = "COMPLETED"
)

type CampaignCreatorSocialNetworkActions struct {
	models.BaseModelIDInt
	CodeName               string       `gorm:"column:code_name" sql:"type:varchar(250);" json:"code_name"`
	Quantity               int          `gorm:"column:quantity"  json:"quantity"`
	CostPrice              float64      `gorm:"column:cost_price"  json:"cost_price"`
	BundlePrice            float64      `gorm:"column:bundle_price"  json:"bundle_price"`
	AcceptedPrice          float64      `gorm:"column:accepted_price"  json:"accepted_price"`
	DraftContentStatus     string       `gorm:"column:draft_content_status" sql:"type:varchar(250);" json:"draft_content_status"`
	FinalContentStatus     string       `gorm:"column:final_content_status" sql:"type:varchar(250);" json:"final_content_status"`
	UploadDraftContentDate sql.NullTime `gorm:"column:upload_draft_content_date" sql:"type:date;" json:"upload_draft_content_date"`
	UploadFinalContentDate sql.NullTime `gorm:"column:upload_final_content_date" sql:"type:date;" json:"upload_final_content_date"`
	PaymentCondition       string       `gorm:"column:payment_condition" sql:"type:varchar(250);" json:"payment_condition"`
	CostCurrency           string       `gorm:"column:cost_currency" sql:"type:varchar(250);" json:"cost_currency"`

	// RELATIONS
	CampaignId                  uint          `gorm:"column:campaign_id" sql:"type:integer;" json:"campaign_id"`
	CreatorId                   uint          `gorm:"column:creator_id" sql:"type:integer;" json:"creator_id"`
	CreatorSocialNetworkId      uint          `gorm:"column:creator_social_network_id" sql:"type:integer;" json:"creator_social_network_id"`
	CampaignPaymentObligationId sql.NullInt64 `gorm:"column:campaign_payment_obligation_id" json:"campaign_payment_obligation_id"`
	CampaignChargeObligationId  sql.NullInt64 `gorm:"column:campaign_charge_obligation_id" json:"campaign_charge_obligation_id"`
}

func (model CampaignCreatorSocialNetworkActions) TableName() string {
	return "campaign_creator_social_network_actions"
}
