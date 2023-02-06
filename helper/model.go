package helper

import (
	"belajar-go-restful-api-latihan1/model/domain"
	"belajar-go-restful-api-latihan1/model/web"
)

func ToBukuResponse(Buku domain.Buku) web.ResponseBuku {
	BukuResponse := web.ResponseBuku{
		Id:       Buku.Id,
		Judul:    Buku.Judul,
		Penulis:  Buku.Penulis,
		Penerbit: Buku.Penerbit,
		Tahun:    Buku.Tahun,
	}
	return BukuResponse
}

func ToBukuResponses(books []domain.Buku) []web.ResponseBuku {
	var responseBooks []web.ResponseBuku
	for _, buku := range books {
		// data := web.ResponseBuku{
		// 	Id:       buku.Id,
		// 	Judul:    buku.Judul,
		// 	Penulis:  buku.Penulis,
		// 	Penerbit: buku.Penerbit,
		// 	Tahun:    buku.Tahun,
		// }
		//responseBooks = append(responseBooks, data)
		responseBooks = append(responseBooks, ToBukuResponse(buku))
	}
	return responseBooks

}
