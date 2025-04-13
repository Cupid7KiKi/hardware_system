package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"hardware_system/service"
)

func RegisterRoutes(router *gin.Engine) {
	// 产品价格查询API
	router.POST("/ks/choose/product_price", func(c *gin.Context) {
		productID := c.PostForm("product_id")
		if productID == "" {
			c.JSON(400, gin.H{
				"success": false,
				"msg":     "产品ID不能为空",
			})
			return
		}
		price := service.GetProductSalePrice(productID)

		if service.TransDecimal(price) == decimal.NewFromInt(0) {
			c.JSON(404, gin.H{
				"success": false,
				"msg":     "未找到该产品的价格信息",
			})
			return
		}

		c.JSON(200, gin.H{
			"success": true,
			"msg":     "ok",
			"data":    service.TransDecimal(price),
		})
	})

	// 可以在这里添加更多API路由...
}
