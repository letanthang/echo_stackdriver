package migrate

import (
	"fmt"
	"sync"
	"time"

	logger "g.ghn.vn/go-common/zap-logger"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

var log = logger.GetLogger("DB - Customer")
var SessionOrginal *mgo.Session
var databaseName string

var (
	MongoDBHosts string
	AuthDatabase string
	AuthUserName string
	AuthPassword string
)

var (
	once sync.Once
	db1  *gorm.DB
)

//Init mongo client
func init() {
	MongoDBHosts = "35.240.214.111"
	AuthDatabase = "admin"
	AuthUserName = "truck_staging_dba"
	AuthPassword = "ZjRiZTg5ZTRmM2MxNzQyY2U2NmM1ZjdjOT"
	databaseName = "customer"
	nativeConnection()

	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}
	db, _ := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		"35.240.214.111", "postgres", "thangtest", "gtsa4934"))

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	// db.DB().SetConnMaxLifetime(time.Second * 10)
	db.DB().Ping()
	db.LogMode(true)
	db1 = db
	fmt.Println("init connected db")

}

func nativeConnection() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mgoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	log.Info("Mongodb connected")
	mgoSession.SetMode(mgo.Monotonic, true)
	SessionOrginal = mgoSession
}

func new() *gorm.DB {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}
	db, _ := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		"35.240.214.111", "postgres", "thangtest", "gtsa4934"))
	// db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	db.LogMode(true)
	return db
}

func GetDB() *gorm.DB {
	once.Do(func() {
		db1 = new()
	})
	return db1
}
