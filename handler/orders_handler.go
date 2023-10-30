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

type orders struct {
	service            service.Orders
	serviceOrderDetail service.OrderDetail
}

func NewOrdersHandler(service service.Orders, serviceOrderDetail service.OrderDetail) *orders {
	return &orders{service, serviceOrderDetail}
}

func (h *orders) Show(c *gin.Context) {
	limit := 10
	page, _ := strconv.Atoi(c.Query("page"))
	offset := limit * (page - 1)

	ordersAll, err := h.service.Show()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	orders, err := h.service.ShowByLimit(limit, offset)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tampilkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	pagination := *helper.PaginationFormat(len(ordersAll), limit, page)

	formatter := mapping.Orders(orders, pagination)
	response := helper.ApiResponse("Data berhasil di tampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *orders) Create(c *gin.Context) {
	var input payload.OrdersPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	uid := c.Value("uid")
	orders, err := h.service.Create(input, uid.(int))
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	_, err = h.serviceOrderDetail.CreateMultiple(input.Order, orders.Id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.OrdersRow(orders)
	response := helper.ApiResponse("Data berhasil ditambahkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *orders) Edit(c *gin.Context) {
	var input payload.OrdersPayload
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data yang dimasukan salah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	id, _ := strconv.Atoi(c.Param("id"))
	orders, err := h.service.Update(id, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di ubah", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	_, err = h.serviceOrderDetail.DeleteByOrderId(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	_, err = h.serviceOrderDetail.CreateMultiple(input.Order, orders.Id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data gagal di tambahkan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.OrdersRow(orders)
	response := helper.ApiResponse("Data berhasil di ubah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *orders) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	orders, err := h.service.Find(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	orderDetail, err := h.serviceOrderDetail.ShowByOrderId(id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}
		response := helper.ApiResponse("Data tidak ditemukan", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := mapping.OrdersDetails(orders, orderDetail)
	response := helper.ApiResponse("Data berhasil ditampilkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *orders) Delete(c *gin.Context) {
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
