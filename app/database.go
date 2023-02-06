package app

import (
	"belajar-go-restful-api-latihan1/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api_latihan")
	helper.PanicHandler(err)
	//manimal jumlah koneksi stanbay 10
	db.SetMaxIdleConns(10)
	//maksimal koneksi yg diijinkan 20
	db.SetMaxOpenConns(20)
	//apabila koneksi yg digunkan lebih dari minimum dan
	//tidak digunakan lagi maka setelah 10 menit akan di set kembali koneksinya
	// menjadi jumlah minimal
	db.SetConnMaxIdleTime(10 * time.Minute)
	//refresh koneksi setiap 1 jam sekali
	//dan menset ke jumlah minimal koneksi
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

}
