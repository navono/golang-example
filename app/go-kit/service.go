package go_kit

import "fmt"

type (
	Service interface {
		Echo(msg string) string
	}
	LocalService struct {
	}

	ServiceMiddleware func(Service) Service
)

func (l LocalService) Echo(msg string) string {
	return fmt.Sprintf("return \"%s\" from server", msg)
}
