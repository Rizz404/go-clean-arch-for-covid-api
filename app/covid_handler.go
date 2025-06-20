package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rizz404/go-clean-arch-for-covid-api/internal/repository/sqlc"
	"github.com/Rizz404/go-clean-arch-for-covid-api/utils"
	"github.com/go-chi/chi/v5"
)

type createRequestPayload struct {
	Nama      string `json:"nama" form:"nama"`
	Kota      string `json:"kota" form:"kota"`
	Sembuh    int64  `json:"sembuh" form:"sembuh"`
	Dirawat   int64  `json:"dirawat" form:"dirawat"`
	Meninggal int64  `json:"meninggal" form:"meninggal"`
	Total     int64  `json:"total" form:"total"`
}

type updateRequestPayload struct {
	Nama      *string `json:"nama" form:"nama"`
	Kota      *string `json:"kota" form:"kota"`
	Sembuh    *int64  `json:"sembuh" form:"sembuh"`
	Dirawat   *int64  `json:"dirawat" form:"dirawat"`
	Meninggal *int64  `json:"meninggal" form:"meninggal"`
	Total     *int64  `json:"total" form:"total"`
}

func (apiCfg *apiConfig) createCovidHandler(w http.ResponseWriter, r *http.Request) {
	var req createRequestPayload
	if err := utils.ParseRequestBody(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Gagal parse request: %v", err))
		return
	}

	if req.Nama == "" || req.Kota == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Nama dan kota harus diisi")
		return
	}

	covid, err := apiCfg.DB.CreateCovid(r.Context(), sqlc.CreateCovidParams{
		Nama:      req.Nama,
		Kota:      req.Kota,
		Sembuh:    req.Sembuh,
		Dirawat:   req.Dirawat,
		Meninggal: req.Meninggal,
		Total:     req.Total,
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Gagal menyimpan data covid: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, covid)
}

func (apiCfg *apiConfig) getCovidsHandler(w http.ResponseWriter, r *http.Request) {
	covids, err := apiCfg.DB.GetCovids(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Gagal mengambil daftar data covid: %v", err))
		return
	}

	// * Inisialisasi ke empty slice
	if covids == nil {
		covids = []sqlc.Covid{}
	}

	utils.RespondWithJSON(w, http.StatusOK, covids)
}

func (apiCfg *apiConfig) getCovidByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	covid, err := apiCfg.DB.GetCovid(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, "Data covid tidak ditemukan")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Gagal mengambil data covid: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, covid)
}

func (apiCfg *apiConfig) updateCovidHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	existingCovid, err := apiCfg.DB.GetCovid(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, "Data covid tidak ditemukan")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Gagal mengambil data covid: %v", err))
		return
	}

	var req updateRequestPayload
	if err := utils.ParseRequestBody(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Gagal parse request: %v", err))
		return
	}

	updatedData := sqlc.UpdateCovidParams{
		ID:        id,
		Nama:      existingCovid.Nama,
		Kota:      existingCovid.Kota,
		Sembuh:    existingCovid.Sembuh,
		Dirawat:   existingCovid.Dirawat,
		Meninggal: existingCovid.Meninggal,
		Total:     existingCovid.Total,
	}

	if req.Nama != nil {
		updatedData.Nama = *req.Nama
	}
	if req.Kota != nil {
		updatedData.Kota = *req.Kota
	}
	if req.Sembuh != nil {
		updatedData.Sembuh = *req.Sembuh
	}
	if req.Dirawat != nil {
		updatedData.Dirawat = *req.Dirawat
	}
	if req.Meninggal != nil {
		updatedData.Meninggal = *req.Meninggal
	}
	if req.Total != nil {
		updatedData.Total = *req.Total
	}

	covid, err := apiCfg.DB.UpdateCovid(r.Context(), updatedData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Gagal mengupdate data covid: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, covid)
}

func (apiCfg *apiConfig) deleteCovidHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	err = apiCfg.DB.DeleteCovid(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Gagal menghapus data covid: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
}
