package go_kit

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type (
	instrumentMw struct {
		requestCount   metrics.Counter
		requestLatency metrics.Histogram
		countResult    metrics.Histogram
		Service
	}
)

func instrumentMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
) ServiceMiddleware {
	return func(next Service) Service {
		return instrumentMw{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			countResult:    countResult,
			Service:        next,
		}
	}
}

func (mw instrumentMw) Echo(msg string) string {
	defer func(begin time.Time) {
		//lvs := []string{"method", "echo", "error", fmt.Sprint(err != nil)}
		lvs := []string{"method", "echo", "error"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.Service.Echo(msg)
}
