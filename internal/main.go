package main

import (
	"log"

	"github.com/AkmalArifin/short-url/internal/db"
	"github.com/AkmalArifin/short-url/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Panic("could not load env file")
	}

	db.InitDB()
	r := gin.Default()
	routes.ServeRouter(r)
	r.Run(":8080")
}
