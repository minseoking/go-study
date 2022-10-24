package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Type string

type RDB struct {
	DBType  Type   `json:"dbType" bson:"dbType"`
	Host    string `json:"host" bson:"host"`
	Port    string `json:"port" bson:"port"`
	AdminId string `json:"adminId" bson:"adminId"`
	AdminPw string `json:"adminPw" bson:"adminPw"`
	DbName  string `json:"dbName" bson:"dbName"`
}

const (
	SQLSERVER  = Type("sqlserver")
	MYSQL      = Type("mysql")
	POSTGRESQL = Type("postgresql")
)

var RDb *gorm.DB

func (r RDB) ConnectRDB() {
	var dial gorm.Dialector
	switch r.DBType {
	case SQLSERVER:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			r.AdminId,
			r.AdminPw,
			r.Host,
			r.Port,
			r.DbName)
		dial = sqlserver.Open(dsn)
	case MYSQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			r.AdminId,
			r.AdminPw,
			r.Host,
			r.Port,
			r.DbName)
		dial = mysql.Open(dsn)
	case POSTGRESQL:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
			r.Host,
			r.AdminId,
			r.AdminPw,
			r.DbName,
			r.Port)
		dial = postgres.Open(dsn)
	}

	db, err := gorm.Open(dial, &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	RDb = db
}
