package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
	"strconv"

	"github.com/emerli/go-projects/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound         = errors.New("user-id non valido o mancante")
	InvalidListFormatError  = errors.New("il campo deve essere una lista di  numero intero")
	InvalidFormatError      = errors.New("il campo deve essere un numero intero")
	FilterNotSupportedError = errors.New("filtro non supportato")
	TooMuchRecordsError     = errors.New("troppi record trovati")
	IdIsMandatoryError      = errors.New("id must be a number")
	ResponseNilError        = errors.New("HTTP response  must be provided")
	BadRequest              = errors.New("bad request")
)

func CreateCity(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var do models.City
		err := ReadRequestBody(r, &do)
		if err != nil {
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusBadRequest)
			return
		}
		var dao models.City
		id, err := dao.Insert(&do, db)
		if err != nil {
			fmt.Println(fmt.Sprintf("error: %s", err.Error()))
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusInternalServerError)
		} else {
			fmt.Println("Sending response")
			response := models.NewBaseResponse(id)
			response.WriteTo(w, http.StatusOK)
		}
	}
}
func GetCity(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response := models.NewBaseResponseError(IdIsMandatoryError)
			response.WriteTo(w, http.StatusBadRequest)
			return
		}
		var dao models.City

		city, err := dao.Get(id, db)
		if err != nil {
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusInternalServerError)
			return
		}

		response := models.NewBaseResponse(city)
		response.WriteTo(w, http.StatusOK)
	}
}

func UpdateCity(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response := models.NewBaseResponseError(IdIsMandatoryError)
			response.WriteTo(w, http.StatusBadRequest)
			return
		}
		var do models.City

		err = ReadRequestBody(r, &do)
		if err != nil {
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusBadRequest)
			return
		}

		var dao models.City

		err = dao.Update(&do, db)

		if err != nil {
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusInternalServerError)
		} else {
			response := models.NewBaseResponse(id)
			response.WriteTo(w, http.StatusOK)
		}
	}
}

func DeleteCity(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			response := models.NewBaseResponseError(IdIsMandatoryError)
			response.WriteTo(w, http.StatusBadRequest)
			return
		}
		var dao models.City

		err = dao.Delete(id, db)

		if err != nil {
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusInternalServerError)
		} else {
			response := models.NewBaseResponseWOData()
			response.WriteTo(w, http.StatusOK)
		}
	}
}

func GetAllCity(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dao models.City
		cities, err := dao.GetAll(db)

		if err != nil {
			response := models.NewBaseResponseError(err)
			response.WriteTo(w, http.StatusInternalServerError)
		} else {
			response := models.NewBaseResponse(cities)
			response.WriteTo(w, http.StatusOK)
		}
	}
}

func ReadRequestBody(r *http.Request, v interface{}) error {
	if r == nil || r.Body == nil {
		return ResponseNilError
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(body), &v); err != nil {
		return err
	}
	return nil
}
