package api

import (
	"encoding/json"
	"gamepricehistory/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Api struct {
	gameService *services.GameService
}

func CreateApis(r *mux.Router, gameService *services.GameService) {
	api := &Api{gameService: gameService}
	r.HandleFunc("/api/game/paginated", api.getGamePriceHistory)
}

func (a *Api) getGamePriceHistory(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	page, _ := strconv.Atoi(values.Get("page"))
	size, _ := strconv.Atoi(values.Get("size"))

	gamelst := a.gameService.GetGameDataPaginated(page, size)

	json.NewEncoder(w).Encode(gamelst)
}
