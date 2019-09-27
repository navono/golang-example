package go_kit

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
)

type (
	echoRequest struct {
		Input string `json:"input"`
	}

	echoResponse struct {
		Output string `json:"output"`
	}

	Endpoints struct {
		EchoEndpoint endpoint.Endpoint
	}
)

func MakeServerEndpoints(s Service) Endpoints {
	var echoEndpoint endpoint.Endpoint
	{
		echoEndpoint = makeEchoEndpoint(s)
		// rate limit
		echoEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(echoEndpoint)
		// circuit breaker
		commandName := "echo"
		hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
			Timeout:                1000 * 30,
			MaxConcurrentRequests:  1000,
			RequestVolumeThreshold: 5,
			SleepWindow:            10000,
			ErrorPercentThreshold:  1,
		})
		echoEndpoint = circuitbreaker.Hystrix(commandName)(echoEndpoint)
	}

	return Endpoints{
		EchoEndpoint: echoEndpoint,
	}
}

func makeEchoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// Assert type of incoming request
		a := ctx.Value("app")
		fmt.Println(a)
		req := request.(echoRequest)

		e := s.Echo(req.Input)

		// Return response
		return echoResponse{Output: e}, nil
	}
}
