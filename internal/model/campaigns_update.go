package model

type UpdateCampaignRequestModel struct {
	CampaignName        string  `json:"campaign_name"`
	CampaignInitialDate string  `json:"campaign_initial_date"`
	CampaignFinalDate   string  `json:"campaign_final_date"`
	CampaignCurrency    string  `json:"campaign_currency"`
	CampaignBudget      float64 `json:"campaign_budget"`
}

type UpdateCampaignResponseModel struct {
	CampaignId          uint    `gorm:"campaign_id" json:"campaign_id"`
	CampaignName        string  `gorm:"campaign_name" json:"campaign_name"`
	CampaignBudget      float64 `gorm:"campaign_budget" json:"campaign_budget"`
	CampaignAgencyId    uint    `gorm:"campaign_agency_id" json:"campaign_agency_id"`
	CampaignManagerId   uint    `gorm:"campaign_manager_id" json:"campaign_manager_id"`
	CampaignCompanyId   uint    `gorm:"campaign_company_id" json:"campaign_company_id"`
	CampaignBundleId    uint    `gorm:"campaign_bundle_id" json:"campaign_bundle_id"`
	CampaignCurrency    string  `gorm:"campaign_currency" json:"campaign_currency"`
	CampaignInitialDate string  `gorm:"campaign_initial_date" json:"campaign_initial_date"`
	CampaignFinalDate   string  `gorm:"campaign_final_date" json:"campaign_final_date"`
}

func (data UpdateCampaignRequestModel) ValidateData() error {
	return nil
}
