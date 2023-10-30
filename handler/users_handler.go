package handler

import (
	"dbo/helper"
	"dbo/mapping"
	"dbo/payload"
	"dbo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type users struct {
	service     service.Users
	authService service.Auth
}

func NewUsersHandler(service service.Users, authService service.Auth) *users {
	return &users{service, authService}
}

func (h *users) Login(c *gin.Context) {
	var input payload.Login

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.service.Login(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Login gagal", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Generate token gagal", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mapping.Auth(loggedinUser, token)
	response := helper.ApiResponse("Login berhasil", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *users) Show(c *gin.Context) {
	users, err := h.service.Show()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.Users(users)
	response := helper.ApiResponse("Data berhasil di tampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *users) Create(c *gin.Context) {
	var input payload.UsersPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	users, err := h.service.Create(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.UsersRow(users)
	response := helper.ApiResponse("Data berhasil ditambahkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *users) Edit(c *gin.Context) {
	var input payload.UsersPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := h.service.Update(id, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di ubah", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.UsersRow(users)
	response := helper.ApiResponse("Data berhasil di ubah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *users) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := h.service.Find(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.UsersRow(users)
	response := helper.ApiResponse("Data berhasil ditampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *users) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.service.Delete(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.ApiResponse("Data berhasil dihapus", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}
