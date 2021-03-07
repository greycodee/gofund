package out

import "github.com/greycodee/gofund/api"

type PrintFormat interface {
	Print(factor Factor)
}

/*
	循环请求所有基金，拼接结果
*/
func fundDetail(factor Factor) []api.GoFund {
	var result []api.GoFund
	// 选择基金接口
	server:=factor.Server
	for _,code:= range factor.Codes{
		f,t:=server.Api(code)
		if t==nil {
			result=append(result,f)
		}
	}
	return result
}
