package repository

import (
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/config/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	OauthDB *gorm.DB

	err error
)

func init() {
	//get configuration
	cfg := config.GetConfig()

	/*initiate connection to all repositories*/

	//connect to endpoint db
	if OauthDB, err = connectMysql(cfg.Database.Oauth); err != nil {
		panic(err)
	}

	if cfg.Server.Debug == "false" { //TO DO: set log mode
		OauthDB.LogMode(false)
	}
}

func connectMysql(db database.Database) (*gorm.DB, error) {
	return gorm.Open(db.Adapter, db.Username+":"+db.Password+"@("+db.Address+":"+db.Port+")/"+db.Database+"?charset=utf8&parseTime=True&loc=Local")
}
