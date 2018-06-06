package main

import (
	"db"
	"log"
	"models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	session *db.Session
	err     error
)

// MongoCon connection string information
const (
	MongoCon = "mongodb://localhost:27017"
)

func main() {

	session, err = db.New(MongoCon, "local")

	if err != nil {
		log.Fatal("mongodb database is not connected")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/dialog", func(c *gin.Context) {
		dialog := models.Dialogs{}
		dialog.Timestamp = time.Now().String()
		dialog.Status = "Active"
		if err := c.ShouldBind(&dialog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := session.Insert("dialogs", dialog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "successfully inserted",
		})
	})

	r.GET("/dialog/:skip/:limit", func(c *gin.Context) {

		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		skip, err := strconv.Atoi(c.Param("skip"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		dialogs, err := session.FindAllByLimitOffset("dialogs", skip, limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if dialogs == nil {
			c.JSON(http.StatusNoContent, gin.H{"error": "No data"})
		}
		c.JSON(http.StatusOK, dialogs)
	})

	r.POST("/dialog/{dialog_id}/voice", func(c *gin.Context) {

	})

	r.POST("/dialog/{dialog_id}/video", func(c *gin.Context) {

	})

	r.POST("/dialog/{dialog_id}/selfie", func(c *gin.Context) {

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
