package test

import (
	"belajar-go-restful-api-latihan1/app"
	"belajar-go-restful-api-latihan1/controller"
	"belajar-go-restful-api-latihan1/helper"
	"belajar-go-restful-api-latihan1/middleware"
	"belajar-go-restful-api-latihan1/model/domain"
	"belajar-go-restful-api-latihan1/repository"
	"belajar-go-restful-api-latihan1/services"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api_latihan_test")
	helper.PanicHandler(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	repository := repository.NewBukuRepository()
	services := services.NewBukuServieImpl(db, validate, repository)
	controller := controller.NewBukuControllerImpl(services)
	router := app.NewRouter(controller)
	return middleware.NewAuthMiddleware(router)
}

func turncateBuku(db *sql.DB) {
	db.Exec("TRUNCATE buku")
}

func TestCreateBukuSuccess(t *testing.T) {
	db := setupDB()
	turncateBuku(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"judul":"abc","penulis":"billy","penerbit":"adii","tahun":2021}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:4000/api/buku", requestBody)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "abc", responseBody["data"].(map[string]interface{})["judul"])

}

func TestCreateBukuFailed(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"judul":"abc","penulis":"","penerbit":"adii","tahun":2022}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:4000/api/buku", requestBody)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateBukuSuccess(t *testing.T) {
	// db := setupDB()
	// //turncateBuku(db)
	// router := setupRouter(db)
	// requestBody := strings.NewReader(`{"judul":"abc","penulis":"billy","penerbit":"adii","tahun":2021}`)
	// request := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/buku/1", requestBody)
	// request.Header.Add("Content-Type", "aplication/json")
	// request.Header.Add("KEY", "RAHASIA")
	// recorder := httptest.NewRecorder()
	// router.ServeHTTP(recorder, request)
	// response := recorder.Result()
	// assert.Equal(t, 200, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)
	// fmt.Print(responseBody)

	// assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// assert.Equal(t, "Ok", responseBody["status"])
	// assert.Equal(t, "abc", responseBody["data"].(map[string]interface{})["judul"])
	//==================================bisa juga yg atas

	db := setupDB()
	turncateBuku(db)

	bukurepository := repository.NewBukuRepository()
	tx, _ := db.Begin()
	buku := bukurepository.Save(context.Background(), tx, domain.Buku{
		Judul:    "abc",
		Penulis:  "billy",
		Penerbit: "adii",
		Tahun:    2021,
	})

	tx.Commit()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"judul":"abc","penulis":"billy","penerbit":"adii","tahun":2021}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/buku/"+strconv.Itoa(buku.Id), requestBody)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, buku.Judul, responseBody["data"].(map[string]interface{})["judul"])

}

func TestUpdateBukuFailed(t *testing.T) {
	db := setupDB()
	turncateBuku(db)

	bukurepository := repository.NewBukuRepository()
	tx, _ := db.Begin()
	buku := bukurepository.Save(context.Background(), tx, domain.Buku{
		Judul:    "abc",
		Penulis:  "billy",
		Penerbit: "adii",
		Tahun:    2021,
	})

	tx.Commit()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"judul":"","penulis":"billy","penerbit":"adii","tahun":2021}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/buku/"+strconv.Itoa(buku.Id), requestBody)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestGetBukuSuccess(t *testing.T) {
	db := setupDB()
	turncateBuku(db)

	bukurepository := repository.NewBukuRepository()
	tx, _ := db.Begin()
	buku := bukurepository.Save(context.Background(), tx, domain.Buku{
		Judul:    "abc",
		Penulis:  "billy",
		Penerbit: "adii",
		Tahun:    2021,
	})
	tx.Commit()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/buku/"+strconv.Itoa(buku.Id), nil)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	// assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// assert.Equal(t, "Ok", responseBody["status"])
	// assert.Equal(t, buku.Id, responseBody["data"].(map[string]interface{})["id"])
	// assert.Equal(t, buku.Judul, responseBody["data"].(map[string]interface{})["judul"])

}

func TestDeleteBukuSuccess(t *testing.T) {
	db := setupDB()
	turncateBuku(db)

	bukurepository := repository.NewBukuRepository()
	tx, _ := db.Begin()
	buku := bukurepository.Save(context.Background(), tx, domain.Buku{
		Judul:    "abc",
		Penulis:  "billy",
		Penerbit: "adii",
		Tahun:    2021,
	})
	tx.Commit()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:4000/api/buku/"+strconv.Itoa(buku.Id), nil)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])

}

func TestListBukuSuccess(t *testing.T) {
	db := setupDB()
	turncateBuku(db)

	bukurepository := repository.NewBukuRepository()
	tx, _ := db.Begin()
	buku1 := bukurepository.Save(context.Background(), tx, domain.Buku{
		Judul:    "aaa",
		Penulis:  "billy",
		Penerbit: "Sinarta",
		Tahun:    2021,
	})

	buku2 := bukurepository.Save(context.Background(), tx, domain.Buku{
		Judul:    "bbb",
		Penulis:  "bima",
		Penerbit: "Adita",
		Tahun:    2021,
	})

	tx.Commit()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/buku", nil)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)
	var books = responseBody["data"].([]interface{})
	book1 := books[0].(map[string]interface{})
	book2 := books[1].(map[string]interface{})
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, buku1.Id, int(book1["id"].(float64)))
	assert.Equal(t, buku1.Judul, book1["judul"])

	assert.Equal(t, buku2.Id, int(book2["id"].(float64)))
	assert.Equal(t, buku2.Judul, book2["judul"])
}

func TestUnAuthorized(t *testing.T) {
	db := setupDB()
	turncateBuku(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/buku/2", nil)
	request.Header.Add("Content-Type", "aplication/json")
	request.Header.Add("KEY", "salah")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//fmt.Print(responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])

}
