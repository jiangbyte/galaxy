package database

import (
	"fmt"
	models2 "galaxy/internal/models"
	"galaxy/pkg/config"
	"galaxy/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	cfg := config.Get().Database.Postgres

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger().LogMode(gormLogger.Warn), // 只显示警告和错误
	})
	if err != nil {
		logger.Error().
			Str("host", cfg.Host).
			Int("port", cfg.Port).
			Err(err).
			Msg("Connection failed")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error().
			Str("host", cfg.Host).
			Int("port", cfg.Port).
			Err(err).
			Msg("Failed to get database instance")
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	logger.Connected("Postgres")
	logger.Database().
		Str("host", cfg.Host).
		Int("port", cfg.Port).
		Str("db", cfg.DBName).
		Msg("Connection details")
}

func GetDB() *gorm.DB {
	if db == nil {
		logger.Error().
			Msg("Database not initialized")
	}
	return db
}

func AutoMigrate() error {
	// 先禁用钩子
	db = db.Session(&gorm.Session{
		SkipHooks: true,
	})

	tables := []interface{}{
		// ==================== 用户认证模块 ====================
		&models2.AuthAccount{},
		&models2.UserInfo{},
		&models2.UserProfile{},
		&models2.UserPreference{},
		&models2.UserStats{},
		&models2.VipInfo{},
		&models2.UserBadge{},
		&models2.BadgeConfig{},
		&models2.VipPrivilege{},
		&models2.VipLevelConfig{},
		&models2.AuthAccountRole{},
		&models2.AuthGroup{},
		&models2.AuthRole{},
		&models2.AuthRoleMenu{},

		// ==================== 系统管理模块 ====================
		&models2.SysDict{},
		&models2.SysLog{},
		&models2.SysMenu{},

		// ==================== 配置 ====================
		&models2.ConfigGroup{},
		&models2.ConfigItem{},

		// ==================== CMS内容管理模块 ====================
		&models2.ContentArticle{},
		&models2.ContentBanner{},
		&models2.ContentCategory{},
		&models2.ContentNotice{},
		&models2.ContentTag{},

		// ==================== 题目管理模块 ====================
		&models2.ProblemInfo{},
		&models2.ProblemTagRel{},
		&models2.ProblemTestCase{},

		// ==================== 提交与判题模块 ====================
		&models2.JudgeSubmit{},
		&models2.JudgeCase{},

		// ==================== 竞赛管理模块 ====================
		&models2.ContestInfo{},
		&models2.ContestAuth{},
		&models2.ContestParticipant{},
		&models2.ContestProblem{},

		// ==================== 用户学习记录模块 ====================
		&models2.RecordCodeLibrary{},
		&models2.RecordSolved{},

		// ==================== 代码相似度检测模块 ====================
		&models2.SimilarityStat{},
		&models2.SimilaritySegment{},
	}

	//return db.AutoMigrate(m...)

	for i, table := range tables {
		logger.Info().
			Str("table", fmt.Sprintf("%T", table)).
			Msg("Migrating")
		if err := db.AutoMigrate(table); err != nil {
			logger.Error().
				Str("table", fmt.Sprintf("%T", table)).
				Int("index", i).
				Err(err).
				Msg("Migration failed")
			return err
		}
	}

	logger.Info().
		Msg("Migration completed")
	return nil
}
