package out

import (
	"encoding/json"
	"fmt"
	"github.com/greycodee/gofund/api"
	"github.com/greycodee/gofund/base"
	"time"
)

// [{"name": "mpd", "instance": "now playing", "full_text": " '${status}' '$1'", "color": "'${color}'"}]
type I3status struct {
	Name string `json:"name"`
	Instance string `json:"instance"`
	FullText string `json:"full_text"`
	Color string `json:"color"`
}

func (i3 I3status) Print(factor Factor)  {
	result:= convert(factor)
	firstAdd:=true
	i:=factor.Interval *1e9
	for  {
		// 判断时间是否在交易时间内
		if base.IsTransDay() {
			result= convert(factor)
			// 重置休市提示
			firstAdd=true
		}else {
			// 第一次添加休市提示
			if firstAdd {
				result=append(result, addTips())
				firstAdd=false
			}
		}
		j,_:=json.Marshal(result)
		fmt.Println(string(j))
		time.Sleep(time.Duration(i))
	}
}

/*
	结果转换
*/
func convert(factor Factor) []I3status {
	var result []I3status
	f:= fundDetail(factor)
	for _,v:=range f {
		var i I3status
		i.Name=v.Name
		i.Instance=v.Code
		i.Color= RED
		if v.Trend==api.DOWN {
			i.Color= GREEN
		}
		i.FullText=v.Name+":"+v.Trend+v.Valuation
		result=append(result,i)
	}
	return result
}

/*
	添加休市文本
*/
func addTips() I3status {
	i3:= I3status{
		Name:     "休市",
		Instance: "close",
		Color:    PINK,
		FullText: "【休市中】",
	}
	return i3
}

const(
	RED="#b72e2e"
	GREEN="#50b72e"
	PINK="#ffe6f9"
)