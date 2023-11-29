package mygorm

import (
	"time"

	"github.com/mrbryside/rbh/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Close func()

func BasicConnection() (*gorm.DB, Close, error) {
	db, err := gorm.Open(mysql.Open(env.Data().DbUrl()), &gorm.Config{})
	if err != nil {
		return nil, func() {}, err
	}
	close := func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}
	return db, close, nil
}

func BasicMainConnection() (*gorm.DB, Close, error) {
	var close func()
	db := &gorm.DB{}
	var err error
	for {
		db, err = gorm.Open(mysql.Open(env.Data().DbUrl()), &gorm.Config{})
		close = func() {
			dbInstance, _ := db.DB()
			_ = dbInstance.Close()
		}
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	return db, close, nil
}
