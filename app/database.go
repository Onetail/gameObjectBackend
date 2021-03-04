package app

import (
	model "gameObjectBackend/app/model"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Database struct {
	app *App
	db  *gorm.DB
}

func (database *Database) Init() {

	dbtype := viper.GetString("database.type")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")

	// Connect to MySQL Database
	log.Printf("Connecting to database server %s:%s/%s ...\n", host, port, dbname)

	db, err := gorm.Open(dbtype, username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True")
	if err != nil {
		host = os.Getenv("MYSQL_HOST")
		dbEnvHost, err := gorm.Open(dbtype, username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True")
		if err != nil {
			log.Fatalf("Failed to connect to database: %v\n", err)
		}
		db = dbEnvHost
	}

	db.LogMode(true)
	database.db = db

	// Initializing models
	db.DropTableIfExists(&model.User{}, &model.UserLogin{}, &model.Attachment{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserLogin{})
	db.AutoMigrate(&model.Attachment{})
}

func (database *Database) GetDb() *gorm.DB {
	return database.db
}