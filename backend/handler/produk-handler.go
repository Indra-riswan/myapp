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

type produkhandler struct {
	service service.ProdukService
}

func NewProdukHandler(service service.ProdukService) *produkhandler {
	return &produkhandler{service}
}

func (h *produkhandler) Create(ctx *gin.Context) {
	var produkdto dto.ProdukDto
	err := ctx.ShouldBindJSON(&produkdto)
	if err != nil {
		res := helper.BuildErrors("Failed Create Produk :", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	result := h.service.CreateProduk(produkdto)
	res := helper.BuildResponse(true, "Ok!", result)
	ctx.JSON(200, res)
}

func (h *produkhandler) AllProduk(ctx *gin.Context) {
	var produks []entity.Produk = h.service.AllProduk()
	res := helper.BuildResponse(true, "Ok!", produks)
	ctx.JSON(200, res)
}

func (h *produkhandler) FindProduk(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := h.service.FindProduk(uint(id))
	if (result == entity.Produk{}) {
		res := helper.BuildErrors("Failed Send Request", "No Produk Found", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Ok!", result)
		ctx.JSON(200, res)

	}
}

func (h *produkhandler) Update(ctx *gin.Context) {
	var produkupdt dto.ProdukDtoUpdt
	err := ctx.ShouldBindJSON(&produkupdt)
	if err != nil {
		res := helper.BuildErrors("Failed Update Produk", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return

	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := h.service.UpdateProduk(uint(id), produkupdt)
	res := helper.BuildResponse(true, "Ok!", result)
	ctx.JSON(200, res)

}

func (h *produkhandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := h.service.DeleteProduk(uint(id))
	res := helper.BuildResponse(true, "Succesed Delete !", result)
	ctx.JSON(http.StatusOK, res)
}

func (h *produkhandler) Query(ctx *gin.Context) {
	nama := ctx.Query("nama")
	// ctx.JSON(200, nama)
	result := h.service.QueryProduk(nama)
	if (result == entity.Produk{}) {
		res := helper.BuildErrors("Failed Send Request", "No Produk Found", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Ok!", result)
		ctx.JSON(200, res)

	}

}
