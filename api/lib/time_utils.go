package lib

import "time"

type (
	TimeNowFunc func() time.Time

	TimeUtils interface {
		Now() time.Time
	}

	defaultTimeUtils struct {
		TimeUtils
	}

	NewTimeUtilsFunc func() TimeUtils
)

var TimeNow TimeNowFunc = time.Now

var NewTimeUtils NewTimeUtilsFunc = func() TimeUtils {
	return &defaultTimeUtils{}
}

func (defaultTimeUtils) Now() time.Time {
	return time.Now()
}
