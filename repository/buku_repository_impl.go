package repository

import (
	"belajar-go-restful-api-latihan1/helper"
	"belajar-go-restful-api-latihan1/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BukuRepositoryImpl struct {
}

func NewBukuRepository() BukuRepository {
	return &BukuRepositoryImpl{}
}

func (repository *BukuRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku {
	sql := "Insert into buku(judul,penulis,penerbit,tahun) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)
	helper.PanicHandler(err)
	id, err := result.LastInsertId()
	helper.PanicHandler(err)
	buku.Id = int(id)
	return buku
}

func (repository *BukuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, buku domain.Buku) domain.Buku {
	sql := "UPDATE buku set judul=?,penulis=?,penerbit=?,tahun=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sql, buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Id)
	helper.PanicHandler(err)

	return buku
}

func (repository *BukuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, BukuId int) {
	sql := "DELETE FROM buku WHERE id=?"
	_, err := tx.ExecContext(ctx, sql, BukuId)
	helper.PanicHandler(err)

}

func (repository *BukuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, BukuId int) (domain.Buku, error) {
	sql := "SELECT id,judul,penulis,penerbit,tahun FROM buku WHERE id=?"
	rows, err := tx.QueryContext(ctx, sql, BukuId)
	helper.PanicHandler(err)
	defer rows.Close()

	buku := domain.Buku{}
	if rows.Next() {
		//jika ada
		err := rows.Scan(&buku.Id, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Tahun)
		helper.PanicHandler(err)
		return buku, nil
	} else {
		return buku, errors.New("BUKU IS NOT FOUND")
	}
}

func (repository *BukuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Buku {
	sql := "SELECT id, judul, penulis,penerbit,tahun FROM buku"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicHandler(err)
	defer rows.Close()
	books := []domain.Buku{}
	for rows.Next() {
		var buku domain.Buku
		err := rows.Scan(&buku.Id, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Tahun)
		helper.PanicHandler(err)
		books = append(books, buku)
	}
	return books

}
