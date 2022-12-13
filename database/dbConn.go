package database

import (
	"thor/api/e-commerce/config"
	"gorm.io/driver/mysql"
	"log"
	"gorm.io/gorm"
)

type Connection struct {
	GormDB *gorm.DB
}

func (conn *Connection) OpenConn() (*gorm.DB, error){
	dbUser := config.DatabaseUser()
	dbPass := config.DatabasePass()
	dbHost := config.DatabaseHost()
	dbPort := config.DatabasePort()

	var err error
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/?charset=utf8mb4&parseTime=True&loc=Local"
	conn.GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return conn.GormDB, nil
}