package instance_types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineInstanceTypeService struct {
	Client *ve.SdkClient
}

func NewInstanceTypeService(c *ve.SdkClient) *VolcengineInstanceTypeService {
	return &VolcengineInstanceTypeService{
		Client: c,
	}
}

func (s *VolcengineInstanceTypeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineInstanceTypeService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "ListInstanceTypes"

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
	results, err = ve.ObtainSdkValue("Result.instance_type_configs", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New(" Result.instance_type_configs is not Slice")
	}
	return data, err
}

func (s *VolcengineInstanceTypeService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineInstanceTypeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineInstanceTypeService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineInstanceTypeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineInstanceTypeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	return callbacks
}

func (s *VolcengineInstanceTypeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineInstanceTypeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		CollectField: "instance_type_configs",
	}
}

func (s *VolcengineInstanceTypeService) ReadResourceId(id string) string {
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
