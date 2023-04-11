package testexample_test

import (
	"context"
	"github.com/eayzenberg-paxos/go-intro/cmd/testexample"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {
	impl := testexample.ServiceImpl{}

	process, err := impl.Process(context.Background(), testexample.ServiceArgs{})

	assert.Equal(t, "", process)
	assert.Equal(t, nil, err)
}
