package model

import (
	"errors"
)

type ListCampaignsRequestModel struct {
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	AgencyId      uint   `json:"agency_id"`
	KeywordSearch string `json:"keyword_search"`
}

type ListCampaignsResponseModel struct {
	CampaignID          uint    `gorm:"column:campaign_id" json:"campaign_id"`
	CampaignCreatedAt   string  `gorm:"column:campaign_created_at" json:"campaign_created_at"`
	CampaignUpdatedAt   string  `gorm:"column:campaign_updated_at" json:"campaign_updated_at"`
	CampaignDeletedAt   string  `gorm:"column:campaign_deleted_at" json:"campaign_deleted_at,omitempty"`
	CampaignName        string  `gorm:"column:campaign_name" json:"campaign_name"`
	CampaignInitialDate string  `gorm:"column:campaign_initial_date" json:"campaign_initial_date"`
	CampaignFinalDate   string  `gorm:"column:campaign_final_date" json:"campaign_final_date"`
	CampaignBudget      float64 `gorm:"column:campaign_budget" json:"campaign_budget"`
	CampaignCurrency    string  `gorm:"column:campaign_currency" json:"campaign_currency"`
	CampaignAgencyID    uint    `gorm:"column:campaign_agency_id" json:"campaign_agency_id"`
	CampaignManagerID   uint    `gorm:"column:campaign_manager_id" json:"campaign_manager_id"`
	CampaignCompanyID   uint    `gorm:"column:campaign_company_id" json:"campaign_company_id"`
	CampaignBundleID    uint    `gorm:"column:campaign_bundle_id" json:"campaign_bundle_id"`
	ManagerEmail        string  `gorm:"column:manager_email" json:"manager_email"`
	ManagerName         string  `gorm:"column:manager_name" json:"manager_name"`
	CompanyName         string  `gorm:"column:company_name" json:"company_name"`
	CompanyEmail        string  `gorm:"column:company_email" json:"company_email"`
}

func (data ListCampaignsRequestModel) ValidateData() error {
	if data.Limit == 0 {
		return errors.New("limit is required")
	} else if data.Page == 0 {
		return errors.New("page is required")
	} else if data.Page == 0 {
		return errors.New("user_id is required")
	}
	return nil
}
