package mw

import (
	"Another-Nikki/dal"
	"time"
)

func GetVisitTime(userID int64) int64 {
	user, err := dal.User{}.GetUserById(userID)
	if err != nil {
		return -1
	}
	return user.VisitTime
}

func SetVisitTime(userID, visitTime int64) {
	dal.User{}.UpdateUser(userID, map[string]interface{}{
		"visit_time": visitTime,
	})
}

func AddVisitTime(userID, val int) {
	visitTime := GetVisitTime(int64(userID))
	sum := visitTime >> 32
	today := visitTime & (1<<32 - 1)
	sum += int64(val)
	today += int64(val)
	visitTime = sum<<32 + today
	SetVisitTime(int64(userID), visitTime)
}

func SetTodayTime(userID int64) {
	t := time.NewTicker(time.Hour)
	for {
		<-t.C
		if time.Now().Hour() != 0 {
			continue
		}
		vistTime := GetVisitTime(userID)
		vistTime = vistTime >> 32 << 32
		SetVisitTime(userID, vistTime)
	}
}
