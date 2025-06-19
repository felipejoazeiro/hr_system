package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"app/pkg/db"
	"app/routes"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis padrão")
	}

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	db.InitDB()

	router:=routes.SetupRouter()

	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Printf("API rodando em http://localhost:%s", port)

	router.Run(":" + port)


}