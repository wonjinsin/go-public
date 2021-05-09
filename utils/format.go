package utils

import "time"

func TimeFormat(t time.Time) string {
	formattedTime := t.Format("2006-01-02 3:4:5")
	return formattedTime
}
