package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"
)

var confPath = flag.String("config","$HOME/.i3/gofund.conf","基金代码配置文件")
var interval = flag.Int64("interval",5,"刷新间隔(秒)")

var urls []string
func init() {
	flag.Parse()
	urls=requestUrls()
}

func main()  {
	result:=fundDetail()
	firstAdd:=true
	i:=*interval*1e9

	for  {
		// 判断时间是否在交易时间内
		if isTransDay() {
			result=fundDetail()
			// 重置休市提示
			firstAdd=true
		}else {
			// 第一次添加休市提示
			if firstAdd {
				result=append(result,addTips())
				firstAdd=false
			}
		}
		j,_:=json.Marshal(result)
		fmt.Println(string(j))
		time.Sleep(time.Duration(i))
	}

}

/*
	循环请求所有基金url，拼接结果
*/
func fundDetail() []i3status {
	var result []i3status
	for _,v:= range urls{
		i3,t:=requestFund(v)
		if t {
			result=append(result,i3)
		}
	}
	return result
}

/*
	添加休市文本
*/
func addTips() i3status{
	i3:=i3status{
		Name: "休市",
		Instance: "close",
		Color: PINK,
		FullText: "【休市中】",
	}
	return i3
}


const(
	RED="#b72e2e"
	GREEN="#50b72e"
	PINK="#ffe6f9"
)

