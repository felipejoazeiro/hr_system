package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotnev"
	"system/pkg/db"
	"system/routes"
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

	router.Get("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Printf("API rodando em http://localhost:%s", port)

	router.Run(":" + port)


}