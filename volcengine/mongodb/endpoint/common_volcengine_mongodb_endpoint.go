package endpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

func mongoDBEndpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'instanceId:endpointId'")
	}
	instanceId := items[0]
	endpointId := items[1]
	d.Set("endpoint_id", endpointId)
	d.Set("instance_id", instanceId)

	endpointService := NewMongoDBEndpointService(m.(*ve.SdkClient))
	endpoint, err := endpointService.GetEndpoint(instanceId, endpointId, "", "")
	if err != nil {
		return []*schema.ResourceData{d}, err
	}
	if objectId, ok := endpoint["ObjectId"]; ok {
		d.Set("object_id", objectId.(string))
	}
	if networkType, ok := endpoint["NetworkType"]; ok {
		d.Set("network_type", networkType)
	}
	nodeIds := make([]string, 0)
	eipIds := make([]string, 0)
	for _, address := range endpoint["DBAddresses"].([]interface{}) {
		logger.DebugInfo("address %v :", address)
		if nodeId, ok := address.(map[string]interface{})["NodeId"]; ok {
			nodeIds = append(nodeIds, nodeId.(string))
		}
		if eipId, ok := address.(map[string]interface{})["EipId"]; ok {
			eipIds = append(eipIds, eipId.(string))
		}
	}
	d.Set("mongos_node_ids", nodeIds)
	d.Set("eip_ids", eipIds)

	return []*schema.ResourceData{d}, nil
}
