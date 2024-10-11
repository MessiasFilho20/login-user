package main

import (
	"log"
	"login-user/prisma/db"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("Uagw|2rWcrvt,KO8£$QWga;5[9LC58Gz8kzH%DH&£q]5iaE2l")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var client *db.PrismaClient

func main() {
	client = db.NewClient()
	err := client.Prisma.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer client.Prisma.Disconnect()

	r := gin.Default()

	r.POST("/regiister")

	r.GET("teste", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello word")
		ctx.JSON(200, gin.H{"messege": "teste okay"})
	})

	r.Run("localhost:8000")
}
