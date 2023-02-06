package controller

import (
	"belajar-go-restful-api-latihan1/helper"
	"belajar-go-restful-api-latihan1/model/web"
	"belajar-go-restful-api-latihan1/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BukuControllerImpl struct {
	Service services.BukuServices
}

func NewBukuControllerImpl(service services.BukuServices) BukuController {
	return &BukuControllerImpl{
		Service: service,
	}
}

func (control *BukuControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	createWebRequest := web.CreateWebRequest{}
	err := decoder.Decode(&createWebRequest)
	helper.PanicHandler(err)
	responseBuku := control.Service.Create(request.Context(), createWebRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseBuku,
	}

	helper.WriteFromResponse(writer, webResponse)

}

func (control *BukuControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	webUpdate := web.UpdateWebRequest{}
	helper.ReadFromRequest(request, &webUpdate)

	BukuId := params.ByName("BukuId")
	id, err := strconv.Atoi(BukuId)
	helper.PanicHandler(err)
	webUpdate.Id = id
	responseBuku := control.Service.Update(request.Context(), webUpdate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseBuku,
	}
	helper.WriteFromResponse(writer, &webResponse)

}

func (control *BukuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Bukuid := params.ByName("BukuId")
	id, err := strconv.Atoi(Bukuid)

	helper.PanicHandler(err)
	control.Service.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteFromResponse(writer, &webResponse)
}

func (control *BukuControllerImpl) FindByid(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	BukuId := params.ByName("BukuId")
	id, err := strconv.Atoi(BukuId)
	helper.PanicHandler(err)
	bukuResponse := control.Service.FindByid(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bukuResponse,
	}
	helper.WriteFromResponse(writer, &webResponse)
}

func (control *BukuControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bukuResponse := control.Service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bukuResponse,
	}

	//helper.WriteFromResponse(writer, &webResponse)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicHandler(err)

}
