package scaling_instance_attach

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func importScalingInstanceAttach(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form ScalingGroupId:InstanceId")
	}
	err = data.Set("scaling_group_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("instance_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}

func formatInstanceIdsRequest(instanceId string) map[string]interface{} {
	var res = make(map[string]interface{}, 0)
	res["InstanceIds.1"] = instanceId
	return res
}

func convertSliceInterfaceToString(org []interface{}) []string {
	res := make([]string, len(org))
	for i, ele := range org {
		res[i] = ele.(string)
	}
	return res
}
