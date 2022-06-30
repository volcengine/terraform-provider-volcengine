package volume_attach

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResourceNotFoundError(t *testing.T) {
	parts := strings.Split("vol-3tzl52wubz3b9fciw7ev:i-4ay59ww7dq8dt9c29hd4", ":")
	assert.True(t, ve.ResourceNotFoundError(fmt.Errorf("volume %s does not associate instances", parts[0])))
}
