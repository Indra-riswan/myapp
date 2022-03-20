package handler

import (
	"net/http"
	"strconv"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/helper"
	"github.com/Indra-riswan/vue-golang-backend2/service"
	"github.com/gin-gonic/gin"
)

type authheandler struct {
	authservice service.AuthService
	jwt         service.JwtService
}

func NewAuthHeandler(auth service.AuthService, jwt service.JwtService) *authheandler {
	return &authheandler{auth, jwt}
}

func (h *authheandler) Login(ctx *gin.Context) {
	var userdto dto.UserDtoLogin
	err := ctx.ShouldBind(&userdto)
	if err != nil {
		res := helper.BuildErrors("Failed Send Request", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authresult := h.authservice.VerifyCredential(userdto.Email, userdto.Password)
	if v, ok := authresult.(entity.User); ok {
		generate := h.jwt.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generate
		res := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(200, res)
		return
	}
	res := helper.BuildErrors("Denied", "Please Check Again Your Email or Password", helper.Emptyobj{})
	ctx.AbortWithStatusJSON(http.StatusConflict, res)

}

func (h *authheandler) Register(ctx *gin.Context) {
	var userdtoupdt dto.UserDto
	err := ctx.ShouldBind(&userdtoupdt)
	if err != nil {
		res := helper.BuildErrors("Failed Send Request", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if !h.authservice.IsDuplicate(userdtoupdt.Email) {
		res := helper.BuildErrors("Failed Send Request", "User Already Exist", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	} else {
		usercreate := h.authservice.CreateUser(userdtoupdt)
		res := helper.BuildResponse(true, "Ok!", usercreate)
		ctx.JSON(200, res)
	}

}
