package model

import (
	"errors"
)

type CreateCampaignRequestModel struct {
	CampaignManagerId                   uint                           `json:"campaign_manager_id"`
	CampaignBundleId                    uint                           `json:"campaign_bundle_id"`
	CampaignCompanyId                   uint                           `json:"campaign_company_id"`
	CampaignAgencyId                    uint                           `json:"campaign_agency_id"`
	CampaignName                        string                         `json:"campaign_name"`
	CampaignInitialDate                 string                         `json:"campaign_initial_date"`
	CampaignFinalDate                   string                         `json:"campaign_final_date"`
	CampaignCurrency                    string                         `json:"campaign_currency"`
	CampaignBudget                      float64                        `json:"campaign_budget"`
	CampaignCreatorSocialNetworkActions []CreateCampaignCreatorActions `json:"campaign_creator_social_network_actions"`
}

type CreateCampaignCreatorActions struct {
	CampaignCreatorSocialNetworkActionCreatorId              uint    `json:"campaign_creator_social_network_action_creator_id"`
	CampaignCreatorSocialNetworkActionCreatorSocialNetworkId uint    `json:"campaign_creator_social_network_action_creator_social_network_id"`
	CampaignCreatorSocialNetworkActionCodeName               string  `json:"campaign_creator_social_network_action_code_name"`
	CampaignCreatorSocialNetworkActionQuantity               int     `json:"campaign_creator_social_network_action_quantity"`
	CampaignCreatorSocialNetworkActionCostPrice              float64 `json:"campaign_creator_social_network_action_cost_price"`
	CampaignCreatorSocialNetworkActionCostCurrency           string  `json:"campaign_creator_social_network_action_cost_currency"`
	CampaignCreatorSocialNetworkActionBundlePrice            float64 `json:"campaign_creator_social_network_action_bundle_price"`
	CampaignCreatorSocialNetworkActionAcceptedPrice          float64 `json:"campaign_creator_social_network_action_accepted_price"`
	CampaignCreatorSocialNetworkActionUploadDraftContentDate string  `json:"campaign_creator_social_network_action_upload_draft_content_date"`
	CampaignCreatorSocialNetworkActionUploadFinalContentDate string  `json:"campaign_creator_social_network_action_upload_final_content_date"`
}

func (data CreateCampaignRequestModel) ValidateData() error {
	if data.CampaignName == "" {
		return errors.New("agency_name is required")
	} else if data.CampaignManagerId == 0 {
		return errors.New("manager_id is required")
	} else if data.CampaignAgencyId == 0 {
		return errors.New("agency_id is required")
	} else if data.CampaignCompanyId == 0 {
		return errors.New("company_id is required")
	} else if data.CampaignFinalDate == "" {
		return errors.New("campaign_final_date is required")
	} else if data.CampaignInitialDate == "" {
		return errors.New("campaign_initial_date is required")
	} else if data.CampaignCurrency == "" {
		return errors.New("campaign_currency is required")
	} else if data.CampaignName == "" {
		return errors.New("campaign_name is required")
	} else if len(data.CampaignCreatorSocialNetworkActions) == 0 {
		return errors.New("campaign_creator_social_network_actions is required")
	}

	// todo: Definir el resto de datos requeridos

	return nil
}

type CreateCampaignResponseModel struct {
	CampaignId uint `json:"campaign_id"`
}
