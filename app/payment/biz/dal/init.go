package dal

import (
	"github.com/LXJ0000/gomall/app/payment/biz/dal/mysql"
	"github.com/LXJ0000/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
