package model

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaign_creator_social_network_actions"
	utils "github.com/luisnquin/blind-creator-test-core-utils"
)

type CreateSocialNetworkActionRequest struct {
	campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions
}

func (data *CreateSocialNetworkActionRequest) Decode(r *http.Request) error {
	err := utils.DecodeData(r, data)
	if err != nil {
		return err
	}

	vars := mux.Vars(r)

	if id, err := strconv.Atoi(vars["campaign_id"]); err != nil {
		return fmt.Errorf("path variable 'campaign_id' is not a valid integer")
	} else {
		data.CampaignId = uint(id)
	}

	if id, err := strconv.Atoi(vars["social_network_id"]); err != nil {
		return fmt.Errorf("path variable 'social_network_id' is not a valid integer")
	} else {
		data.CreatorSocialNetworkId = uint(id)
	}

	if id, err := strconv.Atoi(vars["creator_id"]); err != nil {
		return fmt.Errorf("path variable 'creator_id' is not a valid integer")
	} else {
		data.CreatorId = uint(id)
	}

	return nil
}

func (data CreateSocialNetworkActionRequest) Validate() error {
	if data.CodeName == "" {
		return fmt.Errorf("'code_name' is required")
	}

	if data.CampaignId == 0 {
		return fmt.Errorf("'campaign_id' is required")
	}

	if data.CreatorId == 0 {
		return fmt.Errorf("'creator_id' is required")
	}

	if data.CreatorSocialNetworkId == 0 {
		return errors.New("'creator_social_network_id' is required")
	}

	if data.CostCurrency == "" {
		return errors.New("'cost_currency' is required")
	} else if !utils.Contains([]string{"COP", "EUR", "USD"}, data.CostCurrency) {
		return fmt.Errorf("'cost_currency' %s not supported", data.CostCurrency)
	}

	if data.CostPrice == 0 {
		return errors.New("you must provide a 'cost_price' greater than zero")
	}

	if data.AcceptedPrice == 0 {
		return errors.New("you must provide a 'accepted_price' greater than zero")
	}

	if data.BundlePrice == 0 {
		return errors.New("you must provide a 'bundle_price' greater than zero")
	}

	if data.Quantity == 0 {
		return errors.New("you must provide a 'quantity' greater than zero")
	}

	return nil
}
