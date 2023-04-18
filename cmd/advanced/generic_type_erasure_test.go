package advanced

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func switchByGeneric[T any]() string {
	var t T
	switch reflect.TypeOf(t) {
	case reflect.TypeOf(1):
		return "It was int!"
	case reflect.TypeOf(""):
		return "It was string!"
	default:
		return "I don't know!"
	}
}

func TestSwitchByGeneric(t *testing.T) {
	assert.Equal(t, "It was int!", switchByGeneric[int]())
	assert.Equal(t, "It was string!", switchByGeneric[string]())
	assert.Equal(t, "I don't know!", switchByGeneric[struct{ Runner[float32] }]())
}
