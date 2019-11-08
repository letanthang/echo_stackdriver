package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/letanthang/echo_stackdriver/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	once sync.Once
	db1  *gorm.DB
)

func init() {
	// GetDB()
}

func new() *gorm.DB {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Config.Db.Host, config.Config.Db.User, config.Config.Db.Name, config.Config.Db.Password))
	if err != nil {

		log.Printf("Connect not success to postgres database at host:%s with user:%s and db:%s",
			config.Config.Db.Host, config.Config.Db.User, config.Config.Db.Name)
	}
	// db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// db.DB().SetConnMaxLifetime(time.Nanosecond)
	// db.DB().SetConnMaxLifetime(3 * time.Second)
	db.DB().Ping()
	// db.DB().Exec("SET timezone TO 'Asia/Ho_Chi_Minh';")
	if config.Config.Db.Debug {
		db.LogMode(true)
	}
	return db
}

func GetDB() *gorm.DB {
	once.Do(func() {
		db1 = new()
	})
	return db1
}
