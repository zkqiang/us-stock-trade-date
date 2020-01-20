package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"us-stock-trade-date/config"
	"us-stock-trade-date/data"
	"us-stock-trade-date/router"
)

func main() {
	// 解析命令行参数
	env := flag.String("env", "dev", "running environment")
	flag.Parse()
	if *env != "test" && *env != "prod" {
		*env = "dev"
	}

	// 初始化配置
	config.Init(*env)

	// 生成节日数据
	log.Print("Wait for generate holidays...")
	data.GenHolidays()

	// 初始化路由
	router_ := router.New()

	// 设置监听的端口
	port := os.Getenv("PORT")
	log.Printf("Start listen serve: http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, router_)
	if err != nil {
		log.Fatal("Start error: ", err)
	}
}
