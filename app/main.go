package main

import (
	Handlers "gin-boilerplate/app/handlers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	log "github.com/sirupsen/logrus"
	docs "gin-boilerplate/docs"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	gin.ForceConsoleColor()

	r := gin.Default()
	r.SetTrustedProxies([]string{"*"})

	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.Use(Handlers.Auth())
		v1.GET("/", Handlers.GetIndex)
		v1.GET("/albums", Handlers.GetAlbums)
		v1.GET("/albums/:id", Handlers.GetAlbumByID)
		v1.POST("/albums", Handlers.PostAlbums)
	}

	// open swagger at http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Run("localhost:8080")
}
