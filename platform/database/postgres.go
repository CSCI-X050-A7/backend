package database

import (
	"fmt"

	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDatabase(gormConfig *gorm.Config) error {
	conf := config.Conf
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUsername,
		conf.DBPassword,
		"postgres",
	)
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return err
	}
	// check if db exists
	stmt := fmt.Sprintf(
		"SELECT * FROM pg_database WHERE datname = '%s';",
		conf.DBName,
	)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return rs.Error
	}
	// if not create it
	rec := make(map[string]any)
	rs.Find(rec)
	if len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", conf.DBName)
		if rs := db.Exec(stmt); rs.Error != nil {
			return rs.Error
		}
		// close db connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func connectDatabase(gormConfig *gorm.Config) (err error) {
	conf := config.Conf
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUsername,
		conf.DBPassword,
		conf.DBName,
	)
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return
	}
	err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	return
}

func ConnectPostgres() {
	gormConfig := GetGormConfig()
	if err := createDatabase(&gormConfig); err != nil {
		logrus.Fatalf("failed to create db in Postgres: %+v", err)
	}
	if err := connectDatabase(&gormConfig); err != nil {
		logrus.Fatalf("failed to connect to Postgres: %+v", err)
	}
	if err := MigrateDatabase(); err != nil {
		logrus.Fatalf("failed to migrate in Postgres: %+v", err)
	}
}
