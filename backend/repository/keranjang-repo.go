package repository

import (
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"gorm.io/gorm"
)

type KeranjangRepo interface {
	InsertKeranjang(keranjang entity.Keranjang) entity.Keranjang
	AllKeranjang() []entity.Keranjang
	DeleteKeranjang(keranjang entity.Keranjang)
	FindKeranjang(keranjangID uint) entity.Keranjang
}

type keranjangrepo struct {
	db *gorm.DB
}

func NewKeranjangRepo(db *gorm.DB) *keranjangrepo {
	return &keranjangrepo{db}
}

func (r *keranjangrepo) InsertKeranjang(keranjang entity.Keranjang) entity.Keranjang {
	r.db.Save(&keranjang)
	return keranjang
}

func (r *keranjangrepo) AllKeranjang() []entity.Keranjang {
	var keranjang []entity.Keranjang
	r.db.Find(&keranjang)
	return keranjang
}

func (r *keranjangrepo) DeleteKeranjang(keranjang entity.Keranjang) {
	r.db.Delete(&keranjang)
}

func (r *keranjangrepo) FindKeranjang(keranjangID uint) entity.Keranjang {
	var keranjang entity.Keranjang
	r.db.Find(&keranjang, keranjangID)
	return keranjang
}
