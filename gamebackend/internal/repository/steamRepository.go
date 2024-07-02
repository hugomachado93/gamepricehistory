package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type SteamRepository struct {
	db *sqlx.DB
}

type Vendor string

const (
	STEAM Vendor = "STEAM"
	EPIC  Vendor = "EPIC"
	GOG   Vendor = "GOG"
)

type GameData struct {
	Id              int       `db:"id"`
	AppId           int       `db:"appid"`
	Name            string    `db:"name"`
	Price           float64   `db:"price"`
	CoverUrl        string    `db:"cover_url"`
	Vendor          Vendor    `db:"vendor"`
	CreatedAt       time.Time `db:"created_at"`
	GameDataHistory []GameDataHistory
}

type GameDataHistory struct {
	Price      float64   `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	GameDataId int       `db:"game_data_id"`
}

func NewSteamRepository(db *sqlx.DB) *SteamRepository {
	return &SteamRepository{db: db}
}

func (sr *SteamRepository) SaveGameData(gdata GameData) {
	var gd GameData
	// var gdh GameDataHistory

	r := sr.db.QueryRowx("SELECT id, appid, name, price, cover_url, vendor, created_at FROM game_data gd where gd.appid = $1", gdata.AppId)

	if err := r.StructScan(&gd); err != nil {
		fmt.Println(err)
	}

	if gd.AppId == 0 {
		r, err := sr.db.NamedQuery("INSERT INTO game_data (appid, name, price, cover_url, vendor, created_at) VALUES (:appid, :name, :price, :cover_url, :vendor, :created_at) RETURNING id", gdata)
		if err != nil {
			fmt.Println(err)
		}

		if r.Next() {
			r.Scan(&gdata.Id)
		}
		_, err = sr.db.NamedExec("INSERT INTO game_data_history (price, created_at, game_data_id) VALUES (:price, :created_at, :id)", gdata)
		if err != nil {
			fmt.Println(err)
		}
		r.Close()
		return
	}

	if gdata.Price != gd.Price {
		_, err := sr.db.NamedExec("UPDATE game_data SET price = :price where id = :id", gd)
		if err != nil {
			fmt.Println(err)
		}

		_, err = sr.db.NamedExec("INSERT INTO game_data_history (price, created_at, game_data_id) VALUES (:price, :created_at, :id)", gd)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (sr *SteamRepository) GetGameDataPaginated(page int, size int) []GameData {
	r, err := sr.db.Queryx("select id, appid, name, price, cover_url, vendor, created_at from game_data gd order by name limit $1 offset $2", size, page*size)
	if err != nil {
		fmt.Println(err)
	}

	var gd GameData
	var gdh GameDataHistory

	gdlst := make([]GameData, 0)

	for r.Next() {
		r.StructScan(&gd)

		gdhlst := make([]GameDataHistory, 0)

		r, err := sr.db.Queryx("select price, created_at, game_data_id from game_data_history where game_data_id = $1", gd.Id)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		for r.Next() {
			r.StructScan(&gdh)
			gdhlst = append(gdhlst, gdh)
		}
		gd.GameDataHistory = gdhlst

		gdlst = append(gdlst, gd)

	}

	return gdlst
}
