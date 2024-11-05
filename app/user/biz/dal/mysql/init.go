package mysql

import (
	"fmt"
	"os"

	"github.com/LXJ0000/gomall/app/user/model"
	"github.com/LXJ0000/gomall/app/user/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
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
	utils.MustHandleError(err)

	err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics()))
	utils.MustHandleError(err)

	if os.Getenv("GO_ENV") != "online" {
		err := DB.AutoMigrate(
			&model.User{},
		)
		utils.MustHandleError(err)
	}
}
