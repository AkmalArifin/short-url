package routes

import (
	"net/http"

	"github.com/AkmalArifin/short-url/internal/models"
	"github.com/AkmalArifin/short-url/internal/utils"
	"github.com/gin-gonic/gin"
)

func getUrls(c *gin.Context) {
	urls, err := models.GetURLS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot query data"})
		return
	}

	c.JSON(http.StatusOK, urls)
}

func getStatURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := models.GetStat(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot query data"})
		return
	}

	c.JSON(http.StatusOK, url)
}

func createUrl(c *gin.Context) {
	var url models.URL
	err := c.ShouldBindJSON(&url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	url.ShortCode.SetValid(utils.GenerateShortCode())

	err = url.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot save data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "data created", "url": url})
}

func updateURL(c *gin.Context) {
	var url models.URL
	err := c.ShouldBindJSON(&url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	shortCode := c.Param("shortCode")
	retrieveUrl, err := models.GetStat(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot query data"})
		return
	}

	retrieveUrl.URL.SetValid(url.URL.String)
	err = retrieveUrl.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot save data"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "data updated"})
}

func redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := models.GetStat(shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	url.AccessCount.Int64 += 1
	err = url.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot save data"})
		return
	}

	c.Redirect(http.StatusSeeOther, url.URL.ValueOrZero())
}

func deleteUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := models.GetStat(shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err = url.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot delete data"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "data deleted"})
}
