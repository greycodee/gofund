package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	读取配置返回配置参数
*/
func readConfig() []string {
	var fundCodes []string

	conf,err:=os.Open(os.ExpandEnv(*confPath))
	if err!=nil {
		fmt.Println(err)
		fmt.Printf("读取配置出错,请确认配置是否存在及格式是否正确：%s",*confPath)
		os.Exit(1)
	}
	defer conf.Close()
	b:=bufio.NewReader(conf)
	for  {
		code,_,e:=b.ReadLine()
		if e == io.EOF {
			break
		}
		fundCodes=append(fundCodes,string(code))
	}
	return fundCodes
}


/*
	拼接天天基金url
*/
func requestUrls() []string{
	codes:=readConfig()
	TTFundUrl:="http://fundgz.1234567.com.cn/js/"
	var result []string
	for _,c := range codes{
		url:=TTFundUrl+c+".js"
		result=append(result,url)
	}
	return result
}
