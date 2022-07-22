package instance

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineESCloudInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewESCloudInstanceService(c *ve.SdkClient) *VolcengineESCloudInstanceService {
	return &VolcengineESCloudInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineESCloudInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineESCloudInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeInstances"
		// 重新组织 Filter 的格式
		if filter, filterExist := condition["Filters"]; filterExist {
			newFilter := make([]interface{}, 0)
			for k, v := range filter.(map[string]interface{}) {
				newFilter = append(newFilter, map[string]interface{}{
					"Name":   k,
					"Values": v,
				})
			}
			condition["Filters"] = newFilter
		}
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
		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Instances is not Slice")
		}

		// get instance node info
		for index, ele := range data {
			ins := ele.(map[string]interface{})
			con := &map[string]interface{}{
				"InstanceId": ins["InstanceId"],
			}
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstanceNodes"), con)
			if err != nil {
				return data, err
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, "DescribeInstanceNodes", con, string(respBytes))
			results, err = ve.ObtainSdkValue("Result.Nodes", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			data[index].(map[string]interface{})["Nodes"] = results

			// 插件系统只有在 Running 状态下才存在
			if ins["Status"] != "Running" {
				continue
			}
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstancePlugins"), con)
			if err != nil {
				return data, err
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, "DescribeInstancePlugins", con, string(respBytes))
			results, err = ve.ObtainSdkValue("Result.InstancePlugins", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			data[index].(map[string]interface{})["Plugins"] = results
		}
		return data, err
	})
}

func (s *VolcengineESCloudInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filters": map[string]interface{}{
			"InstanceId": []string{id},
		},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Instance %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineESCloudInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineESCloudInstanceService) WithResourceResponseHandlers(cluster map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineESCloudInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineESCloudInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineESCloudInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineESCloudInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filters.InstanceId",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filters.Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"charge_types": {
				TargetField: "Filters.ChargeType",
				ConvertType: ve.ConvertJsonArray,
			},
			"names": {
				TargetField: "Filters.InstanceName",
				ConvertType: ve.ConvertJsonArray,
			},
			"versions": {
				TargetField: "Filters.Version",
				ConvertType: ve.ConvertJsonArray,
			},
			"zone_ids": {
				TargetField: "Filters.ZoneId",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		IdField:      "InstanceId",
		CollectField: "instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"EnableESPublicNetwork": {
				TargetField: "enable_es_public_network",
			},
			"EnableESPrivateNetwork": {
				TargetField: "enable_es_private_network",
			},
			"ESPublicDomain": {
				TargetField: "es_public_domain",
			},
			"ESPrivateDomain": {
				TargetField: "es_private_domain",
			},
			"ESPrivateEndpoint": {
				TargetField: "es_private_endpoint",
			},
			"ESPublicEndpoint": {
				TargetField: "es_public_endpoint",
			},
			"ESInnerEndpoint": {
				TargetField: "es_inner_endpoint",
			},
			"CPU": {
				TargetField: "cpu",
			},
			"VPC": {
				TargetField: "vpc",
			},
		},
	}
}

func (s *VolcengineESCloudInstanceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ESCloud",
		Version:     "2018-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
