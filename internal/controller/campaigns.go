package controller

import (
	"net/http"
	"strconv"
	"strings"

	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/service"

	"github.com/gorilla/mux"
	utils "github.com/luisnquin/blind-creator-test-core-utils"
)

type CampaignsControllerStruct struct {
	service.CampaignsServiceStruct
}

type CampaignsControllerInterface interface {
	CreateCampaignControllerMethod(http.ResponseWriter, *http.Request)
	GetCampaignByIdControllerMethod(http.ResponseWriter, *http.Request)
	ListCampaignsControllerMethod(http.ResponseWriter, *http.Request)
	UpdateCampaignControllerMethod(http.ResponseWriter, *http.Request)
	SearchCampaignControllerMethod(http.ResponseWriter, *http.Request)
}

func (c CampaignsControllerStruct) CreateCampaignControllerMethod(w http.ResponseWriter, r *http.Request) {
	var err error
	requestData := model.CreateCampaignRequestModel{}

	err = utils.DecodeData(r, &requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	err = requestData.ValidateData()
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	response, err := c.CreateCampaignServiceMethod(requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": response})
	return
}

func (c CampaignsControllerStruct) GetCampaignByIdControllerMethod(w http.ResponseWriter, r *http.Request) {
	campaignId, _ := strconv.Atoi(mux.Vars(r)["campaign_id"])

	campaign, err := c.GetCampaignByIdServiceMethod(uint(campaignId))
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": campaign})
}

func (c CampaignsControllerStruct) ListCampaignsControllerMethod(w http.ResponseWriter, r *http.Request) {
	var err error
	requestData := model.ListCampaignsRequestModel{}

	err = utils.DecodeData(r, &requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	err = requestData.ValidateData()
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	response, err := c.ListCampaignsServiceMethod(requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.SuccessPaginationResponse(w, http.StatusOK, response)
	return
}

func (c CampaignsControllerStruct) UpdateCampaignControllerMethod(w http.ResponseWriter, r *http.Request) {
	var err error
	var campaignId uint
	requestData := model.UpdateCampaignRequestModel{}

	_campaignId, err := strconv.Atoi(mux.Vars(r)["campaign_id"])
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}
	campaignId = uint(_campaignId)

	err = utils.DecodeData(r, &requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	err = requestData.ValidateData()
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	response, err := c.UpdateCampaignServiceMethod(campaignId, requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": response})
	return
}

func (c CampaignsControllerStruct) SearchCampaignControllerMethod(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	q := query.Get("q")

	if strings.TrimSpace(q) == "" {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": "you must provide a search word"})

		return
	}

	results, err := c.SearchCampaignsByQuery(q)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})

		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{
		"status":  "SUCCESS",
		"message": nil,
		"data":    &results,
	})
}
