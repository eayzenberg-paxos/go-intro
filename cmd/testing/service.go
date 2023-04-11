//go:generate mockgen -destination=mocks/service.go -package=testing --build_flags=--mod=mod github.com/eayzenberg-paxos/go-intro/cmd/testing Service
package testing

import "context"
import _ "github.com/golang/mock/mockgen/model"

type ServiceArgs struct {
	Param1 string
	Param2 string
}

type Service interface {
	Process(context.Context, ServiceArgs) (string, error)
}
