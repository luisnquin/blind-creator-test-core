package model

import (
	"errors"
)

type (
	ListCampaignsResponseModel struct {
		*CampaignManagerCompany
		CampaignCreatorSocialNetworkActions []*CampaignCreatorSocialNetworkAction `gorm:"campaign_creator_social_network_actions" json:"campaign_creator_social_network_actions"`
	}

	ListCampaignsDTO struct {
		*CampaignManagerCompany
		*CampaignCreatorSocialNetworkAction
	}

	CampaignManagerCompany struct {
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

	CampaignCreatorSocialNetworkAction struct {
		ActionCodeName                     string  `gorm:"action_code_name" json:"action_code_name"`
		ActionQuantity                     uint    `gorm:"action_quantity" json:"action_quantity"`
		ActionCostPrice                    float64 `gorm:"action_cost_price" json:"action_cost_price"`
		ActionBundlePrice                  float64 `gorm:"action_bundle_price" json:"action_bundle_price"`
		ActionAcceptedPrice                float64 `gorm:"action_accepted_price" json:"action_accepted_price"`
		ActionCostCurrency                 string  `gorm:"action_cost_currency" json:"action_cost_currency"`
		ActionDraftContentStatus           string  `gorm:"action_draft_content_status" json:"action_draft_content_status"`
		ActionFinalContentStatus           string  `gorm:"action_final_content_status" json:"action_final_content_status"`
		ActionUploadDraftContentDate       string  `gorm:"action_upload_draft_content_date" json:"action_upload_draft_content_date"`
		ActionUploadFinalContentDate       string  `gorm:"action_upload_final_content_date" json:"action_upload_final_content_date"`
		ActionPaymentCondition             string  `gorm:"action_payment_condition" json:"action_payment_condition"`
		ActionCreatorId                    uint    `gorm:"action_creator_id" json:"action_creator_id"`
		ActionCreatorName                  string  `gorm:"action_creator_name" json:"action_creator_name"`
		ActionCreatorAvatar                string  `gorm:"action_creator_avatar" json:"action_creator_avatar"`
		ActionCreatorEmail                 string  `gorm:"action_creator_email" json:"action_creator_email"`
		ActionCreatorSocialNetworkId       uint    `gorm:"action_creator_social_network_id" json:"action_creator_social_network_id"`
		ActionCreatorSocialNetworkName     string  `gorm:"action_creator_social_network_name" json:"action_creator_social_network_name"`
		ActionCreatorSocialNetworkUsername string  `gorm:"action_creator_social_network_username" json:"action_creator_social_network_username"`
	}
)

type ListCampaignsRequestModel struct {
	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	AgencyId      uint   `json:"agency_id"`
	KeywordSearch string `json:"keyword_search"`
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
