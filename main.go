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
)

var (
	//go:embed web
	build embed.FS
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
	router := middleware.NewRouter(userController, build)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	appPort := os.Getenv("APP_PORT")

	log.Println("Server running on port " + appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, corsMiddleware.Handler(router)))
}
