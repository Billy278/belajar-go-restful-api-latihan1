package exception

import (
	"belajar-go-restful-api-latihan1/helper"
	"belajar-go-restful-api-latihan1/model/web"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func Errorhandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationErrors(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}
	fmt.Println("errHandler")
	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, er interface{}) bool {
	exception, ok := er.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		helper.WriteFromResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriteFromResponse(writer, webResponse)
		return true

	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL STATUS ERROR",
		Data:   err,
	}
	helper.WriteFromResponse(writer, webResponse)
}
