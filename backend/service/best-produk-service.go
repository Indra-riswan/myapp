package service

import (
	"fmt"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/repository"
	"github.com/mashingan/smapping"
)

type BestProdukService interface {
	CreateBestProduk(bestproduk dto.BestProdukDto) entity.BestProduk
	AllBestproduk() []entity.BestProduk
	FindBestProduk(bestprodukID uint) entity.BestProduk
}

type bestprodukservice struct {
	repo repository.BestProdukRepo
}

func NewBestProdukService(repo repository.BestProdukRepo) *bestprodukservice {
	return &bestprodukservice{repo}
}

func (s *bestprodukservice) CreateBestProduk(bestproduk dto.BestProdukDto) entity.BestProduk {
	var bestproduks = entity.BestProduk{}
	err := smapping.FillStruct(&bestproduks, smapping.MapFields(&bestproduk))
	if err != nil {
		fmt.Println("Failed Insert Produk", err.Error())
	}

	s.repo.InsertProduk(bestproduks)
	return bestproduks
}

func (s *bestprodukservice) AllBestproduk() []entity.BestProduk {
	return s.repo.AllBestProduk()
}

func (s *bestprodukservice) FindBestProduk(bestprodukID uint) entity.BestProduk {
	return s.repo.FindBestProduk(bestprodukID)
}
