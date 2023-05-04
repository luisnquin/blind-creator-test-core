package migrations

import (
	"github.com/jinzhu/gorm"
	"go-backend-challenge/core-models-private-library/models/agencies"
	"go-backend-challenge/core-models-private-library/models/campaign_creator_social_network_actions"
	"go-backend-challenge/core-models-private-library/models/campaigns"
	"go-backend-challenge/core-models-private-library/models/companies"
	"go-backend-challenge/core-models-private-library/models/creator_social_networks"
	"go-backend-challenge/core-models-private-library/models/creators"
	"go-backend-challenge/core-models-private-library/models/user_agency_relations"
	"go-backend-challenge/core-models-private-library/models/users"
	"go-backend-challenge/environment"
	"log"
)

func ApplyMigrations(dbConnection *gorm.DB) {

	if !environment.ApplyMigrations {
		return
	}

	dbConnection.LogMode(true)

	// MODELS
	log.Println("users...")
	dbConnection.AutoMigrate(&users.User{})

	log.Println("agencies...")
	dbConnection.AutoMigrate(&agencies.Agency{})

	log.Println("user_agency_relations...")
	dbConnection.AutoMigrate(&user_agency_relations.UserAgencyRelation{})

	log.Println("creators...")
	dbConnection.AutoMigrate(&creators.Creator{})

	log.Println("creator_social_networks...")
	dbConnection.AutoMigrate(&creator_social_networks.CreatorSocialNetworkAccount{})

	log.Println("companies...")
	dbConnection.AutoMigrate(&companies.Company{})

	log.Println("campaigns...")
	dbConnection.AutoMigrate(&campaigns.Campaign{})

	log.Println("campaign_creator_social_network_actions...")
	dbConnection.AutoMigrate(&campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions{})

	// RELATIONS

	log.Println("user_agency_relations...")
	dbConnection.Model(&user_agency_relations.UserAgencyRelation{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("agency_id", "agencies(id)", "CASCADE", "CASCADE")

	log.Println("creator_social_networks...")
	dbConnection.Model(&creator_social_networks.CreatorSocialNetworkAccount{}).
		AddForeignKey("creator_id", "creators(id)", "CASCADE", "CASCADE")

	log.Println("companies...")
	dbConnection.Model(&companies.Company{}).
		AddForeignKey("agency_id", "agencies(id)", "CASCADE", "CASCADE").
		AddForeignKey("manager_id", "users(id)", "CASCADE", "CASCADE")

	log.Println("campaigns...")
	dbConnection.Model(&campaigns.Campaign{}).
		AddForeignKey("agency_id", "agencies(id)", "CASCADE", "CASCADE").
		AddForeignKey("manager_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")

	log.Println("campaign_creator_social_network_actions ...")
	dbConnection.Model(&campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions{}).
		AddForeignKey("campaign_id", "campaigns(id)", "CASCADE", "CASCADE").
		AddForeignKey("creator_id", "creators(id)", "CASCADE", "CASCADE").
		AddForeignKey("creator_social_network_id", "creator_social_networks(id)", "CASCADE", "CASCADE")

}
