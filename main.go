package main

import (
	"belajar-go-restful-api-latihan1/app"
	"belajar-go-restful-api-latihan1/controller"
	"belajar-go-restful-api-latihan1/helper"
	"belajar-go-restful-api-latihan1/middleware"
	"belajar-go-restful-api-latihan1/repository"
	"belajar-go-restful-api-latihan1/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	bukurepository := repository.NewBukuRepository()
	Bukuservices := services.NewBukuServieImpl(db, validate, bukurepository)
	bukucontroller := controller.NewBukuControllerImpl(Bukuservices)

	// router := httprouter.New()
	// router.GET("/api/buku", controller.FindAll)
	// router.GET("/api/buku/:BukuId", controller.FindByid)
	// router.POST("/api/buku", controller.Create)
	// router.PUT("/api/buku/", controller.Update)
	// router.DELETE("/api/buku/:BukuId", controller.Delete)
	// router.PanicHandler = exception.Errorhandler
	router := app.NewRouter(bukucontroller)
	server := http.Server{
		Addr:    "localhost:4000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicHandler(err)

}
