package base

import (
	"time"
)

/*
	判断是否是交易日 [后续加上法定节假日]
*/
func IsTransDay() bool {
	t := time.Now()
	if t.Weekday()==time.Weekday(6) || t.Weekday()==time.Weekday(0) {
		return false
	}
	if isHoliday() {
		return false
	}
	return isTransTime()
}
func isTransTime() bool {
	start:=time.Now().Format("2006-01-02")+" 09:30:00"
	end:=time.Now().Format("2006-01-02")+" 11:30:00"
	start2:=time.Now().Format("2006-01-02")+" 13:00:00"
	end2:=time.Now().Format("2006-01-02")+" 15:00:00"
	loc,_:=time.LoadLocation("Local")

	tS,_:=time.ParseInLocation("2006-01-02 15:04:05",start,loc)
	tE,_:=time.ParseInLocation("2006-01-02 15:04:05",end,loc)
	tS2,_:=time.ParseInLocation("2006-01-02 15:04:05",start2,loc)
	tE2,_:=time.ParseInLocation("2006-01-02 15:04:05",end2,loc)

	tN:=time.Now()
	if (tN.Before(tE) && tN.After(tS)) || (tN.Before(tE2) && tN.After(tS2)) {
		return true
	}
	return false
}

/*
	2021法定节假日判断
*/
func isHoliday() bool {
	nowDay:=time.Now().Format("20060102")
	if holidays[nowDay] {
		return false
	}
	return true
}

/*
	在工作日时放假的法定节假日数据，不包含在周末的法定节假日数据
*/
var holidays = map[string]bool{
	"20210101":true,
	"20210211":true,
	"20210212":true,
	"20210215":true,
	"20210216":true,
	"20210217":true,
	"20210405":true,
	"20210503":true,
	"20210504":true,
	"20210505":true,
	"20210614":true,
	"20210920":true,
	"20210921":true,
	"20211001":true,
	"20211004":true,
	"20211005":true,
	"20211006":true,
	"20211007":true,
}