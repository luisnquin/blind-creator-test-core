package controller

import (
	"go-backend-challenge/internal/repository"
	"go-backend-challenge/internal/service"
)

type CustomControllerStruct struct {
	Campaigns CampaignsControllerInterface
}

func NewControl() CustomControllerStruct {
	agenciesDBConnection := repository.InitAgenciesDB()
	return CustomControllerStruct{
		Campaigns: CampaignsControllerStruct{
			CampaignsServiceStruct: service.CampaignsServiceStruct{
				AgenciesDbRepository: repository.AgenciesDbRepository{
					DB: agenciesDBConnection,
				},
			},
		},
	}
}
