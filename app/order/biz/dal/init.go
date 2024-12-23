package dal

import (
	"gomall/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
