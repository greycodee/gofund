package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

var confPath = flag.String("config","$HOME/.i3/gofund.conf","基金代码配置文件")
func main()  {
	urls:=requestUrls()
	var result []i3status
	for _,v:= range urls{
		i3,t:=requestFund(v)
		if t {
			result=append(result,i3)
		}
	}
	j,_:=json.Marshal(result)
	fmt.Println(string(j))
}


const(
	RED="#b72e2e"
	GREEN="#50b72e"
)

