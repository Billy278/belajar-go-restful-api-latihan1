package web

type UpdateWebRequest struct {
	Id       int    `validate:"required" json:"id"`
	Judul    string `validate:"required" json:"judul"`
	Penulis  string `validate:"required" json:"penulis"`
	Penerbit string `validate:"required" json:"penerbit"`
	Tahun    int    `validate:"required" json:"tahun"`
}
