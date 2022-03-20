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

type bestprodukhandler struct {
	service service.BestProdukService
}

func NewBestProdukHandler(service service.BestProdukService) *bestprodukhandler {
	return &bestprodukhandler{service}
}

func (h *bestprodukhandler) Create(ctx *gin.Context) {
	var bestProdukDto dto.BestProdukDto
	err := ctx.ShouldBindJSON(&bestProdukDto)
	if err != nil {
		res := helper.BuildErrors("Failed Insert To Database", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		result := h.service.CreateBestProduk(bestProdukDto)
		res := helper.BuildResponse(true, "Ok!", result)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (h *bestprodukhandler) FindProduk(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := h.service.FindBestProduk(uint(id))
	if (result == entity.BestProduk{}) {
		res := helper.BuildErrors("Invalid", "No Data Found", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Data Produk", result)
		ctx.JSON(200, res)
	}
}

func (h *bestprodukhandler) AllBestProduk(ctx *gin.Context) {
	var bestproduk []entity.BestProduk = h.service.AllBestproduk()
	result := helper.BuildResponse(true, "Data Semua Best Produk", bestproduk)
	ctx.JSON(200, result)
}
