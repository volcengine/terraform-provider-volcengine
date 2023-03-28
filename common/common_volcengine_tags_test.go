package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FilterSystemTags(t *testing.T) {
	tags := []interface{}{
		map[string]interface{}{
			"Key":   "tag1",
			"Value": "value1",
		},
		map[string]interface{}{
			"Key":   "tag2",
			"Value": "value2",
		},
		map[string]interface{}{
			"Key":   "volc:ecs:linkedresource",
			"Value": "trn:ecs:cn-beijing:222222",
		},
	}

	res := FilterSystemTags(tags)
	assert.Equal(t, res, []interface{}{
		map[string]interface{}{
			"Key":   "tag1",
			"Value": "value1",
		},
		map[string]interface{}{
			"Key":   "tag2",
			"Value": "value2",
		},
	})
}
