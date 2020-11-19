package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang-bootcamp/config"

	"github.com/gin-gonic/gin"
)

type SyncDataController struct{}

type Episode struct {
	Name    string `json:"name"`
	Season  int    `json:"season"`
	Number  int    `json:"number"`
	Airdate string `json:"airdate"`
}

func (s SyncDataController) Status(c *gin.Context) {
	config := config.GetConfig()
	response, err := http.Get(config.GetString("data.endPoint"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed sync",
		})

	} else {
		var episodes []Episode

		json.Unmarshal(responseData, &episodes)

		file, err := os.Create(config.GetString("data.outputPath"))
		checkError("Cannot create file", err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, episode := range episodes {
			var row []string

			row = append(row, episode.Name)
			row = append(row, strconv.Itoa(episode.Season))
			row = append(row, strconv.Itoa(episode.Number))
			row = append(row, episode.Airdate)

			err := writer.Write(row)
			checkError("Cannot write to file", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data synced",
		})
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
