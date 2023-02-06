package web

type CreateWebRequest struct {
	Judul    string `validate:"required" json:"judul"`
	Penulis  string `validate:"required" json:"penulis"`
	Penerbit string `validate:"required" json:"penerbit"`
	Tahun    int    `validate:"required" json:"tahun"`
}
