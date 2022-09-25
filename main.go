package main

import (
	"assignment-4/models"
	"assignment-4/utils"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

const (
	FILE_NAME        = "raw/stats.json"
	DEFAULT_LOCATION = "Asia/Jakarta"
)

func updateStatsData() {
	file, err := utils.GetFile(FILE_NAME)
	if err != nil {
		log.Panic("Failed to load file", err.Error())
	}

	var status models.Status
	err = json.Unmarshal(file, &status)
	if err != nil {
		log.Panic("Failed to parse JSON", err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	status.Stats.Water = rand.Intn(99) + 1
	rand.Seed(time.Now().UnixNano())
	status.Stats.Wind = rand.Intn(99) + 1
	status.UpdatedAt = time.Now()

	result, err := json.Marshal(status)
	if err != nil {
		log.Panic("Failed to parse JSON", err.Error())
	}

	err = utils.SaveFile(FILE_NAME, result)
	if err != nil {
		log.Panic("Failed to save file", err.Error())
	}
}

func createScheduler() {
	loc, err := time.LoadLocation(DEFAULT_LOCATION)
	if err != nil {
		log.Panic("Failed to load time", err.Error())
	}
	time.Local = loc

	s := gocron.NewScheduler(time.Local)
	s.Every(15).Seconds().Do(updateStatsData)
	s.StartAsync()
}

func statsHandler(ctx *gin.Context) {
	file, err := utils.GetFile(FILE_NAME)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.CreateResponse(
				http.StatusBadRequest,
				"Failed to get stats file",
				nil,
			),
		)
		return
	}

	var status models.Status
	err = json.Unmarshal(file, &status)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.CreateResponse(
				http.StatusBadRequest,
				"Failed to get parse JSON",
				nil,
			),
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		utils.CreateResponse(
			http.StatusOK,
			"Success get data",
			status,
		),
	)
}

func main() {
	createScheduler()

	router := gin.Default()
	router.GET("/", statsHandler)
	router.Run(":3000")
}
