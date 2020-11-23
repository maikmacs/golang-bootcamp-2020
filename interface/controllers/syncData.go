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

	"github.com/maikmacs/golang-bootcamp-2020/config"

	"github.com/gin-gonic/gin"
)

// SyncDataController -
type SyncDataController struct{}

// Episode - Episode Structure
type Episode struct {
	Name    string `json:"name"`
	Season  int    `json:"season"`
	Number  int    `json:"number"`
	Airdate string `json:"airdate"`
}

// Status - Sync Status
func (s SyncDataController) Status(c *gin.Context) {
	response, err := http.Get(config.Structure.Data.Source)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed sync",
		})
	} else {
		var episodes []Episode

		json.Unmarshal(responseData, &episodes)

		writeCSV(episodes)

		c.JSON(http.StatusOK, gin.H{
			"message": "Data synced",
		})
	}
}

func writeCSV(episodes []Episode) {
	file, err := os.Create(config.Structure.Data.Output)
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
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
