package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func requestFund(u string) (i3status,bool)  {
	var res fund
	var i3 i3status

	param:=url.Values{}
	URL,e:=url.Parse(u)
	if e!=nil {
		return i3,false
	}
	param.Set("rt",strconv.FormatInt(time.Now().Unix(),10))
	URL.RawQuery = param.Encode()
	reqUrl:=URL.String()
	resp, err := http.Get(reqUrl)
	if err != nil {
		return i3,false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return i3,false
	}
	body, err := ioutil.ReadAll(resp.Body)
	bodyString:=string(body)
	// 删除尾部2个字符
	bodyString=bodyString[:len(bodyString)-2]
	// 删除头部8个字符
	bodyString=bodyString[8:]

	_=json.Unmarshal([]byte(bodyString),&res)
	gszzl,_ := strconv.ParseFloat(res.Gszzl,64)
	i3.Instance=res.Fundcode
	i3.Name=res.Name
	i3.Color=GREEN
	j:="↓"
	if gszzl >=0 {
		j="↑"
		i3.Color=RED
	}
	i3.FullText=res.Name+":"+j+res.Gszzl
	return i3,true
}
