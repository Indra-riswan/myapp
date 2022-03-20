package main

import (
	"github.com/Indra-riswan/vue-golang-backend2/config"
	"github.com/Indra-riswan/vue-golang-backend2/entity"
	"github.com/Indra-riswan/vue-golang-backend2/handler"
	"github.com/Indra-riswan/vue-golang-backend2/repository"
	"github.com/Indra-riswan/vue-golang-backend2/service"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB = config.ConnectionDB()
	produkrepo                 = repository.NewProdukRepo(db)
	authrepo                   = repository.NewAuthRepo(db)
	userrepo                   = repository.NewUserRepo(db)
	keranjnagrepo              = repository.NewKeranjangRepo(db)
	bestprod                   = repository.NewBestProdukRepo(db)
	authservice                = service.NewAuthService(authrepo)
	jwtservice                 = service.NewJwtService()
	produkservice              = service.NewProdukService(produkrepo)
	userservice                = service.NewUserService(userrepo)
	keranjangservice           = service.NewKeranjangService(keranjnagrepo)
	bestprodukservice          = service.NewBestProdukService(bestprod)
	authhandler                = handler.NewAuthHeandler(authservice, jwtservice)
	produkhandler              = handler.NewProdukHandler(produkservice)
	userhandler                = handler.NewUserHandler(userservice, jwtservice, authservice)
	keranjanghandler           = handler.NewKeranjangHandler(keranjangservice, jwtservice)
	bestprodukhandler          = handler.NewBestProdukHandler(bestprodukservice)
	jwthandler                 = handler.AuthorizedJwt(jwtservice)
)

func main() {
	defer config.CloseConnectionDB(db)
	db.AutoMigrate(&entity.User{}, &entity.Produk{}, &entity.Keranjang{}, entity.BestProduk{})

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "http://localhost:8080",
		Methods:         "GET, PUT, POST, DELETE, HEAD, OPTIONS, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	auth := r.Group("auth")
	{
		auth.POST("/login", authhandler.Login)
		auth.POST("/register", authhandler.Register)
	}

	produk := r.Group("produk", jwthandler)
	{
		produk.POST("/create", produkhandler.Create)
		produk.GET("/produks", produkhandler.AllProduk)
		produk.DELETE("/:id", produkhandler.Delete)
		produk.GET("/:id", produkhandler.FindProduk)
		produk.PUT("/update/:id", produkhandler.Update)
		produk.GET("/query", produkhandler.Query)
	}

	bestproduk := r.Group("bestproduk", jwthandler)
	{
		bestproduk.GET("produks", bestprodukhandler.AllBestProduk)
		bestproduk.GET("produk/:id", bestprodukhandler.FindProduk)
		bestproduk.POST("create", bestprodukhandler.Create)
	}

	user := r.Group("user", jwthandler)
	{
		user.PUT("/update/", userhandler.Update)
		user.GET("/profil/", userhandler.ProfilUser)
		user.GET("/keranjang/", userhandler.UserKeranjangs)
	}

	keranjang := r.Group("keranjang", jwthandler)
	{
		keranjang.POST("/create", keranjanghandler.Create)
		keranjang.GET("/keranjangs", keranjanghandler.AllKeranjang)
		keranjang.DELETE("/delete/:id", keranjanghandler.Delete)
		keranjang.GET("/find/:id", keranjanghandler.FindKeranjang)
	}
	r.Run(":3000")

}
