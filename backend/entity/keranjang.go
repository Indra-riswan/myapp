package entity

type Keranjang struct {
	ID        uint   `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama      string `gorm:"type:char(100)" json:"nama"`
	Harga     uint   `gorm:"type:int" json:"harga"`
	Jumlah    uint   `gorm:"type:int" json:"jumlah"`
	Deskripsi string `gorm:"varchar(200)" json:"deskripsi"`
	UserID    uint   `gorm:"not null" json:"-"`
}
