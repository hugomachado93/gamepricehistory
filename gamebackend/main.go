package main

import (
	"fmt"
	"gamepricehistory/internal/api"
	"gamepricehistory/internal/crons"
	"gamepricehistory/internal/infrastructure"
	"gamepricehistory/internal/infrastructure/database"
	"gamepricehistory/internal/repository"
	"gamepricehistory/internal/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	sr := repository.NewSteamRepository(database.GetDbConnection())

	sapi := infrastructure.NewSteamApi()

	gs := services.NewGameService(sr)

	ss := services.NewSteamService(sr, sapi)

	cronService := crons.CreateCronService(ss)

	cronService.RunFetchSteam()

	r := mux.NewRouter()

	api.CreateApis(r, gs)

	fmt.Println("Serving server on port 8080")

	h := cors.Default().Handler(r)
	http.ListenAndServe(":8080", h)
}
