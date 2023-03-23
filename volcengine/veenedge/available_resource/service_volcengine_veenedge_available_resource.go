package available_resource

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineResourceService struct {
	Client *ve.SdkClient
}

func NewResourceService(c *ve.SdkClient) *VolcengineResourceService {
	return &VolcengineResourceService{
		Client: c,
	}
}

func (s *VolcengineResourceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineResourceService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "ListAvailableResourceInfo"

	bytes, _ := json.Marshal(condition)
	logger.Debug(logger.ReqFormat, action, string(bytes))

	if condition == nil {
		resp, err = s.Client.UniversalClient.DoCall(universalGet(action), nil)
		if err != nil {
			return data, err
		}
	} else {
		resp, err = s.Client.UniversalClient.DoCall(universalGet(action), &condition)
		if err != nil {
			return data, err
		}
	}
	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, condition, string(respBytes))
	results, err = ve.ObtainSdkValue("Result.regions", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New(" Result.regions is not Slice")
	}
	return data, err
}

func (s *VolcengineResourceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineResourceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineResourceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineResourceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineResourceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	return callbacks
}

func (s *VolcengineResourceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineResourceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_type": {
				TargetField: "instance_type",
			},
			"cloud_disk_type": {
				TargetField: "cloud_disk_type",
			},
			"bandwith_limit": {
				TargetField: "bandwith_limit",
			},
		},
		ContentType:  ve.ContentTypeJson,
		CollectField: "regions",
	}
}

func (s *VolcengineResourceService) ReadResourceId(id string) string {
	return id
}

func universalGet(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veenedge",
		Version:     "2021-04-30",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
