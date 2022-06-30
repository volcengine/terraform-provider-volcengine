package scaling_instance_attach

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var scalingInstanceImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	if err := data.Set("scaling_group_id", data.Id()); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

func formatInstanceIdsRequest(instanceIds []string) map[string]interface{} {
	var res = make(map[string]interface{}, 0)
	for i, id := range instanceIds {
		res[fmt.Sprintf("InstanceIds.%d", i+1)] = id
	}
	return res
}

func convertSliceInterfaceToString(org []interface{}) []string {
	res := make([]string, len(org))
	for i, ele := range org {
		res[i] = ele.(string)
	}
	return res
}
