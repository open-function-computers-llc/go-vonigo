package vonigo

import (
	"strconv"
	"time"
)

func getMinutesAgo(minutes string) (string, string) {

	i, _ := strconv.Atoi(minutes)
	endTime := makeTimestampMilli()
	startTime := endTime - int64(i*60*1000)

	// convert int64 to strings
	endTimeString := strconv.FormatInt(endTime, 10)
	startTimeString := strconv.FormatInt(startTime, 10)
	return endTimeString, startTimeString
}

func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func makeTimestampMilli() int64 {
	return unixMilli(time.Now())
}
