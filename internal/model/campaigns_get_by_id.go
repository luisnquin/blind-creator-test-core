package model

type GetCampaignDetailsResponseModel struct {
	CampaignID          uint    `json:"campaign_id"`
	CampaignCreatedAt   string  `json:"campaign_created_at"`
	CampaignUpdatedAt   string  `json:"campaign_updated_at"`
	CampaignName        string  `json:"campaign_name"`
	CampaignInitialDate string  `json:"campaign_initial_date"`
	CampaignFinalDate   string  `json:"campaign_final_date"`
	CampaignBudget      float64 `json:"campaign_budget"`
	CampaignCurrency    string  `json:"campaign_currency"`
	CampaignAgencyID    uint    `json:"campaign_agency_id"`
	CampaignManagerID   uint    `json:"campaign_manager_id"`
	CampaignCompanyID   uint    `json:"campaign_company_id"`
	ManagerEmail        string  `json:"manager_email"`
	ManagerName         string  `json:"manager_name"`
	CompanyName         string  `json:"company_name"`
	CompanyEmail        string  `json:"company_email"`
}
