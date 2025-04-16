package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"hardware_system/service"
	"net/http"
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
	// 处理聊天请求
	router.POST("/ks/info/chat", func(c *gin.Context) {
		var req service.RequestBody
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "错误请求"})
			return
		}

		// 这里可以替换为实际的AI处理逻辑
		//reply := service.ProcessAIMessage(req.Message)
		reply, err := service.CallDeepseekAPI(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "错误response"})
		}

		messages := []service.Message{}

		// 更新聊天记录
		messages = append(messages,
			service.Message{Sender: "user", Content: req.Messages[0].Content},
			service.Message{Sender: "ai", Content: reply.Choices[0].Message.Content},
		)
		fmt.Println("Messages中的信息：", messages)
		// 返回AI回复
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"Reply":   reply.Choices[0].Message.Content,
		})
	})
}
