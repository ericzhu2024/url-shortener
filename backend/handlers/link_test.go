package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"url-shortener/cache"
	"url-shortener/db"
)

// TestMain 在所有测试运行前执行，负责初始化
func TestMain(m *testing.M) {
	// 设置测试环境变量
	os.Setenv("DB_DSN", "host=localhost user=admin password=password dbname=urlshortener port=5432 sslmode=disable")
	os.Setenv("REDIS_ADDR", "localhost:6379")

	// 初始化数据库和 Redis
	db.Init()
	cache.Init()

	// 运行所有测试
	os.Exit(m.Run())
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/links", CreateLink)
	r.GET("/api/links/:code/stats", GetStats)
	r.GET("/r/:code", RedirectLink)
	return r
}

func TestCreateLink_Valid(t *testing.T) {
	r := setupRouter()

	body := `{"original":"https://google.com"}`
	req, _ := http.NewRequest("POST", "/api/links", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestCreateLink_InvalidURL(t *testing.T) {
	r := setupRouter()

	body := `{"original":"not-a-url"}`
	req, _ := http.NewRequest("POST", "/api/links", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetStats_NotFound(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/api/links/notexist/stats", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestRedirect_NotFound(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/r/notexist", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}