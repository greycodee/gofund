package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type fund struct {
	Fundcode string `json:"fundcode"`
	Name string `json:"name"`
	Jzrq string `json:"jzrq"`
	Dwjz string `json:"dwjz"`
	Gsz string `json:"gsz"`
	Gszzl string `json:"gszzl"`
	Gztime string `json:"gztime"`
}
// [{"name": "mpd", "instance": "now playing", "full_text": " '${status}' '$1'", "color": "'${color}'"}]
type i3status struct {
	Name string `json:"name"`
	Instance string `json:"instance"`
	FullText string `json:"full_text"`
	Color string `json:"color"`
}
// {"name":"hh","instance":"test","full_text":"ddd","color":"#b72e2e"}
func main()  {
	//fmt.Println(RED)
	var i3 i3status
	i3.Name="hh"
	i3.Color=RED
	i3.Instance="test"
	i3.FullText="ddd"
	j,_:=json.Marshal(i3)
	fmt.Println(string(j))
	//fundDetail()
	//
	//s:="sdd"
	//fmt.Println(s[1:])

	//for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
	//	for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
	//		for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
	//			fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
	//		}
	//		fmt.Println("")
	//	}
	//	fmt.Println("")
	//}

}

func fundDetail(){

	url:="http://fundgz.1234567.com.cn/js/003095.js?rt="+strconv.FormatInt(time.Now().Unix(),10)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString:=string(body)
	bodyString=bodyString[:len(bodyString)-2]
	bodyString=bodyString[8:]
	fmt.Println(bodyString)
	fmt.Println(resp.StatusCode)
	var res fund
	_=json.Unmarshal([]byte(bodyString),&res)
	fmt.Println(res.Name)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
		fmt.Printf("\033[1;32;40m%s\033[0m\n","Red.")
	}
	gszzl,_ := strconv.ParseFloat(res.Gszzl,64)
	if gszzl>0 {
		fmt.Printf("\033[1;%d;40m%s%s%.2f\033[0m\u001B[1;37;40m|\u001B[0m",31,res.Name,":↑",gszzl)
		fmt.Printf("\033[1;%d;40m%s%s%.2f\033[0m",32,res.Name,":↓",gszzl)
	}else {
		fmt.Printf("\033[1;%d;40m%s%s%.2f\033[0m\n",32,res.Name,":↓",gszzl)
	}
}

func p()  {
	float,err := strconv.ParseFloat("0.55",64)
	if err!=nil {
		fmt.Println(err)
	}else {
		if float<0 {
			fmt.Println("负数")
		}else {
			fmt.Println("正数")
		}
	}
}

const(
	RED="#b72e2e"
	GREEN="#50b72e"
)

