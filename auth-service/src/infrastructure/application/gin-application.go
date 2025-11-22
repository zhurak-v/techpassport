package application

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/zhurak-v/techpassport/auth-service/src/adapters"
	"github.com/zhurak-v/techpassport/auth-service/src/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func GinApplication(
	migrations database.IMigrations,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery(), gzip.Gzip(gzip.BestSpeed))
	_ = app.SetTrustedProxies([]string{})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router := adapters.NewRouterSetup(app)
	router.SetupRoutes()

	migrations.Up()

	return app
}
