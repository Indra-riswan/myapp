package service

import (
	"fmt"
	"log"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/repository"
	"github.com/mashingan/smapping"
)

type KeranjangService interface {
	CreateKeranjang(keranjang dto.KeranjangDto) entity.Keranjang
	AllKeranjang() []entity.Keranjang
	DeleteKeranjang(ID uint)
	AllowEdit(UserID string, keranjangid uint) bool
	FindKeranjang(keranjangID uint) entity.Keranjang
}

type keranjangservice struct {
	repository repository.KeranjangRepo
}

func NewKeranjangService(repo repository.KeranjangRepo) *keranjangservice {
	return &keranjangservice{repo}
}

func (s *keranjangservice) CreateKeranjang(keranjang dto.KeranjangDto) entity.Keranjang {
	var keranjangs = entity.Keranjang{}
	err := smapping.FillStruct(&keranjangs, smapping.MapFields(&keranjang))
	if err != nil {
		log.Panic(err.Error())
	}
	s.repository.InsertKeranjang(keranjangs)
	return keranjangs
}

func (s *keranjangservice) AllKeranjang() []entity.Keranjang {
	return s.repository.AllKeranjang()
}

func (s *keranjangservice) DeleteKeranjang(keranjangID uint) {
	keranjang := s.repository.FindKeranjang(keranjangID)
	s.repository.DeleteKeranjang(keranjang)
}

func (s *keranjangservice) AllowEdit(UserID string, keranjangid uint) bool {
	keranjang := s.repository.FindKeranjang(keranjangid)
	id := fmt.Sprintf("%v", keranjang.UserID)
	return UserID == id

}

func (s *keranjangservice) FindKeranjang(keranjangID uint) entity.Keranjang {
	return s.repository.FindKeranjang(keranjangID)
}
