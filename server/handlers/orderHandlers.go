package handlers

import (
	"kotoshop/models"
	"kotoshop/postgres"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Address string `json:"address"`
}



// CreateOrder godoc
// @Summary      Создает заказ
// @Description  Создает заказ пользователя из продуктов его корзины
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param address body CreateOrderRequest true "Данные заказа"
// @Param Authorization header string true "Токен в формате Bearer {token}" default(Bearer )
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/order/create [post]
func CreateOrder(c *gin.Context) {
	userID := c.GetUint("userID")

	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("error on parsing address: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"error on parsing order address",
		})
		return 
	}

	var cart models.Cart 

	if err := postgres.DB.Where("user_id = ?", userID).Preload("Items").First(&cart).Error; err != nil {
		log.Printf("error on getting user cart: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on getting user cart",
		})
		return 
	}

	order := models.Order {
		UserID: userID,
		Status: "Создан",
		Total: cart.Total,
		Address: req.Address,
		Date: time.Now(),
	}

	if err := postgres.DB.Create(&order).Error; err != nil {
		log.Printf("error on creating order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on creaing order",
		})
		return 
	}

	var orderItems []models.OrderItem 
	for _, item := range cart.Items {
		orderItems = append(orderItems, models.OrderItem{
			OrderID: order.ID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
		})
	}

	if err := postgres.DB.Create(&orderItems).Error; err != nil {
		log.Printf("error on creating order items %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on creating order items",
		})
		return 
	}

	if err := postgres.DB.Delete(&cart).Error; err != nil {
		log.Printf("error on deleting user cart %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on deleting user cart",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"order successfully created",
		"order_number":order.OrderNumber,
		"order_status":order.Status,
		"date":order.Date.Format("2006-01-02"),
	})
}

// GetUserOrders godoc
// @Summary      Получить заказы пользователя
// @Description  Получает все заказы пользователя
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Токен в формате Bearer {token}" default(Bearer )
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/order/get_all [get]
func GetUserOrders(c *gin.Context) {
	userID := c.GetUint("userID")

	var orders []models.Order 

	if err := postgres.DB.Select("total, address, status, order_number, date").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on getting user orders",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}