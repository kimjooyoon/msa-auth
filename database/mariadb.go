package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"msa-auth/members"
	"os"
)

type DSN string

func MysqlNewDSN() DSN {
	dsn := os.Getenv("DSN")
	return DSN(dsn)
}

var db *gorm.DB

func MysqlConnection(dataSourceName DSN) *gorm.DB {
	if db != nil {
		return db
	}
	newDb, err := gorm.Open(mysql.Open(string(dataSourceName)), &gorm.Config{})

	if err != nil {
		log.Panicf("new mysql connection err\nerr=%v", err)
		return nil
	}
	db = newDb

	return db
}

func AutoMigrate() {
	fmt.Println("migrate Mysql")
	dsn := MysqlNewDSN()
	db := MysqlConnection(dsn)

	err1 := db.AutoMigrate(&members.Members{})
	if err1 != nil {
		log.Fatalln(err1)
	}

	sqlDB, errDB := db.DB()
	if errDB != nil {
		log.Fatalln(errDB)
	}

	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(sqlDB)
}
