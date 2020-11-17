package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SyncDataController struct{}

func (s SyncDataController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Sync Controller",
	})

}
