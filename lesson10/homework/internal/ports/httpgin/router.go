package httpgin

import (
	"github.com/gin-gonic/gin"
	"homework10/internal/app"
)

func AppRouter(r *gin.RouterGroup, a app.App) {
	r.POST("/ads", createAd(a)) // Метод для создания объявления (ad)
	r.POST("/users", createUser(a))
	r.PUT("/ads/:ad_id/status", changeAdStatus(a)) // Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
	r.PUT("/users/:user_id", changeUser(a))
	r.PUT("/ads/:ad_id", updateAd(a)) // Метод для обновления текста(Text) или заголовка(Title) объявления
	r.GET("/ads", listAds(a))
	r.GET("/ads/by_title", getAdsByTitle(a))
	r.GET("/ads/:ad_id", getAdById(a))
}
