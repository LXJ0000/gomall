package mysql

import (
	"fmt"
	"os"

	"github.com/LXJ0000/gomall/demo/demo_proto/biz/model"
	"github.com/LXJ0000/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dns := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)
	DB, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		panic(err)
	}
}
