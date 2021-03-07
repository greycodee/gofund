package out

import (
	"fmt"
	"github.com/greycodee/gofund/api"
	"github.com/greycodee/gofund/base"
	"strings"
	"time"
)

type Terminal struct {

}

func (t Terminal) Print(factor Factor)  {
	result:=t.convert(factor)
	var temp []string
	for  {
		if base.IsTransDay() {
			result=t.convert(factor)
			temp=result
		}else{
			temp=append(result,fmt.Sprintf("【休市中:%s】",time.Now().Format("2006-01-02 15:04:05")))
		}
		rs:=strings.Join(temp,"|")
		fmt.Printf("%s\r",rs)
		time.Sleep(time.Duration(factor.Interval))
	}
}

func (t Terminal) convert(factor Factor) []string  {
	funds:=fundDetail(factor)
	var result []string
	for _,v:=range funds{
		color:=31
		if v.Trend==api.DOWN {
			color=32
		}
		s:=fmt.Sprintf("\033[1;%dm%s:%s%s\033[0m",color,v.Name,v.Trend,v.Valuation)
		result=append(result,s)
	}
	return result
}