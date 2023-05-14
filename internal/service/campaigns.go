package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/repository"

	"github.com/jinzhu/gorm"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaign_creator_social_network_actions"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaigns"
	utils "github.com/luisnquin/blind-creator-test-core-utils"
)

type CampaignsServiceInterface interface {
	CreateCampaignServiceMethod()
	GetCampaignByIdServiceMethod()
	ListCampaignsServiceMethod()
	UpdateCampaignServiceMethod()
}

type CampaignsServiceStruct struct {
	repository.AgenciesDbRepository
}

func (c CampaignsServiceStruct) CreateCampaignServiceMethod(
	requestData model.CreateCampaignRequestModel,
) (
	model.CreateCampaignResponseModel,
	error,
) {
	campaignToCreate := campaigns.Campaign{}
	response := model.CreateCampaignResponseModel{}

	// validate manager
	manager, err := c.AgenciesDbRepository.GetUserById(requestData.CampaignManagerId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response, fmt.Errorf("an error happened trying to fetch user: %s", err.Error())
	}

	// format dates
	_initialDate, err := time.Parse("2006-01-02", requestData.CampaignInitialDate)
	if err == nil {
		campaignToCreate.InitialDate = sql.NullTime{Time: _initialDate, Valid: true}
	}
	_finalDate, err := time.Parse("2006-01-02", requestData.CampaignFinalDate)
	if err == nil {
		campaignToCreate.FinalDate = sql.NullTime{Time: _finalDate, Valid: true}
	}

	// create campaign
	campaignToCreate.Name = requestData.CampaignName
	campaignToCreate.Budget = requestData.CampaignBudget
	campaignToCreate.AgencyId = requestData.CampaignAgencyId
	campaignToCreate.ManagerId = manager.ID
	campaignToCreate.CompanyId = requestData.CampaignCompanyId
	campaignToCreate.BundleId = sql.NullInt64{Valid: true, Int64: int64(requestData.CampaignBundleId)}
	campaignToCreate.Currency = requestData.CampaignCurrency
	campaignToCreate, err = c.AgenciesDbRepository.CreateCampaign(campaignToCreate)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response, fmt.Errorf("an error happened trying to create agency: %s", err.Error())
	}

	// create actions
	for _, action := range requestData.CampaignCreatorSocialNetworkActions {
		actionToCreate := campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions{}
		actionToCreate.CodeName = action.CampaignCreatorSocialNetworkActionCodeName
		actionToCreate.Quantity = action.CampaignCreatorSocialNetworkActionQuantity
		actionToCreate.CostPrice = action.CampaignCreatorSocialNetworkActionCostPrice
		actionToCreate.CostCurrency = action.CampaignCreatorSocialNetworkActionCostCurrency
		actionToCreate.BundlePrice = action.CampaignCreatorSocialNetworkActionBundlePrice
		actionToCreate.AcceptedPrice = action.CampaignCreatorSocialNetworkActionAcceptedPrice
		actionToCreate.FinalContentStatus = campaign_creator_social_network_actions.ContentStatusPending
		actionToCreate.DraftContentStatus = campaign_creator_social_network_actions.ContentStatusPending
		actionToCreate.CampaignId = campaignToCreate.ID
		actionToCreate.CreatorId = action.CampaignCreatorSocialNetworkActionCreatorId
		actionToCreate.CreatorSocialNetworkId = action.CampaignCreatorSocialNetworkActionCreatorSocialNetworkId
		actionToCreate, err = c.AgenciesDbRepository.CreateCampaignCreatorSocialNetworkAction(actionToCreate)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return response, fmt.Errorf("an error happened trying to create agency: %s", err.Error())
		}
	}

	// define response
	response.CampaignId = campaignToCreate.ID

	return response, err
}

func (c CampaignsServiceStruct) SearchCampaignsByQuery(query string,
) ([]campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions, error) {
	query = strings.TrimSpace(strings.ToUpper(query))

	return c.AgenciesDbRepository.SearchCampaignsByQuery(query)
}

