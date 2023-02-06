package services

import (
	"belajar-go-restful-api-latihan1/model/web"
	"context"
)

type BukuServices interface {
	Create(ctx context.Context, request web.CreateWebRequest) web.ResponseBuku
	Update(ctx context.Context, request web.UpdateWebRequest) web.ResponseBuku
	Delete(ctx context.Context, bukuId int)
	FindByid(ctx context.Context, bukuId int) web.ResponseBuku
	FindAll(ctx context.Context) []web.ResponseBuku
}
