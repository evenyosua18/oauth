package repository

import (
	"github.com/evenyosua18/oauth/config"
	"github.com/evenyosua18/oauth/config/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
}

func connectMysql(db database.Database) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(db.Username + ":" + db.Password + "@(" + db.Address + ":" + db.Port + ")/" + db.Database + "?charset=utf8&parseTime=True&loc=Local"))
}
