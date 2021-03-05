package main

import (
	"encoding/json"
	"flag"
	"fmt"
)


// {"name":"hh","instance":"test","full_text":"ddd","color":"#b72e2e"}
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

