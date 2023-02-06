package app

import (
	"belajar-go-restful-api-latihan1/controller"
	"belajar-go-restful-api-latihan1/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.BukuController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/buku", controller.FindAll)
	router.GET("/api/buku/:BukuId", controller.FindByid)
	router.POST("/api/buku", controller.Create)
	router.PUT("/api/buku/:BukuId", controller.Update)
	router.DELETE("/api/buku/:BukuId", controller.Delete)
	router.PanicHandler = exception.Errorhandler

	return router
}
