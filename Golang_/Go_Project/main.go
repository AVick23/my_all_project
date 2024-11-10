package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер
	r := gin.Default()

	// Указываем папку для статических файлов
	r.Static("/static", "./static")

	// Указываем папку для HTML-шаблонов
	r.LoadHTMLGlob("templates/*")

	// Определяем маршрут для главной страницы
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Главная страница",
		})
	})

	// Определяем маршрут для страницы каталогов
	r.GET("/catalog", func(c *gin.Context) {
		c.HTML(200, "catalog.html", gin.H{
			"title": "Каталог магазина",
		})
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
