//go:generate mockgen -destination=mocks/service.go -source=./service.go Service
package testexample

import "context"
import _ "github.com/golang/mock/mockgen/model"

type ServiceArgs struct {
	Param1 string
	Param2 string
}

type Service interface {
	Process(context.Context, ServiceArgs) (string, error)
}

type ServiceImpl struct {
	TestMode bool
}

func (s ServiceImpl) Process(ctx context.Context, args ServiceArgs) (string, error) {
	return "", nil
}
