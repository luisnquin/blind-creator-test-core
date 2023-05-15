package controller

import (
	"github.com/luisnquin/blind-creator/test-core/internal/repository"
	"github.com/luisnquin/blind-creator/test-core/internal/service"
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
