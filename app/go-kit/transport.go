package go_kit

import (
	"context"
	"encoding/json"
	"fmt"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

func makeHTTPHandler(e *echo.Echo, s Service) {
	ends := MakeServerEndpoints(s)
	options := []kitHttp.ServerOption{
		//kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(kitHttp.DefaultErrorEncoder),
	}

	//rateBucket := rate.NewLimiter(rate.Every(time.Second*1), 3)
	//echoH := newTokenBucketLimitWithBuildIn(rateBucket)(echoEndpoint)

	e.GET("/echo", echo.WrapHandler(kitHttp.NewServer(
		ends.EchoEndpoint,
		decodeEchoRequest,
		encodeResponse,
		options...,
	)))
}

func decodeEchoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	a := ctx.Value("app")
	fmt.Println(a)

	// Decode incoming request with JSON body as an echoRequest
	var request echoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	// Return decoded request
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
