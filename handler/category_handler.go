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

type category struct {
	service service.Category
}

func NewCategoryHandler(service service.Category) *category {
	return &category{service}
}

// Show godoc
// @Summary      Show a category
// @Description  get all data category
// @Tags         category
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Category
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /category [get]
func (h *category) Show(c *gin.Context) {
	category, err := h.service.Show()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.Category(category)
	response := helper.ApiResponse("Data berhasil di tampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

// Create godoc
// @Summary      Add a category
// @Description  add data category
// @Tags         category
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Category
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /category [get]
func (h *category) Create(c *gin.Context) {
	var input payload.CategoryPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	category, err := h.service.Create(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.CategoryRow(category)
	response := helper.ApiResponse("Data berhasil ditambahkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *category) Edit(c *gin.Context) {
	var input payload.CategoryPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.service.Update(id, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di ubah", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.CategoryRow(category)
	response := helper.ApiResponse("Data berhasil di ubah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *category) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.service.Find(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.CategoryRow(category)
	response := helper.ApiResponse("Data berhasil ditampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *category) Delete(c *gin.Context) {
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
