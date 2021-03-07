package out

import "github.com/greycodee/gofund/api"

/*
	输出传参
*/
type Factor struct {
	Server 		api.Fund		// API
	Codes 		[]string		// 基金代码列表
	Interval 	int64			// 间隔输出时间(秒)
}
