package volume

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sizeConvertFunc(t *testing.T) {
	assert.Equal(t, sizeConvertFunc(10), 10)
	assert.Equal(t, sizeConvertFunc("20"), 20)
}
