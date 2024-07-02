package services

import (
	"fmt"
	"gamepricehistory/internal/infrastructure"
	"gamepricehistory/internal/repository"
	"gamepricehistory/internal/utils"
	"strconv"
	"time"
)

type ISteamService interface {
	FetchSteamData()
}

type SteamService struct {
	steamRepository *repository.SteamRepository
	steamApi        *infrastructure.SteamApi
}

type SteamGameData struct {
}

func NewSteamService(sr *repository.SteamRepository, sapi *infrastructure.SteamApi) *SteamService {
	return &SteamService{steamRepository: sr, steamApi: sapi}
}

func (ss *SteamService) FetchAndSaveSteamData() error {

	apps, err := ss.steamApi.FetchSteamApiAllGames()

	if err != nil {
		return fmt.Errorf("failed to request steam game appid: %s", err)
	}

	applst := apps.Applist.Apps

	count := 0

	for i := 0; i < len(applst); i++ {
		fmt.Println(i)
		appid := strconv.Itoa(applst[i].Appid)

		var app *infrastructure.App

		utils.Retry(func() error {
			app, err = ss.steamApi.FetchSteamGameInfo(appid)
			return err
		}, 10, 10)

		if err != nil {
			return fmt.Errorf("failed to request steam game info: %s", err)
		}

		if !app.Success {
			continue
		}

		count++

		gd := repository.GameData{AppId: applst[i].Appid, Name: app.Data.Name, Price: app.Data.PriceOverview.Final, CoverUrl: app.Data.HeaderImage, Vendor: repository.STEAM, CreatedAt: time.Now()}

		ss.steamRepository.SaveGameData(gd)

		if count >= 100 {
			time.Sleep(time.Minute * 5)
			count = 0
		}
	}

	return nil
}
