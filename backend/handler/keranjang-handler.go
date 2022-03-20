package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Indra-riswan/vue-golang-backend2/dto"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/helper"
	"github.com/Indra-riswan/vue-golang-backend2/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type keranjanghandler struct {
	service service.KeranjangService
	jwt     service.JwtService
}

func NewKeranjangHandler(service service.KeranjangService, jwt service.JwtService) *keranjanghandler {
	return &keranjanghandler{service, jwt}
}

func (h *keranjanghandler) Create(ctx *gin.Context) {
	var keranjangDto dto.KeranjangDto
	err := ctx.ShouldBindJSON(&keranjangDto)
	if err != nil {
		res := helper.BuildErrors("Failed Insert to Keranjang", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	getheader := ctx.GetHeader("Authorization")
	token, err := h.jwt.ValidateToken(getheader)
	if err != nil {
		helper.BuildErrors("No Token Found :", err.Error(), helper.Emptyobj{})
	} else {
		claims := token.Claims.(jwt.MapClaims)
		id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["users"]))
		keranjangDto.UserID = uint(id)
		result := h.service.CreateKeranjang(keranjangDto)
		res := helper.BuildResponse(true, "Insert To Keranjang", result)
		ctx.JSON(http.StatusOK, res)
	}
}

func (h *keranjanghandler) Delete(ctx *gin.Context) {
	var keranjang entity.Keranjang
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.BuildErrors("Failed send Request", "No Data Found", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}
	keranjang.ID = uint(id)
	getheader := ctx.GetHeader("Authorization")
	t, err := h.jwt.ValidateToken(getheader)
	if err != nil {
		res := helper.BuildErrors("Invalid Token", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	claims := t.Claims.(jwt.MapClaims)
	userid := fmt.Sprintf("%v", claims["users"])
	if h.service.AllowEdit(userid, keranjang.ID) {
		h.service.DeleteKeranjang(uint(id))
		res := helper.BuildResponse(true, "Succesed Delete Keranjang", helper.Emptyobj{})
		ctx.JSON(200, res)
	} else {
		res := helper.BuildErrors("Failed Send Request", "Forbidden", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, res)
	}

}

func (h *keranjanghandler) FindKeranjang(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	keranjang := h.service.FindKeranjang(uint(id))
	if (keranjang == entity.Keranjang{}) {
		res := helper.BuildErrors("Failed Send Request", "No Data Found", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		res := helper.BuildResponse(true, "Ok!", keranjang)
		ctx.JSON(http.StatusFound, res)
	}
}

func (h *keranjanghandler) AllKeranjang(ctx *gin.Context) {
	var keranjangs []entity.Keranjang = h.service.AllKeranjang()
	res := helper.BuildResponse(true, "Data Keranjang", keranjangs)
	ctx.JSON(200, res)
}
