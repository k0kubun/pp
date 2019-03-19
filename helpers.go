package pp

import "time"

func StrToTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, str)
	return t
}
