package repository

import (
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"gorm.io/gorm"
)

type UserRepo interface {
	UpdateUser(user entity.User) entity.User
	ProfilUser(ID uint) entity.User
	Keranjangs(ID uint) entity.User
}

type userrepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userrepo {
	return &userrepo{db}
}

func (r *userrepo) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = HashAndSalt([]byte(user.Password))
	} else {
		var tempuser entity.User
		r.db.Find(tempuser, user.ID)
		user.Password = tempuser.Password
	}
	r.db.Save(&user)
	return user
}

func (r *userrepo) ProfilUser(ID uint) entity.User {
	var user entity.User
	r.db.Find(&user, ID)
	return user
}

func (r *userrepo) Keranjangs(ID uint) entity.User {
	var user entity.User
	r.db.Preload("Keranjangs").Find(&user, ID)
	return user
}
