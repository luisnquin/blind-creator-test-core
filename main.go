package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/luisnquin/blind-creator/test-core/environment"
	"github.com/luisnquin/blind-creator/test-core/internal/controller"
	"github.com/luisnquin/blind-creator/test-core/internal/repository"
	router "github.com/luisnquin/blind-creator/test-core/internal/routes"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/luisnquin/blind-creator-test-core-models/migrations"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	// godotenv.Load()
	// SetFlags config Initializing log flags
	log.SetFlags(log.Flags() | log.Lshortfile)

	environment.InitializeEnv()

	if environment.ApplyMigrations {
		migrations.ApplyMigrations(repository.InitAgenciesDB())
	}

	// Negroni: web middleware-focused library. It is tiny, non-intrusive, and encourages use of net/http Handlers.
	app := negroni.Classic()

	/* Initialize Routes */
	routes := router.ApiRouter(controller.NewControl())
	app.UseHandler(routes)

	// CORS Settings
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: strings.Split(environment.CorsWhitelist, ","),
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowedHeaders: []string{"*"},
	})
	handler := corsOpts.Handler(app)

	isLocalEnvironment := os.Getenv("_LAMBDA_SERVER_PORT") == "" && os.Getenv("_AWS_LAMBDA_RUNTIME_API") == ""
	if !isLocalEnvironment {
		lambda.Start(httpadapter.New(handler).ProxyWithContext)
		return
	}

	port := environment.ServerPort

	fmt.Printf(fmt.Sprintf("http server is running on port :%d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
