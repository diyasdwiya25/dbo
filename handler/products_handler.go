package handler

import (
	"dbo/helper"
	"dbo/mapping"
	"dbo/payload"
	"dbo/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type products struct {
	service service.Products
}

func NewProductsHandler(service service.Products) *products {
	return &products{service}
}

func (h *products) Show(c *gin.Context) {
	products, err := h.service.Show()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.Products(products)
	response := helper.ApiResponse("Data berhasil di tampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *products) Create(c *gin.Context) {
	var input payload.ProductsPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	fmt.Println(input)
	products, err := h.service.Create(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.ProductsRow(products)
	response := helper.ApiResponse("Data berhasil ditambahkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *products) Edit(c *gin.Context) {
	var input payload.ProductsPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	id, _ := strconv.Atoi(c.Param("id"))
	products, err := h.service.Update(id, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di ubah", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.ProductsRow(products)
	response := helper.ApiResponse("Data berhasil di ubah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *products) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	products, err := h.service.Find(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.ProductsRow(products)
	response := helper.ApiResponse("Data berhasil ditampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *products) Delete(c *gin.Context) {
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
