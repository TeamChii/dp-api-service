package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	log "github.com/sirupsen/logrus"
)

// "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/postgres"
// log "github.com/sirupsen/logrus"

var db *gorm.DB

//var session *mgo.Session

func InitPostgres() {

	var err error
	db, err = gorm.Open("postgres", "host=35.240.143.157 port=15432 user=dpoint dbname=dp_db password=dpointpass sslmode=disable")
	//db, err = gorm.Open("postgres", "host=172.17.0.3 port=5432 user=dpoint dbname=dp_db password=dpointpass sslmode=disable")
	//db.LogMode(true)
	if err != nil {
		log.Error(err.Error())
		panic("failed to connect database")
	}
}

func DBsql() *gorm.DB {

	return db
}

/*
func DBmongo() *mgo.Session {
	var err error
	session, err := mgo.Dial("35.198.242.63:27008")
	if err != nil {
		panic("failed to connect database")
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}
*/
