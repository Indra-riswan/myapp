package repository

import (
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"gorm.io/gorm"
)

type ProdukRepo interface {
	InsertProduk(produk entity.Produk) entity.Produk
	AllProduk() []entity.Produk
	UpdateProduk(produk entity.Produk) entity.Produk
	DeleteProduk(produk entity.Produk) entity.Produk
	FindProduk(ID uint) entity.Produk
	QueryProduk(nama string) entity.Produk
}

type produkrepo struct {
	db *gorm.DB
}

func NewProdukRepo(db *gorm.DB) *produkrepo {
	return &produkrepo{db}
}

func (r *produkrepo) InsertProduk(produk entity.Produk) entity.Produk {
	r.db.Save(&produk)
	return produk
}

func (r *produkrepo) AllProduk() []entity.Produk {
	var produks []entity.Produk
	r.db.Find(&produks)
	return produks
}

func (r *produkrepo) UpdateProduk(produk entity.Produk) entity.Produk {
	r.db.Save(&produk)
	return produk
}

func (r *produkrepo) DeleteProduk(produk entity.Produk) entity.Produk {
	r.db.Delete(&produk)
	return produk
}

func (r *produkrepo) FindProduk(ID uint) entity.Produk {
	var produk entity.Produk
	r.db.Find(&produk, ID)
	return produk
}

func (r *produkrepo) QueryProduk(nama string) entity.Produk {
	var produk entity.Produk
	r.db.Where("nama = ?", nama).Find(&produk)
	return produk

}
