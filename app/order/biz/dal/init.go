package dal

import (
	"github.com/LXJ0000/gomall/app/order/biz/dal/mysql"
	"github.com/LXJ0000/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
