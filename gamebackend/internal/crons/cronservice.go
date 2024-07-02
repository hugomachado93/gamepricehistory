package crons

import (
	"gamepricehistory/internal/services"
)

type CronsService struct {
	steamService *services.SteamService
}

func CreateCronService(steamService *services.SteamService) *CronsService {
	return &CronsService{steamService: steamService}
}

func (cs *CronsService) RunFetchSteam() {
	// c := cron.New()
	// c.AddFunc("* * * * *", cs.steamService.FetchSteamData)
	// c.Start()

	go cs.steamService.FetchAndSaveSteamData()
}
