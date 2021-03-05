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

func main()  {

	fundDetail()

	s:="sdd"
	fmt.Println(s[1:])

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
		fmt.Printf("\033[1;31;40m↑%.2f\033[0m\n",gszzl)
	}else {
		fmt.Printf("\033[1;32;40m↓%.2f\033[0m\n",gszzl)
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
