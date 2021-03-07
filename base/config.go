package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	读取配置返回配置参数
*/
func ReadConfig(path string) []string {
	var fundCodes []string

	conf,err:=os.Open(os.ExpandEnv(path))
	if err!=nil {
		fmt.Println(err)
		fmt.Printf("读取配置出错,请确认配置是否存在及格式是否正确：%s",path)
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
