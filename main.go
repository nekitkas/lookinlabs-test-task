package main

import (
	"embed"
	"log"
	"lookinlabs-test/config"
	"lookinlabs-test/controller"
	"lookinlabs-test/middleware"
	"lookinlabs-test/model"
	"lookinlabs-test/repository"
	"net/http"
	"os"

	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

var (
	//go:embed web
	build embed.FS
	g     errgroup.Group
)

func main() {
	pgConfig, err := config.NewPGSQLConfig()
	if err != nil {
		log.Fatal(err)
	}

	connection := repository.NewConnection(pgConfig)
	if err = connection.DB().AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}

	userController := controller.NewUserController(*connection)
	router := middleware.NewRouter(userController)
	frontendRouter := middleware.NewFrontedRouter(build)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	appPort := os.Getenv("API_PORT")
	frontedPort := os.Getenv("FRONTEND_PORT")

	apiServer := &http.Server{
		Addr:    ":" + appPort,
		Handler: corsMiddleware.Handler(router),
	}

	frontedServer := &http.Server{
		Addr:    ":" + frontedPort,
		Handler: frontendRouter,
	}

	log.Println("Starting API server on port", appPort)
	g.Go(func() error {
		return apiServer.ListenAndServe()
	})

	log.Println("Starting Frontend server on port", frontedPort)
	g.Go(func() error {
		return frontedServer.ListenAndServe()
	})

	if err = g.Wait(); err != nil {
		log.Fatal(err)
	}
}
