package service

import (
	"fmt"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	UserUpdate(userid uint, user dto.UserDtoUpdt) entity.User
	ProfilUser(ID uint) entity.User
	UserKeranjangs(ID uint) entity.User
}

type userservice struct {
	repository repository.UserRepo
}

func NewUserService(repository repository.UserRepo) *userservice {
	return &userservice{repository}
}

func (s *userservice) UserUpdate(userid uint, user dto.UserDtoUpdt) entity.User {
	users := s.repository.ProfilUser(userid)
	err := smapping.FillStruct(&users, smapping.MapFields(&user))
	if err != nil {
		fmt.Println("Failed Create Update User :", err.Error())
	}
	s.repository.UpdateUser(users)
	return users
}

func (s *userservice) ProfilUser(ID uint) entity.User {
	return s.repository.ProfilUser(ID)
}

func (s *userservice) UserKeranjangs(ID uint) entity.User {
	return s.repository.Keranjangs(ID)
}
