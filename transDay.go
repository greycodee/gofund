package main

import (
	"time"
)

/*
	判断是否是交易日 [后续加上法定节假日]
*/
func isTransDay() bool {
	t := time.Now()
	if t.Weekday()==time.Weekday(6) || t.Weekday()==time.Weekday(0) {
		return false
	}
	return isTransTime()
}
func isTransTime() bool {
	start:=time.Now().Format("2006-01-02")+" 08:55:00"
	end:=time.Now().Format("2006-01-02")+" 15:05:00"
	loc,_:=time.LoadLocation("Local")

	tS,_:=time.ParseInLocation("2006-01-02 15:04:05",start,loc)
	tE,_:=time.ParseInLocation("2006-01-02 15:04:05",end,loc)

	tN:=time.Now()
	if tN.Before(tE) && tN.After(tS) {
		return true
	}
	return false
}