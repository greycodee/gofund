package main

import (
	"flag"
	"github.com/greycodee/gofund/api"
)

var confPath = flag.String("config","$HOME/.i3/gofund.conf","基金代码配置文件")
var interval = flag.Int64("interval",5,"刷新间隔(秒)")
var o = flag.Int64("o",1,"输出格式:\n1    i3status格式\n2    终端命令行格式")
var fundServer = flag.Int64("api",1,"api服务选择:\n1    天天基金api")

var codes []string
func init() {
	flag.Parse()
	codes=readConfig()
}

func main()  {
	out:=chosePrint()
	out.Print()
}

func chosePrint() PrintFormat{
	switch *o {
	case 1:
		return i3status{}
	case 2:
		return i3status{}
	default:
		panic("所选输出格式不存在")
	}
}

func choseApi() api.Fund {
	switch *fundServer {
		case 1:
			return api.TTFund{}
		default:
			panic("选择api服务不支持")
	}
}

/*
	循环请求所有基金，拼接结果
*/
func fundDetail() []api.GoFund {
	var result []api.GoFund
	// 选择基金接口
	server:=choseApi()
	for _,code:= range codes{
		f,t:=server.Api(code)
		if t==nil {
			result=append(result,f)
		}
	}
	return result
}

type PrintFormat interface {
	Print()
}

