package out

import (
	"encoding/json"
	"fmt"
	"github.com/greycodee/gofund/api"
	"github.com/greycodee/gofund/base"
	"time"
)

// [{"name": "mpd", "instance": "now playing", "full_text": "'${status}' '$1'", "color": "'${color}'"}]
type I3status struct {
	Name string `json:"name"`
	Instance string `json:"instance"`
	FullText string `json:"full_text"`
	Color string `json:"color"`
}

func (i3 I3status) Print(factor Factor)  {
	result:= i3.convert(factor)
	var temp []I3status
	for  {
		// 判断时间是否在交易时间内
		if base.IsTransDay() {
			result= i3.convert(factor)
			temp=result
		}else {
			temp=append(result,i3.addTips())
		}
		j,_:=json.Marshal(temp)
		fmt.Println(string(j))
		time.Sleep(time.Duration(factor.Interval))
	}
}

/*
	结果转换
*/
func (i3 I3status) convert(factor Factor) []I3status {
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
func (i3 I3status) addTips() I3status {
	i:= I3status{
		Name:     "休市",
		Instance: "close",
		Color:    PINK,
		FullText: fmt.Sprintf("【休市中:%s】",time.Now().Format("2006-01-02 15:04:05")),
	}
	return i
}

var (
	RED="#b72e2e"
	GREEN="#50b72e"
	PINK="#ffe6f9"
)