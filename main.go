package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/components/login"
	"hardware_system/routes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	// 引入theme2登录页面主题，如不用，可以不导入
	_ "github.com/GoAdminGroup/components/login/theme2"
	_ "github.com/GoAdminGroup/themes/adminlte" // ui theme
	_ "github.com/GoAdminGroup/themes/sword"    // ui theme
	_ "github.com/GoAdminGroup/themes/sword/separation"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"hardware_system/models"
	"hardware_system/pages"
	"hardware_system/service"
	"hardware_system/tables"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	// 先注册自定义路由
	routes.RegisterRoutes(r)

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	// 使用登录页面组件
	login.Init(login.Config{
		Theme:         "theme2",
		CaptchaDigits: 4, // 使用图片验证码，这里代表多少个验证码数字
		// 使用腾讯验证码，需提供appID与appSecret
		// TencentWaterProofWallData: login.TencentWaterProofWallData{
		//    AppID:"",
		//    AppSecret: "",
		// }
	})

	// 引入我们定义的login组件
	//template.AddLoginComp(login.Get())

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}
	service.SetDb(eng.MysqlConnection())

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/ks", pages.GetIndex2)
	eng.HTMLFile("GET", "/ks/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})
	//r.GET("/ks/login", func(c *gin.Context) {
	//	// 获取客户端 IP（自动处理 X-Forwarded-For）
	//	clientIP := c.ClientIP()
	//
	//	// 确保是 IPv4
	//	ip := net.ParseIP(clientIP)
	//	if ip != nil && ip.To4() != nil {
	//		clientIP = ip.To4().String()
	//	}
	//
	//	c.String(200, "客户端 IPv4: %s", clientIP)
	//})
	//eng.HTMLFile("GET", "ks/info/chat", "./public/assets/template/test.html", map[string]interface{}{
	//	"msg": "test-chat",
	//})
	// 初始化聊天记录
	messages := []service.Message{
		{Sender: "ai", Content: "你好！我是AI助手，有什么可以帮你的吗？"},
	}
	eng.HTMLFile("GET", "ks/info/chat", "./public/assets/template/chat.tmpl", map[string]interface{}{
		"Messages": messages,
	})

	// 处理聊天请求
	r.POST("/ks/info/chat", func(c *gin.Context) {
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

	//eng.AddConfig(config.Config{})

	//// 加载模板
	//r.LoadHTMLGlob("public/pages/404.tmpl")
	//
	//// 配置 404 处理
	//r.NoRoute(func(c *gin.Context) {
	//	c.HTML(http.StatusNotFound, "public/pages/404.tmpl", gin.H{
	//		"title": "自定义 404 页面",
	//	})
	//})

	models.Init(eng.MysqlConnection())

	srv := &http.Server{
		Addr:    ":9022",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Print("closing database connection")
	eng.MysqlConnection().Close()

	log.Println("Server exiting")
}
