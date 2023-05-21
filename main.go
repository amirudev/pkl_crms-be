package main

import (
	"github.com/amirudev/pkl_crms-be/config"
	"github.com/amirudev/pkl_crms-be/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	err := config.InitDB(
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
	)

	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", config.DBConn)
		c.Next()
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	}))

	routes.SetupRouter(r)

	r.Run()
}
