package main

import (
	"example/apps/config"
	"example/apps/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error

	config.DB, err = config.DBConnect()

	if err != nil {
		panic("DB Connection not working")
	}

	defer config.DB.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Allow requests from your Next.js frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                                // Allow credentials (like cookies or authorization headers)
	}))

	routes.SetupRouter(r)

	r.Run(":8080")
}
