package router

import (
	"net/http"

	"github.com/maikmacs/golang-bootcamp-2020/config"
	"github.com/maikmacs/golang-bootcamp-2020/interface/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter - Start Server and Router
func InitRouter() {
	r := newRouter()
	r.Run(":" + config.Structure.Server.Port)
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
