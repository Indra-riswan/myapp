package entity

type BestProduk struct {
	ID      uint64 `gorm:"primaryKey; autoIncrement" json:"id"`
	Nama    string `gorm:"type:char(100)" json:"nama"`
	Harga   uint   `json:"harga"`
	IsReady bool   `gorm:"type:bool;default:true" json:"is_ready"`
	Gambar  string `gorm:"type:varchar(200);unique" json:"gambar"`
}
