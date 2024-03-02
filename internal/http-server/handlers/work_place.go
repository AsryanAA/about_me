package handlers

import (
	"about_me/internal/http-server/services"
	"about_me/internal/models"
	"about_me/internal/storage/sqlite"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Создание новой записи Место работы
// @Desription Создает новую запись в таблице
// @Tags Место работы (work_place)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.WorkPlace
// @Failure 404
// @Router /create [post]
func CreateWorkPlace(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newWorkPlace models.WorkPlace
		if err := ctx.BindJSON(&newWorkPlace); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Некорректные данные",
			})
			return
		}

		err := services.CreateWorkPlace(newWorkPlace, storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Создание не произошло",
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusCreated, newWorkPlace)
			return
		}
	}
}

// @Summary Чтение всех записей Места работы
// @Desription Читает записи из таблицы
// @Tags Места работы (work_places)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.WorkPlace
// @Failure 404
// @Router /read [get]
func ReadWorkPlaces(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workPlaces, err := services.ReadWorkPlaces(storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Чтение не произошло",
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, workPlaces)
			return
		}
	}
}

// @Summary Обновление записи Место работы
// @Desription Обноляет запись в таблице
// @Tags Место работы (work_place)
// @Param auth_id path string true "AuthId"
// @Param password_web path string true "PasswordWeb"
// @Success 200 {object} models.WorkPlace
// @Failure 404
// @Router /update [patch]
func UpdateWorkPlace(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateWorkPlace models.WorkPlace
		if err := ctx.BindJSON(&updateWorkPlace); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Некорректные данные",
			})
			return
		}

		err := services.UpdateWorkPlace(updateWorkPlace, storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Обноление не произошло",
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, updateWorkPlace.Id)
			return
		}
	}
}
