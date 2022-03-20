package dto

type KeranjangDto struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Harga     uint   `json:"harga" form:"harga" binding:"required,number"`
	Jumlah    uint   `json:"jumlah" form:"jumlah" binding:"required,number"`
	Deskripsi string `json:"deskripsi,omitempty"`
	UserID    uint   `json:"user_id,omitempty" form:"user_id,omitempty"`
}
