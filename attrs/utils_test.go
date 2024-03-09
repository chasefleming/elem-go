package attrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataAttr(t *testing.T) {
	actual := DataAttr("foobar")
	expected := "data-foobar"
	assert.Equal(t, expected, actual)
}
