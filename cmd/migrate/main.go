package main

import (
	"galaxy/pkg/config"
	"galaxy/pkg/database"
	"log"
)

func main() {
	// 加载配置
	config.Load("configs/config.yaml")

	// 初始化数据库
	database.Init()

	// 执行自动迁移
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	log.Println("数据库迁移完成")
}
