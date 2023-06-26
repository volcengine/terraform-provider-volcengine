package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortAndStartTransJson1(t *testing.T) {
	req := map[string]interface{}{
		"Filter.ClusterId": "12345",
	}
	target := map[string]interface{}{
		"Filter": map[string]interface{}{
			"ClusterId": "12345",
		},
	}
	resp, _ := SortAndStartTransJson(req)
	assert.Equal(t, resp, target)
}

func TestSortAndStartTransJson2(t *testing.T) {
	req, _ := SortAndStartTransJson(map[string]interface{}{
		"Filter.Ids.1": "id123",
		"Filter.Ids.2": "id456",
	})
	target := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []interface{}{"id123", "id456"},
		},
	}
	resp, _ := SortAndStartTransJson(req)
	assert.Equal(t, resp, target)
}

func TestSortAndStartTransJson3(t *testing.T) {
	req, _ := SortAndStartTransJson(map[string]interface{}{
		"Filter.ClusterId": "12345",
		"Filter.Ids.1":     "id123",
		"Filter.Ids.2":     "id456",

		"Filter.Nets.1.Subnet": "subnet1",
		"Filter.Nets.2.Subnet": "subnet2",
		"Filter.Nets.3.Subnet": "subnet3",
	})
	target := map[string]interface{}{
		"Filter": map[string]interface{}{
			"ClusterId": "12345",
			"Ids":       []interface{}{"id123", "id456"},
			"Nets": []interface{}{
				map[string]interface{}{
					"Subnet": "subnet1",
				},
				map[string]interface{}{
					"Subnet": "subnet2",
				},
				map[string]interface{}{
					"Subnet": "subnet3",
				},
			},
		},
	}
	resp, _ := SortAndStartTransJson(req)
	assert.Equal(t, resp, target)

	str := `{"Filter":{"ClusterId":"12345","Ids":["id123","id456"],"Nets":[{"Subnet":"subnet1"},{"Subnet":"subnet2"},{"Subnet":"subnet3"}]}}`
	bytes, _ := json.Marshal(resp)
	assert.Equal(t, str, string(bytes))
}
