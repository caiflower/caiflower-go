package limiter

import (
	"time"

	"golang.org/x/time/rate"
)

const (
	Fail    = "fail"
	Success = "success"
)

var limiter *rate.Limiter

func init() {
	// b: 桶的大小 r: 每秒允许的
	limiter = rate.NewLimiter(10, 10000)
}

func Run() string {
	if limiter.AllowN(time.Now(), 1) {
		return Success
	} else {
		return Fail
	}
}
