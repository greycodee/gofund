package main

import (
	"flag"
	"github.com/greycodee/gofund/api"
	"github.com/greycodee/gofund/base"
	"github.com/greycodee/gofund/out"
)

var confPath = flag.String("config","$HOME/.fund/gofund.conf","基金代码配置文件")
var interval = flag.Int64("interval",5,"刷新间隔(秒)")
var o = flag.Int64("o",1,"输出格式:\n1    i3status格式\n2    终端命令行格式")
var fundServer = flag.Int64("api",1,"api服务选择:\n1    天天基金api")

var codes []string
func init() {
	flag.Parse()
	codes= base.ReadConfig(*confPath)
}

func main()  {
	op:=chosePrint()
	factor:=out.Factor{
		Codes: codes,
		Server: choseApi(),
		Interval: *interval*1e9,
	}
	op.Print(factor)
}

func chosePrint() out.PrintFormat{
	switch *o {
	case 1:
		return out.I3status{}
	case 2:
		return out.Terminal{}
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

