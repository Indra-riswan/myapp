package repository

import (
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"gorm.io/gorm"
)

type BestProdukRepo interface {
	InsertProduk(produk entity.BestProduk) entity.BestProduk
	AllBestProduk() []entity.BestProduk
	FindBestProduk(bestprodakID uint) entity.BestProduk
}

type bestprodukrepo struct {
	db *gorm.DB
}

func NewBestProdukRepo(dbi *gorm.DB) *bestprodukrepo {
	return &bestprodukrepo{dbi}
}

func (r *bestprodukrepo) InsertProduk(produk entity.BestProduk) entity.BestProduk {
	r.db.Save(&produk)
	return produk
}

func (r *bestprodukrepo) AllBestProduk() []entity.BestProduk {
	var bestproduk []entity.BestProduk
	r.db.Find(&bestproduk)
	return bestproduk
}

func (r *bestprodukrepo) FindBestProduk(bestprodakID uint) entity.BestProduk {
	var bestproduk entity.BestProduk
	r.db.Find(&bestproduk, bestprodakID)
	return bestproduk
}
