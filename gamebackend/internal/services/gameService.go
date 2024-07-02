package services

import "gamepricehistory/internal/repository"

type GameService struct {
	sr *repository.SteamRepository
}

func NewGameService(sr *repository.SteamRepository) *GameService {
	return &GameService{sr: sr}
}

func (gs *GameService) GetGameDataPaginated(page int, size int) []repository.GameData {
	gdatalst := gs.sr.GetGameDataPaginated(page, size)

	return gdatalst
}
