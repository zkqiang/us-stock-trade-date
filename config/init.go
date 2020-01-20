package config

import (
	"github.com/joho/godotenv"
)

// 初始化配置项
func Init(env string) {
	// 先读取根目录的环境配置
	_ = godotenv.Load()

	// 再根据环境变量读取不同的环境配置
	_ = godotenv.Load("config/env/" + env + ".conf")
}
