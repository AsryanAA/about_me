package view

import (
	"about_me/internal/http-server/services"
	"about_me/internal/storage/sqlite"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadWorkPlaces(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, _ := services.ReadWorkPlaces(storage)
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"data": data,
		})
	}
}
