package migration

import (
	"github.com/evenyosua18/oauth/app/repository/oauth_db/model"
	"github.com/evenyosua18/oauth/config"
	dbUtil "github.com/evenyosua18/oauth/util/db"
	"github.com/jinzhu/gorm"
)

func InitDatabase() (err error) {
	//get configuration
	cfg := config.GetConfig()

	//open connection
	var db *gorm.DB
	db, err = dbUtil.ConnectMysql(cfg.Database.Oauth)

	if err != nil {
		return
	}

	db.AutoMigrate(&model.Endpoint{}, &model.Scope{}, &model.Role{}, &model.User{}, &model.OauthClient{}, &model.AuthorizationCode{}, &model.AccessToken{}, &model.RefreshToken{})
	return
}
