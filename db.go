package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	MARIADB_HOSTNAME = os.Getenv("MARIADB_HOSTNAME")
	MARIADB_DATABASE = os.Getenv("MARIADB_DATABASE")
	MARIADB_USERNAME = os.Getenv("MARIADB_USERNAME")
	MARIADB_PASSWORD = os.Getenv("MARIADB_PASSWORD")
)

func SetupDatabase() (*gorm.DB, error) {
	var err error
	//tmp
	if MARIADB_HOSTNAME == "" {
		MARIADB_HOSTNAME = "localhost"
	}
	if MARIADB_DATABASE == "" {
		MARIADB_DATABASE = "ctf_box"
	}
	if MARIADB_USERNAME == "" {
		MARIADB_USERNAME = "root"
	}

	if MARIADB_PASSWORD == "" {
		MARIADB_PASSWORD = "password"
	}

	// データベース接続
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", MARIADB_USERNAME, MARIADB_PASSWORD, MARIADB_HOSTNAME, MARIADB_DATABASE))
	if err != nil {
		return nil, err
	}
	if err := initDB(db); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(1024) // デフォルトだと2
	db.DB().SetConnMaxLifetime(0)      // 一応セット
	db.DB().SetConnMaxIdleTime(0)      // 一応セット go1.15以上

	return db, nil
}

func initDB(db *gorm.DB) error {
	return nil
}
