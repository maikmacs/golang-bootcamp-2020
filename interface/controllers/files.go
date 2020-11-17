package controllers

import (
	"github.com/gin-gonic/gin"
)

type FilesController struct{}

func (f FilesController) Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Files Controller",
	})
}
