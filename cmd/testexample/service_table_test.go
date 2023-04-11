package testexample_test

import (
	"context"
	"github.com/eayzenberg-paxos/go-intro/cmd/testexample"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testProcessCase struct {
	name           string
	args           testexample.ServiceArgs
	expectedResult string
	expectedError  error
}

var cases = []testProcessCase{
	{
		name:           "no error",
		args:           testexample.ServiceArgs{Param1: "something"},
		expectedResult: "",
		expectedError:  nil,
	},
}

func TestProcessWithTable(t *testing.T) {
	for _, processCase := range cases {
		t.Run(processCase.name, func(t *testing.T) {
			impl := testexample.ServiceImpl{}

			process, err := impl.Process(context.Background(), processCase.args)

			assert.Equal(t, processCase.expectedResult, process)
			assert.Equal(t, processCase.expectedError, err)
		})
	}
}
