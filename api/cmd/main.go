package main

import (
	"github.com/gin-gonic/gin",
	"github.com/joho/godotnev",
	"log",
	"os"
)

func main(){
	err := gotdotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis padrão")
	}

	port := os.Gotenv("PORT")
	if port == ""{
		port = "8080"
	}

	router := gin.Default()

	routerGet("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Printf("API rodando em http://localhost:%s", port)

	router.Run(":" + port)
}