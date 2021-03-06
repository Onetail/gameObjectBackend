package app

import (
	"log"
	"strconv"

	_ "gameObjectBackend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HTTPServer struct {
	App    *App
	engine *gin.Engine
	host   string
	port   int
}

func (hs *HTTPServer) Init(app *App) {

	hs.App = app
	hs.host = viper.GetString("http_server.host")
	hs.port = viper.GetInt("http_server.port")

	// Initializing HTTP server
	hs.engine = gin.Default()

	url := ginSwagger.URL("/swagger/doc.json")
	hs.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, "DELETE")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	hs.engine.Use(cors.New(corsConfig))
}

func (hs *HTTPServer) Start() {
	log.Printf("Listening on http://%s:%d", hs.host, hs.port)
	hs.engine.Run(hs.host + ":" + strconv.Itoa(hs.port))
}

func (hs *HTTPServer) GetEngine() *gin.Engine {
	return hs.engine
}
