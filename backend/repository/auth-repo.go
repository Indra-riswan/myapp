package repository

import (
	"fmt"

	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo interface {
	VerifyCredential(email, password string) interface{}
	IsDuplicate(email string) (tx *gorm.DB)
	InsertUser(user entity.User) entity.User
}

type authrepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authrepo {
	return &authrepo{db}
}

func (r *authrepo) VerifyCredential(email, password string) interface{} {
	var user entity.User
	res := r.db.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (r *authrepo) IsDuplicate(email string) (tx *gorm.DB) {
	var user entity.User
	return r.db.Where("email = ?", email).Take(&user)

}

func HashAndSalt(passwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(passwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(hash)
}

func (r *authrepo) InsertUser(user entity.User) entity.User {
	user.Password = HashAndSalt([]byte(user.Password))
	r.db.Save(&user)
	return user
}
