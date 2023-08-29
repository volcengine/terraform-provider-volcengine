package common

import (
	"fmt"

	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
)

var rateInfoMap map[string]*Rate

func init() {
	rateInfoMap = map[string]*Rate{
		"ecs.RunInstance.2020-04-01": {
			Limiter:   rate.NewLimiter(4, 10),
			Semaphore: semaphore.NewWeighted(14),
		},
		"ecs.DescribeInstance.2020-04-01": {
			Limiter:   rate.NewLimiter(4, 10),
			Semaphore: semaphore.NewWeighted(14),
		},
	}
}

type Rate struct {
	Limiter   *rate.Limiter
	Semaphore *semaphore.Weighted
}

func GetRateInfoMap(svc, action, version string) *Rate {
	key := fmt.Sprintf("%s.%s.%s", svc, action, version)
	return rateInfoMap[key]
}
