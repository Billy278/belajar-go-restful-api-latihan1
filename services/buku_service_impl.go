package services

import (
	"belajar-go-restful-api-latihan1/exception"
	"belajar-go-restful-api-latihan1/helper"
	"belajar-go-restful-api-latihan1/model/domain"
	"belajar-go-restful-api-latihan1/model/web"
	"belajar-go-restful-api-latihan1/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type BukuServiceImpl struct {
	DB             *sql.DB
	Validate       *validator.Validate
	BukuRepository repository.BukuRepository
}

func NewBukuServieImpl(db *sql.DB, validate *validator.Validate, bukuRepository repository.BukuRepository) BukuServices {
	return &BukuServiceImpl{
		DB:             db,
		Validate:       validate,
		BukuRepository: bukuRepository,
	}
}

func (service *BukuServiceImpl) Create(ctx context.Context, request web.CreateWebRequest) web.ResponseBuku {
	//sebelum kita buat transaksinya kita buat Validatevalidasinya
	err := service.Validate.Struct(request)
	helper.PanicHandler(err)

	tx, err := service.DB.Begin()
	helper.PanicHandler(err)

	defer helper.CommitOrRollback(tx)
	Buku := domain.Buku{
		Judul:    request.Judul,
		Penulis:  request.Penulis,
		Penerbit: request.Penerbit,
		Tahun:    request.Tahun,
	}
	Buku = service.BukuRepository.Save(ctx, tx, Buku)

	return helper.ToBukuResponse(Buku)

}

func (service *BukuServiceImpl) Update(ctx context.Context, request web.UpdateWebRequest) web.ResponseBuku {
	//validasi dulu data request nya
	err := service.Validate.Struct(request)
	helper.PanicHandler(err)
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	Buku, err := service.BukuRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	Buku.Judul = request.Judul
	Buku.Penerbit = request.Penerbit
	Buku.Penulis = request.Penulis
	Buku.Tahun = request.Tahun
	Buku = service.BukuRepository.Update(ctx, tx, Buku)
	return helper.ToBukuResponse(Buku)
}

func (service *BukuServiceImpl) Delete(ctx context.Context, bukuId int) {
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)
	buku, err := service.BukuRepository.FindById(ctx, tx, bukuId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BukuRepository.Delete(ctx, tx, buku.Id)
}

func (service *BukuServiceImpl) FindByid(ctx context.Context, bukuId int) web.ResponseBuku {
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	Buku, err := service.BukuRepository.FindById(ctx, tx, bukuId)
	helper.PanicHandler(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBukuResponse(Buku)
}

func (service *BukuServiceImpl) FindAll(ctx context.Context) []web.ResponseBuku {
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)
	Books := service.BukuRepository.FindAll(ctx, tx)
	return helper.ToBukuResponses(Books)

}
