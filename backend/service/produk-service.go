package service

import (
	"fmt"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/repository"
	"github.com/mashingan/smapping"
)

type ProdukService interface {
	CreateProduk(produk dto.ProdukDto) entity.Produk
	AllProduk() []entity.Produk
	FindProduk(ID uint) entity.Produk
	UpdateProduk(produkid uint, produk dto.ProdukDtoUpdt) entity.Produk
	DeleteProduk(produkid uint) entity.Produk
	QueryProduk(nama string) entity.Produk
}

type produkservice struct {
	repository repository.ProdukRepo
}

func NewProdukService(repository repository.ProdukRepo) *produkservice {
	return &produkservice{repository}
}

func (s *produkservice) CreateProduk(produk dto.ProdukDto) entity.Produk {
	var produks = entity.Produk{}
	err := smapping.FillStruct(&produks, smapping.MapFields(&produk))
	if err != nil {
		fmt.Println("Failed Create Produk :", err.Error())
	}
	s.repository.InsertProduk(produks)
	return produks
}

func (s *produkservice) AllProduk() []entity.Produk {
	return s.repository.AllProduk()
}

func (s *produkservice) FindProduk(ID uint) entity.Produk {
	return s.repository.FindProduk(ID)
}

func (s *produkservice) UpdateProduk(produkid uint, produk dto.ProdukDtoUpdt) entity.Produk {
	produks := s.repository.FindProduk(produkid)
	err := smapping.FillStruct(&produks, smapping.MapFields(&produk))
	if err != nil {
		fmt.Println("Failed Update Produk :", err.Error())
	}
	s.repository.InsertProduk(produks)
	return produks
}

func (s *produkservice) DeleteProduk(produkid uint) entity.Produk {
	produk := s.repository.FindProduk(produkid)
	s.repository.DeleteProduk(produk)
	return produk
}

func (s *produkservice) QueryProduk(nama string) entity.Produk {
	return s.repository.QueryProduk(nama)
}
