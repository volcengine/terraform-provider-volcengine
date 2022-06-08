package node

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var vkeNodeImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) > 100 {
		return []*schema.ResourceData{data}, fmt.Errorf("import nodes cannot exceed 100")
	}
	if len(items) < 1 {
		return []*schema.ResourceData{data}, fmt.Errorf("import nodes cannot at least 1")
	}
	return []*schema.ResourceData{data}, nil
}
