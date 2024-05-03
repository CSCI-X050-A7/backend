package database

import (
	"time"

	"github.com/CSCI-X050-A7/backend/app/controller"
	"github.com/CSCI-X050-A7/backend/app/model"
	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// DB gorm connector
var DB *gorm.DB

func GetGormConfig() gorm.Config {
	logLevel := gorm_logger.Silent
	if config.Conf.DBEcho {
		logLevel = gorm_logger.Info
	}
	gormConfig := gorm.Config{Logger: gorm_logger.New(
		logrus.StandardLogger(),
		gorm_logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)}
	return gormConfig
}

func MigrateDatabase() error {
	return DB.AutoMigrate(
		&model.Movie{},
		&model.User{},
		&model.Booking{},
		&model.Card{},
		&model.Promotion{},
		&model.Ticket{},
	)
}

func CreateDemoUser() error {
	if !config.Conf.Debug {
		return nil
	}
	password, _ := controller.GeneratePasswordHash([]byte("123456"))
	user := model.User{
		IsActive: true,
		IsAdmin:  true,
		UserName: "demo",
		Email:    "demo@example.com",
		Password: password,
		Name:     "Demo User",
	}
	// check user already exists
	result := DB.Where(model.User{Email: user.Email}).
		Or(model.User{UserName: user.UserName}).Find(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 0 {
		return nil
	}
	return DB.Create(&user).Error
}
