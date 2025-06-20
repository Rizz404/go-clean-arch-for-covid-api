package main

import (
	"github.com/Rizz404/go-clean-arch-for-covid-api/internal/repository/sqlc"
	"github.com/oklog/ulid/v2"
)

type Covid struct {
	ID        string `json:"id"`
	Nama      string `json:"nama"`
	Kota      string `json:"kota"`
	Sembuh    int32  `json:"sembuh"`
	Dirawat   int32  `json:"dirawat"`
	Meninggal int32  `json:"meninggal"`
	Total     int32  `json:"total"`
}

func dbCovidToCovid(dbCovid sqlc.Covid) Covid {
	return Covid{
		ID:        ulid.ULID(dbCovid.ID).String(),
		Nama:      dbCovid.Nama,
		Kota:      dbCovid.Kota,
		Sembuh:    dbCovid.Sembuh,
		Dirawat:   dbCovid.Dirawat,
		Meninggal: dbCovid.Meninggal,
		Total:     dbCovid.Total,
	}
}

func dbCovidsToCovids(dbCovids []sqlc.Covid) []Covid {
	covids := []Covid{}
	for _, dbCovid := range dbCovids {
		covids = append(covids, dbCovidToCovid(dbCovid))
	}
	return covids
}
