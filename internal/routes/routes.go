package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-backend-challenge/core-utils-private-library"
	"go-backend-challenge/environment"
	"go-backend-challenge/internal/controller"
	"net/http"
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

	fmt.Println("Available Routes:")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println(t, methods)
		return nil
	})

	return r
}