func (c CampaignsServiceStruct) GetCampaignByIdServiceMethod(
	campaignId uint,
) (
	model.GetCampaignDetailsResponseModel,
	error,
) {
	campaign, err := c.AgenciesDbRepository.GetCampaignById(campaignId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.GetCampaignDetailsResponseModel{}, err
	}

	manager, err := c.AgenciesDbRepository.GetUserById(campaign.ManagerId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.GetCampaignDetailsResponseModel{}, err
	}

	company, err := c.AgenciesDbRepository.GetCompanyById(campaign.CompanyId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.GetCampaignDetailsResponseModel{}, err
	}

	return model.GetCampaignDetailsResponseModel{
		CampaignID:          campaign.ID,
		CampaignCreatedAt:   campaign.CreatedAt.Time.Format("2006-01-02"),
		CampaignUpdatedAt:   campaign.UpdatedAt.Time.Format("2006-01-02"),
		CampaignName:        campaign.Name,
		CampaignInitialDate: campaign.FinalDate.Time.Format("2006-01-02"),
		CampaignFinalDate:   campaign.FinalDate.Time.Format("2006-01-02"),
		CampaignBudget:      campaign.Budget,
		CampaignCurrency:    campaign.Currency,
		CampaignAgencyID:    campaign.AgencyId,
		CampaignManagerID:   campaign.ManagerId,
		CampaignCompanyID:   campaign.CompanyId,
		ManagerEmail:        manager.Email,
		ManagerName:         manager.FirstName.String,
		CompanyName:         company.Name,
		CompanyEmail:        company.Email.String,
	}, err
}

func (c CampaignsServiceStruct) ListCampaignsServiceMethod(
	requestData model.ListCampaignsRequestModel,
) (
	utils.GormPaginationData,
	error,
) {
	return c.AgenciesDbRepository.ListAgencyCampaigns(
		utils.GormPaginationData{
			Limit: requestData.Limit,
			Page:  requestData.Page,
		},
		requestData.AgencyId,
		requestData.KeywordSearch,
	)
}

func (c CampaignsServiceStruct) UpdateCampaignServiceMethod(
	campaignId uint,
	requestData model.UpdateCampaignRequestModel,
) (
	model.UpdateCampaignResponseModel,
	error,
) {
	campaign, err := c.AgenciesDbRepository.GetCampaignById(
		campaignId,
	)
	if err != nil {
		return model.UpdateCampaignResponseModel{}, fmt.Errorf("an error happened trying to fetch campaign: %s", err.Error())
	}
	if campaign.ID == 0 {
		return model.UpdateCampaignResponseModel{}, fmt.Errorf("an error happened trying to fetch campaign")
	}

	if campaign.Name != requestData.CampaignName && requestData.CampaignName != "" {
		campaign.Name = requestData.CampaignName
	}

	if campaign.Budget != requestData.CampaignBudget && requestData.CampaignBudget != 0 {
		campaign.Budget = requestData.CampaignBudget
	}

	if campaign.Currency != requestData.CampaignCurrency && requestData.CampaignCurrency != "" {
		campaign.Currency = requestData.CampaignCurrency
	}

	campaign, err = c.AgenciesDbRepository.UpdateCampaign(campaign)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.UpdateCampaignResponseModel{}, fmt.Errorf("an error happened trying to update campaign: %s", err.Error())
	}

	response := model.UpdateCampaignResponseModel{}
	response.CampaignId = campaign.ID
	response.CampaignName = campaign.Name
	response.CampaignBudget = campaign.Budget
	response.CampaignAgencyId = campaign.AgencyId
	response.CampaignManagerId = campaign.ManagerId
	response.CampaignCompanyId = campaign.CompanyId
	response.CampaignCurrency = campaign.Currency
	response.CampaignInitialDate = campaign.InitialDate.Time.Format("2006-01-02")
	response.CampaignFinalDate = campaign.FinalDate.Time.Format("2006-01-02")
	return response, err
}
