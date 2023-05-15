package router

import (
	"fmt"
	"net/http"
	"os"

	"go-backend-challenge/environment"
	"go-backend-challenge/internal/controller"

	"github.com/gorilla/mux"
	utils "github.com/luisnquin/blind-creator-test-core-utils"
)

func ApiRouter(c controller.CustomControllerStruct) *mux.Router {
	r := mux.NewRouter()

	basicAuth := utils.BasicAuth(
		environment.BasicAuthUsername,
		environment.BasicAuthPassword,
	)
	r.Use(basicAuth)

	r.NotFoundHandler = http.HandlerFunc(utils.NotFoundHandler)

	r.HandleFunc(
		"/v1/campaigns",
		c.Campaigns.CreateCampaignControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/campaigns/list",
		c.Campaigns.ListCampaignsControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/campaigns/{campaign_id:[0-9]+}",
		c.Campaigns.GetCampaignByIdControllerMethod,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/v1/campaigns/{campaign_id:[0-9]+}",
		c.Campaigns.UpdateCampaignControllerMethod,
	).Methods(http.MethodPut)

	r.HandleFunc(
		"/v1/campaigns/search",
		c.Campaigns.SearchCampaignControllerMethod,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/v1/campaigns/{campaign_id:[0-9]+}/creators/{creator_id:[0-9]+}/social-networks/{social_network_id:[0-9]+}/action",
		c.Campaigns.CreateCampaignSocialNetworkActionControllerMethod,
	).Methods(http.MethodPost)

	fmt.Println("Available Routes:")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err := route.GetMethods()
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout, t, methods)

		return nil
	})

	return r
}
