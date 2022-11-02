package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mitchellh/mapstructure"
)

type database struct {
	Username string
	Password string
	Address  string
	Port     string
	Database string
	Adapter  string
}

func ConnectMysql(in interface{}) (*gorm.DB, error) {
	//decode
	var db database
	if err := mapstructure.Decode(in, &db); err != nil {
		return nil, err
	}

	//connect to db
	return gorm.Open(db.Adapter, db.Username+":"+db.Password+"@("+db.Address+":"+db.Port+")/"+db.Database+"?charset=utf8&parseTime=True&loc=Local")
}
