package utils

import "time"
import "github.com/beevik/ntp"

const (
	GOTIME = "2006-01-02 15:04:05"
)

func TimeHMS() string {
	//2020-08-25 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(GOTIME)
	return timeStr
}

func TimeUNix() int64 {
	// 时间戳
	times := NewNtp()
	//return time.Now().Unix()
	return times.Unix()
}

func NewNtp() time.Time {
	response, err := ntp.Query("ntp.aliyun.com")
	if err != nil {
		return time.Time{}
	}
	time := time.Now().Add(response.ClockOffset)
	return time
}
