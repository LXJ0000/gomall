package dal

import (
	"github.com/LXJ0000/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/LXJ0000/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
