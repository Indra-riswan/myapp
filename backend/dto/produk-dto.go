package dto

type ProdukDto struct {
	Nama   string `json:"nama" form:"nama" binding:"required"`
	Harga  uint   `json:"harga" form:"harga" binding:"required"`
	Gambar string `json:"gambar" form:"gambar" binding:"required"`
}

type ProdukDtoUpdt struct {
	Nama    string `json:"nama" form:"nama" binding:"required"`
	Harga   uint   `json:"harga" form:"harga" binding:"required"`
	IsReady bool   `json:"is_ready"`
	Gambar  string `json:"gambar" form:"gambar" binding:"required"`
}
