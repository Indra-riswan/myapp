package service

import (
	"fmt"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	IsDuplicate(email string) bool
	CreateUser(user dto.UserDto) entity.User
}

type authservice struct {
	repository repository.AuthRepo
}

func NewAuthService(repository repository.AuthRepo) *authservice {
	return &authservice{repository}
}

func ComparePassword(pass, plainpass []byte) bool {
	passwd := []byte(pass)
	err := bcrypt.CompareHashAndPassword(passwd, plainpass)
	if err != nil {
		fmt.Println("Failed Compare :", err.Error())
		return false
	}
	return true
}

func (s *authservice) VerifyCredential(email, password string) interface{} {
	res := s.repository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		compare := ComparePassword([]byte(v.Password), []byte(password))
		if v.Email == email && compare {
			return res
		}
		return false
	}
	return false
}

func (s *authservice) IsDuplicate(email string) bool {
	res := s.repository.IsDuplicate(email)
	return res.Error != nil
}

func (s *authservice) CreateUser(user dto.UserDto) entity.User {
	var userdto = entity.User{}
	err := smapping.FillStruct(&userdto, smapping.MapFields(&user))
	if err != nil {
		fmt.Println("Failed Create New User :", err.Error())
	}
	s.repository.InsertUser(userdto)
	return userdto
}
