//go:generate mockgen -destination=mocks/service.go -package=testing github.com/eayzenberg-paxos/go-intro/cmd/testing Service
package testing

import "context"
import _ "github.com/golang/mock/mockgen/model"

type ServiceArgs struct {
	Param1 string
	Param2 string
}

type Service interface {
	Process(context.Context, ServiceArgs) string
}

type ServiceImpl struct {
	dependency1 string
	dependency2 string
}

func (s *ServiceImpl) Process(ctx context.Context, args ServiceArgs) string {
	return ""
}
