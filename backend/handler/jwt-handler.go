package handler

import (
	"log"
	"net/http"

	"github.com/Indra-riswan/vue-golang-backend2/helper"
	"github.com/Indra-riswan/vue-golang-backend2/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizedJwt(jwtService service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			respons := helper.BuildErrors("Failed to procces request ", "No token found,", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, respons)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[users] :", claims["users"])
			log.Println("Claims[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			respons := helper.BuildErrors("erorr wrong token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
		}
	}
}
