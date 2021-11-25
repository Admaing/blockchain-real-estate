package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/togettoyou/blockchain-real-estate/application/blockchain"
	_ "github.com/togettoyou/blockchain-real-estate/application/docs"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/setting"
	"github.com/togettoyou/blockchain-real-estate/application/routers"
	"github.com/togettoyou/blockchain-real-estate/application/service"
	"log"
	"net/http"
	"time"
)

func init() {
	// 初始化时会通过app.init将相关配置文件通过Mapto映射到结构体中
	setting.Setup()
}

// @title 基于区块链技术的房地产交易系统api文档
// @version 1.0
// @description 基于区块链技术的房地产交易系统api文档
// @contact.name togettoyou
// @contact.email zoujh99@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		log.Printf("时区设置失败 %s", err)
	}
	log.Printf("12312")
	time.Local = timeLocal
	blockchain.Init()
	go service.Init()
	// 设置模式，test,realease , debug三种模式
	gin.SetMode(setting.ServerSetting.RunMode)
	// 初始化路由信息
	routersInit := routers.InitRouter()
	// 服务器超时时间
	readTimeout := setting.ServerSetting.ReadTimeout
	// 服务器写入时间
	writeTimeout := setting.ServerSetting.WriteTimeout
	// 服务器运行端口
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	// 返回服务控制信息
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	// server.ListenAndServe()创建服务并且监听
	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}
}
