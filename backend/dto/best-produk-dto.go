package dto

type BestProdukDto struct {
	Nama   string `json:"nama" form:"nama" binding:"required"`
	Harga  uint   `json:"harga" form:"harga" binding:"required"`
	Gambar string `json:"gambar" form:"gambar" binding:"required"`
}
