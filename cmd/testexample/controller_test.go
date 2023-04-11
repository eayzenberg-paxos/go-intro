package testexample_test

import (
	"context"
	"fmt"
	"github.com/eayzenberg-paxos/go-intro/cmd/testexample"
	"github.com/eayzenberg-paxos/go-intro/cmd/testexample/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func TestHandle(t *testing.T) {
	var (
		observedZapCore zapcore.Core
		observedLogs    *observer.ObservedLogs
		observedLogger  *zap.Logger
		ctrl            *gomock.Controller
		service         *mock_testexample.MockService
		ctx             context.Context
		controller      testexample.Controller
	)

	setup := func() {
		observedZapCore, observedLogs = observer.New(zap.ErrorLevel)
		observedLogger = zap.New(observedZapCore)
		ctrl = gomock.NewController(t)
		service = mock_testexample.NewMockService(ctrl)
		ctx = context.Background()
		controller = testexample.NewController(service, observedLogger)
	}

	t.Run("service error", func(t *testing.T) {
		setup()
		service.EXPECT().Process(gomock.Eq(ctx), gomock.Any()).
			Return("", fmt.Errorf("error"))

		assert.Equal(t, "", controller.Process(ctx))
		require.Equal(t, 1, observedLogs.Len())
		assert.Equal(t, "Process error: error",
			observedLogs.All()[0].Message)
	})

	t.Run("basic flow", func(t *testing.T) {
		setup()
		service.EXPECT().Process(gomock.Eq(ctx), gomock.Any()).
			Return("", nil)

		assert.Equal(t, "", controller.Process(ctx))
		require.Equal(t, 0, observedLogs.Len())
	})
}
