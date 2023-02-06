package repository

import (
	"belajar-go-restful-api-latihan1/model/domain"
	"context"
	"database/sql"
)

type BukuRepository interface {
	Save(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku
	Update(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku
	Delete(ctx context.Context, tx *sql.Tx, BukuId int)
	FindById(ctx context.Context, tx *sql.Tx, BukuId int) (domain.Buku, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Buku
}
