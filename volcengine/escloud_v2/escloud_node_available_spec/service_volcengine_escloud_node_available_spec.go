package escloud_node_available_spec

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEscloudNodeAvailableSpecsService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEscloudNodeAvailableSpecsService(c *ve.SdkClient) *VolcengineEscloudNodeAvailableSpecsService {
	return &VolcengineEscloudNodeAvailableSpecsService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEscloudNodeAvailableSpecsService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEscloudNodeAvailableSpecsService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeNodeAvailableSpecs"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		result, err := ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if resultMap, ok := result.(map[string]interface{}); ok {
			results = []interface{}{resultMap}
		}

		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result is not Slice ")
		}
		return data, err
	})
}

func (s *VolcengineEscloudNodeAvailableSpecsService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineEscloudNodeAvailableSpecsService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineEscloudNodeAvailableSpecsService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEscloudNodeAvailableSpecsService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEscloudNodeAvailableSpecsService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEscloudNodeAvailableSpecsService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEscloudNodeAvailableSpecsService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "node_specs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"AZAvailableSpecsSoldOut": {
				TargetField: "az_available_specs_sold_out",
			},
			"CPU": {
				TargetField: "cpu",
			},
		},
	}
}

func (s *VolcengineEscloudNodeAvailableSpecsService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ESCloud",
		Version:     "2023-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
