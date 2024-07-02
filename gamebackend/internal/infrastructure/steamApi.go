package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const apiKey string = "DCAA5D99E3509DBA7ECB2B750BD53AD7"

var RequestLimitReached = errors.New("RequestLimitReached")

type SteamApi struct {
	client *http.Client
}

// This type implements the http.RoundTripper interface
type LoggingRoundTripper struct {
	Proxied  http.RoundTripper
	Delay    int
	MaxRetry int
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	// Do "before sending requests" actions here.
	fmt.Printf("Sending request to %v\n", req.URL)

	// Send the request, get the response (or the error)
	res, err := lrt.Proxied.RoundTrip(req)

	fmt.Printf("Received %v response\n", res.Status)

	return res, err
}

func NewSteamApi() *SteamApi {
	client := &http.Client{Timeout: 60 * time.Second, Transport: LoggingRoundTripper{Proxied: http.DefaultTransport, Delay: 60, MaxRetry: 10}}
	return &SteamApi{client: client}
}

type Applst struct {
	Applist struct {
		Apps []struct {
			Appid int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

type App struct {
	Data struct {
		Name          string `json:"name"`
		PriceOverview struct {
			Currency        string  `json:"currency"`
			Initial         int     `json:"initial"`
			Final           float64 `json:"final"`
			DiscountPercent string  `json:"discount_percent"`
		} `json:"price_overview"`
		HeaderImage string `json:"header_image"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (sa *SteamApi) FetchSteamApiAllGames() (*Applst, error) {
	apps := &Applst{}
	url := fmt.Sprintf("http://api.steampowered.com/ISteamApps/GetAppList/v0002/?key=%s&format=json", apiKey)
	resp, err := sa.client.Get(url)

	if err != nil {
		return nil, err
	}

	json.NewDecoder(resp.Body).Decode(apps)

	return apps, nil
}

func (sa *SteamApi) FetchSteamGameInfo(appid string) (*App, error) {
	app := &App{}
	r, _ := http.NewRequest("GET", "http://store.steampowered.com/api/appdetails", nil)
	q := r.URL.Query()
	q.Add("appids", appid)

	res, _ := sa.client.Do(r)

	if res.StatusCode == 429 {
		return nil, RequestLimitReached
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("falha ao buscar dados na Steam")
	}

	var s map[string]any

	json.NewDecoder(res.Body).Decode(&s)

	rm, _ := json.Marshal(s[appid])

	json.Unmarshal(rm, &app)

	return app, nil
}
