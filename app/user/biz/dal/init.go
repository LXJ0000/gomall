package dal

import (
	"github.com/LXJ0000/gomall/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
