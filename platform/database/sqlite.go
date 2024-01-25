package database

import (
	"database/sql"

	"github.com/google/uuid"
	sqliteGo "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite() {
	gormConfig := GetGormConfig()
	const CustomDriverName = "sqlite3_extended"
	const File = "./test.db"
	sql.Register(CustomDriverName,
		&sqliteGo.SQLiteDriver{
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
		},
	)
	conn, err := sql.Open(CustomDriverName, File)
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(sqlite.Dialector{
		DriverName: CustomDriverName,
		DSN:        File,
		Conn:       conn,
	}, &gormConfig)
	if err != nil {
		logrus.Fatalf("failed database setup. error: %v", err)
	}
	if err := MigrateDatabase(); err != nil {
		logrus.Fatalf("failed to migrate in Postgres: %+v", err)
	}
	if err := CreateDemoUser(); err != nil {
		logrus.Infof("failed to create demo user: %+v", err)
	}
}
