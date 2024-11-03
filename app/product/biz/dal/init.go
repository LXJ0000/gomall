package dal

import (
	"github.com/LXJ0000/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
