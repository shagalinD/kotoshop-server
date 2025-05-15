package handlers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetProductImage(c *gin.Context) {
	// 1. Получаем имя файла из параметров URL
	filename := c.Param("filename")
	if filename == "" {
			c.JSON(400, gin.H{"error": "filename is required"})
			return
	}

	// 2. Безопасное формирование пути к файлу
	safePath := filepath.Join("productImages", filepath.Base(filename))
	
	// 3. Проверка существования файла
	if _, err := os.Stat(safePath); os.IsNotExist(err) {
			c.JSON(404, gin.H{"error": "image not found"})
			return
	}

	// 4. Определение MIME-типа
	contentType := "image/jpeg"
	switch filepath.Ext(filename) {
	case ".png":
			contentType = "image/png"
	case ".gif":
			contentType = "image/gif"
	case ".webp":
			contentType = "image/webp"
	}

	// 5. Отправка файла с кэшированием
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Header("Content-Type", contentType)
	c.File(safePath)
}