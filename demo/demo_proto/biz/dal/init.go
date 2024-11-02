package dal

import (
	"github.com/LXJ0000/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/LXJ0000/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
