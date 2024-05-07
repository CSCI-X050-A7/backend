package database

import (
	"database/sql"
	"database/sql/driver"

	"github.com/CSCI-X050-A7/backend/pkg/config"
	"github.com/google/uuid"
	sqliteGo "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// adapter design pattern
func SqliteUUIDAdapter() (string, driver.Driver) {
	return "sqlite3_extended", &sqliteGo.SQLiteDriver{
		ConnectHook: func(conn *sqliteGo.SQLiteConn) error {
			err := conn.RegisterFunc(
				"uuid_generate_v4",
				func(arguments ...interface{}) (string, error) {
					return uuid.NewString(), nil // Return a string value.
				},
				true,
			)
			return err
		},
	}
}

func ConnectSqlite() {
	gormConfig := GetGormConfig()
	name, adapter := SqliteUUIDAdapter()
	sql.Register(name, adapter)
	conn, err := sql.Open(name, config.Conf.DBFilename)
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(sqlite.Dialector{
		DriverName: name,
		DSN:        config.Conf.DBFilename,
		Conn:       conn,
	}, &gormConfig)
	if err != nil {
		logrus.Fatalf("failed database setup. error: %v", err)
	}
	if err := MigrateDatabase(); err != nil {
		logrus.Fatalf("failed to migrate in Sqlite: %+v", err)
	}
	if err := CreateDemoUser(); err != nil {
		logrus.Infof("failed to create demo user: %+v", err)
	}
}
