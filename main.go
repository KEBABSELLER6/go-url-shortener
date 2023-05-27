package main

import (
	"net/http"

	"github.com/KEBABSELLER6/go-url-shortener/redis"
	"github.com/gin-gonic/gin"
)

type ShortenedUrl struct {
	ShortID string `json:"shortID"`
	Url     string `json:"original"`
}

type UrlRequest struct {
	Url string `json:"url"`
}

func createShortenedUrl(c *gin.Context) {
	var newUrl UrlRequest

	if err := c.ShouldBindJSON(&newUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newShortId, redisErr := redis.SetIfNotExist(newUrl.Url)

	if redisErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": redisErr})
	} else {
		newShortenedUrl := ShortenedUrl{ShortID: newShortId, Url: newUrl.Url}
		c.JSON(http.StatusCreated, newShortenedUrl)
	}

}

func main() {
	router := gin.Default()
	router.POST("/shorten", createShortenedUrl)

	router.Run("localhost:8080")
}
