package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/helper"
	"github.com/Indra-riswan/vue-golang-backend2/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userhandler struct {
	service service.UserService
	jwt     service.JwtService
	auth    service.AuthService
}

func NewUserHandler(service service.UserService, jwt service.JwtService, auth service.AuthService) *userhandler {
	return &userhandler{service, jwt, auth}
}

func (h *userhandler) Update(ctx *gin.Context) {
	var userdtoupdt dto.UserDtoUpdt
	err := ctx.ShouldBindJSON(&userdtoupdt)
	if err != nil {
		res := helper.BuildErrors("Failed Create User", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if !h.auth.IsDuplicate(userdtoupdt.Email) {
		res := helper.BuildErrors("Failed Send Request", "User Already Exist", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}
	autheader := ctx.GetHeader("Authorization")
	token, err := h.jwt.ValidateToken(autheader)
	if err != nil {
		res := helper.BuildErrors("Failed Update User ", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["users"]))
	userdtoupdt.ID = uint(id)
	result := h.service.UserUpdate(userdtoupdt.ID, userdtoupdt)
	res := helper.BuildResponse(true, "Ok!", result)
	ctx.JSON(200, res)
}

func (h *userhandler) ProfilUser(ctx *gin.Context) {
	autheader := ctx.GetHeader("Authorization")
	token, err := h.jwt.ValidateToken(autheader)
	if err != nil {
		res := helper.BuildErrors("Erorr Access Profil User", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["users"]))
	user := h.service.ProfilUser(uint(id))
	res := helper.BuildResponse(true, "Ok!", user)
	ctx.JSON(200, res)
}

func (h *userhandler) UserKeranjangs(ctx *gin.Context) {
	autheader := ctx.GetHeader("Authorization")
	token, err := h.jwt.ValidateToken(autheader)
	if err != nil {
		res := helper.BuildErrors("Erorr Access User Keranjang", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["users"]))
	user := h.service.UserKeranjangs(uint(id))
	res := helper.BuildResponse(true, "Ok!", user)
	ctx.JSON(200, res)
}
