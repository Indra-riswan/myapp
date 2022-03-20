package entity

type User struct {
	ID         uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama       string      `gorm:"type: varchar(255)" json:"nama"`
	Email      string      `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password   string      `gorm:"->;<-;not null" json:"-"`
	Token      string      `gorm:"-" json:"token,omitempty"`
	Keranjangs []Keranjang `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"keranjangs"`
}
