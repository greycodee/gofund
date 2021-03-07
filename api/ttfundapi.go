package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
	天天基金接口返回
*/
type TTFund struct {
	Fundcode string `json:"fundcode"`
	Name string `json:"name"`
	Jzrq string `json:"jzrq"`
	Dwjz string `json:"dwjz"`
	Gsz string `json:"gsz"`
	Gszzl string `json:"gszzl"`
	Gztime string `json:"gztime"`
}

/*
	天天基金接口
*/
func (t TTFund) Api(code string) (GoFund,error) {
	var goFund GoFund
	var ttFund TTFund
	ttFundUrl:="http://fundgz.1234567.com.cn/js/"+code+".js"
	param:=url.Values{}
	URL,e:=url.Parse(ttFundUrl)
	if e!=nil {
		return goFund,e
	}
	param.Set("rt",strconv.FormatInt(time.Now().Unix(),10))
	URL.RawQuery = param.Encode()
	reqUrl:=URL.String()
	resp, err := http.Get(reqUrl)
	if err != nil {
		return goFund,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return goFund,http.ErrNotSupported
	}
	body, err := ioutil.ReadAll(resp.Body)
	bodyString:=string(body)
	// 删除尾部2个字符
	bodyString=bodyString[:len(bodyString)-2]
	// 删除头部8个字符
	bodyString=bodyString[8:]

	_=json.Unmarshal([]byte(bodyString),&ttFund)
	goFund.Name=ttFund.Name
	goFund.Code=ttFund.Fundcode
	goFund.UpdateTime=ttFund.Gztime
	goFund.Value=ttFund.Dwjz
	goFund.NowValue=ttFund.Gsz
	goFund.Valuation=ttFund.Gszzl
	goFund.Trend=UP
	gszzl,_ := strconv.ParseFloat(ttFund.Gszzl,64)
	if gszzl < 0 {
		goFund.Trend=DOWN
	}
	return goFund,nil
}