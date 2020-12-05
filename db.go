package main

import (
	"fmt"
	"os"

	"github.com/CTFBox/CTFBox/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	MYSQL_HOSTNAME = os.Getenv("MYSQL_HOSTNAME")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
	MYSQL_USERNAME = os.Getenv("MYSQL_USERNAME")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
)

func SetupDatabase() (*gorm.DB, error) {
	var err error
	//tmp
	if MYSQL_HOSTNAME == "" {
		MYSQL_HOSTNAME = "db"
	}
	if MYSQL_DATABASE == "" {
		MYSQL_DATABASE = "ctf_box"
	}
	if MYSQL_USERNAME == "" {
		MYSQL_USERNAME = "root"
	}

	if MYSQL_PASSWORD == "" {
		MYSQL_PASSWORD = "password"
	}

	// データベース接続
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOSTNAME, MYSQL_DATABASE))
	if err != nil {
		return nil, err
	}
	if err := initDB(db); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(1024) // デフォルトだと2
	db.DB().SetConnMaxLifetime(0) // 一応セット
	db.DB().SetConnMaxIdleTime(0) // 一応セット go1.15以上

	return db, nil
}

func initDB(db *gorm.DB) error {
	db.AutoMigrate(&(repository.Problem{}))
	return nil
}
