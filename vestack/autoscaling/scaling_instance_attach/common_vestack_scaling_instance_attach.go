package scaling_instance_attach

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var scalingInstanceImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	var (
		setMap      = make(map[string]bool)
		instanceIds []string
	)
	if len(items) < 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	for _, id := range items[1:] {
		if len(id) == 0 {
			continue
		}
		if _, ok := setMap[id]; ok {
			continue
		}
		setMap[id] = true
		instanceIds = append(instanceIds, id)
	}
	if err := data.Set("scaling_group_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("instance_ids", instanceIds); err != nil {
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
