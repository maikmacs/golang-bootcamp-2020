package router

import (
	"golang-bootcamp/config"
	"golang-bootcamp/interface/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter - Start Server and Router
func InitRouter() {
	config := config.GetConfig()
	r := newRouter()
	r.Run(":" + config.GetString("server.port"))
}

func newRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	sync := new(controllers.SyncDataController)
	router.GET("/sync-data", sync.Status)

	return router
}
