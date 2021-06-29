package utils

import "time"

func TimeFormat(t time.Time) string {
	formattedTime := t.Format("2006-01-02 3:4:5")
	return formattedTime
}

func GetNow() time.Time {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Seoul")

	//set timezone,
	now := time.Now().In(loc)

	return now
}

func ErrorToStr(err error) string {
	return err.Error()
}
