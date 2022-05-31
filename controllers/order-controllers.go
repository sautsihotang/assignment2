package controllers

import (
	"assignment2/config"
	"assignment2/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	db := config.GetDB()

	var data models.CreateOrder

	var items []models.Item

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(data)
	orders := models.Order{
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
	}

	db.Create(&orders)

	id_order := orders.ID
	for _, v := range data.Item {
		item := models.Item{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     orders.ID,
		}

		items = append(items, item)
	}
	fmt.Println(items)
	result := db.Create(&items)
	log.Println(id_order, result.RowsAffected)

	var returnData interface{}

	returnData = models.Order{
		ID:           items[0].OrderID,
		CustomerName: data.CustomerName,
		OrderedAt:    time.Now(),
		Item:         items,
	}

	c.JSON(http.StatusOK, returnData)

}

func GetAllOrder(c *gin.Context) {
	db := config.GetDB()

	orders := []models.Order{}

	err := db.Preload("Item").Find(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)

}

func GetOrderById(c *gin.Context) {
	db := config.GetDB()

	order := models.Order{}

	var id_order = c.Param("orderID")

	err := db.Where("id = ?", id_order).Find(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	items := []models.Item{}

	err = db.Where("order_id = ?", id_order).Find(&items).Error
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	resp := models.CreateOrder{}
	resp.CustomerName = order.CustomerName
	resp.OrderedAt = order.OrderedAt
	resp.Item = items
	c.JSON(http.StatusOK, resp)
}

func DeleteOrderById(c *gin.Context) {
	db := config.GetDB()

	id_str := c.Param("orderID")
	id_param, _ := strconv.Atoi(id_str)

	order := models.Order{}
	item := models.Item{}

	err := db.Where("order_id = ?", id_param).Delete(&item).Error

	rows := db.Where("ID = ?", id_param).Delete(&order).RowsAffected

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Msg": "Maaf... Data Tidak Ditemukan",
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Msg": "Data Berhasil Di Hapus",
	})
}

func UpdateOrder(c *gin.Context) {
	db := config.GetDB()
	id_str := c.Param("orderID")
	id_param, _ := strconv.Atoi(id_str)

	var data = models.CreateOrder{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		ID:           uint(id_param),
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
		Item:         data.Item,
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	c.JSON(http.StatusOK, order)
}

func DeleteAllOrder(c *gin.Context) {
	db := config.GetDB()

	orders := models.Order{}

	err := db.Delete(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Msg": "Data Berhasil Di Hapus",
	})

}
