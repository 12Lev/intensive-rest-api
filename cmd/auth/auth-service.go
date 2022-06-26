package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
	_ "intensive-rest-api/docs"
	"intensive-rest-api/pkg/routes"
	"intensive-rest-api/pkg/utils"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Outdoor API
// @version 2.0
// @description Документация
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
// @BasePath /
func main() {
	port := utils.GetEnv("PORT", ":8000")
	routers := mux.NewRouter()
	routers.PathPrefix("/api/v1/docs").Handler(httpSwagger.WrapHandler)
	routers.HandleFunc("/api/v1/add-user", routes.PostAddUser).Methods("POST", "OPTIONS")
	routers.HandleFunc("/api/v1/token", routes.PostAuth).Methods("POST", "OPTIONS")
	routers.HandleFunc("/api/v1/get-data", routes.AuthRequired(routes.GetOnlyAuthorized)).Methods("GET", "OPTIONS")
	s := &http.Server{
		Addr:         port,
		WriteTimeout: time.Minute * 5,
		ReadTimeout:  time.Minute * 5,
		IdleTimeout:  time.Second * 1,
		Handler:      routers,
	}

	go func() {
		fmt.Print("Base... [http://localhost", port, "/api/v1", "]\n")
		fmt.Print("Base... [http://localhost", port, "/api/v1/docs/index.html", "]\n")
		utils.CheckErr(s.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
