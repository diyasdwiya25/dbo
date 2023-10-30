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

type customers struct {
	service service.Customers
}

func NewCustomersHandler(service service.Customers) *customers {
	return &customers{service}
}

func (h *customers) Show(c *gin.Context) {
	limit := 10
	page, _ := strconv.Atoi(c.Query("page"))
	offset := limit * (page - 1)

	customersAll, err := h.service.Show()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	customers, err := h.service.ShowByLimit(limit, offset)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	pagination := *helper.PaginationFormat(len(customersAll), limit, page)

	formatter := mapping.Customers(customers, pagination)
	response := helper.ApiResponse("Data berhasil di tampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *customers) Create(c *gin.Context) {
	var input payload.CustomersPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	customers, err := h.service.Create(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.CustomersRow(customers)
	response := helper.ApiResponse("Data berhasil ditambahkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *customers) Edit(c *gin.Context) {
	var input payload.CustomersPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	id, _ := strconv.Atoi(c.Param("id"))
	customers, err := h.service.Update(id, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di ubah", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.CustomersRow(customers)
	response := helper.ApiResponse("Data berhasil di ubah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *customers) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customers, err := h.service.Find(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.CustomersRow(customers)
	response := helper.ApiResponse("Data berhasil ditampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *customers) Delete(c *gin.Context) {
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
