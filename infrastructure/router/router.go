package router

import (
	"golang-bootcamp/interface/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

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
