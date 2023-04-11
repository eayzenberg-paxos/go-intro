package testexample

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type Controller struct {
	service Service
	logger  *zap.Logger
}

func NewController(service Service, logger *zap.Logger) Controller {
	return Controller{
		service: service,
		logger:  logger,
	}
}

func (c *Controller) Process(ctx context.Context) string {
	result, err := c.service.Process(ctx, ServiceArgs{})
	if err != nil {
		c.logger.Error(fmt.Sprintf("Process error: %s", err.Error()))
		return ""
	}
	return result
}
