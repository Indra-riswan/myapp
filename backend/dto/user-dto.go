package dto

type UserDto struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserDtoLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserDtoUpdt struct {
	ID       uint   `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
