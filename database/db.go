package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/khasyah-fr/phinco-attendance/config"
)

func InitializeDB(dbc *config.DbConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.Name)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
