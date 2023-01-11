package utils

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// SharedEnvironment ...
type SharedEnvironment struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
	Cfg *ini.File
}

type Config struct {
	*ini.File
}

// Env ...
var Env *SharedEnvironment

func init() {
	if Env == nil {
		Env = new(SharedEnvironment)
	}
}

func InitDB(cfg *ini.File) *gorm.DB {

	user := cfg.Section("").Key("user").String()
	dbname := cfg.Section("").Key("dbname").String()
	pass := cfg.Section("").Key("password").String()
	dbHost := cfg.Section("").Key("dbhost").String()

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, user, dbname, pass))

	db.LogMode(true)

	if err != nil {
		log.Println("Could not connect to the database:", err)
	}

	return db
}
