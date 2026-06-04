package handlers

import (
	"net/http"
	"net/url"
	"time"
	"url-shortener/cache"
	"url-shortener/db"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"

)

func isValidURL(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}
	return true
}

func CreateLink(c *gin.Context) {
	var input struct {
		Original string `json:"original"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Original == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid URL"})
		return
	}

	if !isValidURL(input.Original) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid URL starting with http:// or https://"})
		return
	}


	if code, err := cache.GetOriginal(input.Original); err == nil {
		cache.Refresh("original:"+input.Original, 24*time.Hour)
		cache.Refresh(code, 24*time.Hour)
		c.JSON(http.StatusOK, gin.H{
			"code":      code,
			"short_url": "http://localhost:8080/r/" + code,
			"original":  input.Original,
		})
		return
	}


	var existing models.Link
	if result := db.DB.Where("original = ?", input.Original).First(&existing); result.Error == nil {
		cache.Set(existing.Code, existing.Original, 24*time.Hour)
		cache.SetOriginal(existing.Original, existing.Code, 24*time.Hour)
		c.JSON(http.StatusOK, gin.H{
			"code":      existing.Code,
			"short_url": "http://localhost:8080/r/" + existing.Code,
			"original":  existing.Original,
		})
		return
	}

	var code string
	for {
		code = utils.GenerateCode()
		var clash models.Link
		if db.DB.Where("code = ?", code).First(&clash).Error != nil {
			break
		}
	}

	link := models.Link{
		Code:     code,
		Original: input.Original,
	}
	db.DB.Create(&link)

	cache.Set(code, input.Original, 24*time.Hour)
	cache.SetOriginal(input.Original, code, 24*time.Hour)

	c.JSON(http.StatusCreated, gin.H{
		"code":      code,
		"short_url": "http://localhost:8080/r/" + code,
		"original":  input.Original,
	})
}

func RedirectLink(c *gin.Context) {
	code := c.Param("code")


	original, err := cache.Get(code)
	if err != nil {

		var link models.Link
		if result := db.DB.Where("code = ?", code).First(&link); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
			return
		}

		original = link.Original

		cache.Set(code, original, 24*time.Hour)

		db.DB.Model(&link).Update("clicks", link.Clicks+1)
	} else {
		go func() {
			var link models.Link
			db.DB.Where("code = ?", code).First(&link)
			db.DB.Model(&link).Update("clicks", link.Clicks+1)
		}()
	}

	c.Redirect(http.StatusFound, original)
}

// GetStats returns click statistics for a short link
func GetStats(c *gin.Context) {
	code := c.Param("code")

	var link models.Link
	if result := db.DB.Where("code = ?", code).First(&link); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, link)
}