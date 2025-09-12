package dao

import (
	"database/sql"
	"fmt"
	"keymesh/models"
	"keymesh/utils/config"
	"log"

	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func init() {
	initClient()
}

func initClient() {
	once.Do(func() {
		initDBWithRetry()
	})
}

// initDBWithRetry 函数用于初始化数据库连接，并包含重试机制
func initDBWithRetry() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&readTimeout=10s&writeTimeout=10s",
		config.DB_User, config.DB_Pass, config.DB_Host, config.DB_Port, config.DB_Name,
	)

	var err error
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})
		if err == nil {
			break
		}
		if !fiber.IsChild() {
			log.Printf("数据库连接失败，重试 %d/%d: %v", i+1, maxRetries, err)
		}
		time.Sleep(time.Second * 2)
	}

	if err != nil {
		if !fiber.IsChild() {
			log.Fatalf("数据库连接最终失败: %v", err)
		}
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		if !fiber.IsChild() {
			log.Fatalf("获取 sqlDB 失败: %v", err)
		}
	}

	// 优化连接池配置
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)

	// 健康检查
	go healthCheck(sqlDB)

	// 自动迁移
	if err := AutoMigrate(); err != nil {
		if !fiber.IsChild() {
			log.Fatalf("数据库表创建失败: %v", err)
		}
	}

	if !fiber.IsChild() {
		log.Println("数据库连接成功")
	}
}

// 健康检查goroutine
func healthCheck(sqlDB *sql.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := sqlDB.Ping(); err != nil {
			if !fiber.IsChild() {
				log.Printf("数据库健康检查失败: %v", err)
			}
			// 尝试重新连接
			initClient()
			return
		}
	}
}

// 获取数据库连接（带重试）
func GetDB() *gorm.DB {
	if db == nil {
		initClient()
	}
	return db
}

// 自动注册表
func AutoMigrate() error {
	return db.AutoMigrate(
		&models.User{},
		&models.Record{},
		&models.OAuthProvider{},
	)
}
