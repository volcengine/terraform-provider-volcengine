package vpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVpcService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVpcService(c *ve.SdkClient) *VolcengineVpcService {
	return &VolcengineVpcService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVpcService) GetClient() *ve.SdkClient {
	return s.Client
}

func interfaceSlice2String(ele []interface{}) []string {
	var res []string
	for _, i := range ele {
		res = append(res, i.(string))
	}
	return res
}

func (s *VolcengineVpcService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "limit", "page", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListVPCInstances"
		condition["with_resource_statistic"] = true

		if ids, ok := condition["vpc_identity_list"]; ok {
			condition["vpc_identity_list"] = strings.Join(interfaceSlice2String(ids.([]interface{})), ",")
		}

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
		results, err = ve.ObtainSdkValue("Result.vpc_instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New(" Result.vpc_instances is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVpcService) ReadResource(resourceData *schema.ResourceData, vpcId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if vpcId == "" {
		vpcId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"vpc_identity_list": []interface{}{vpcId},
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
		return data, fmt.Errorf("Vpc %s not exist ", vpcId)
	}
	return data, err
}

func (s *VolcengineVpcService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "error")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("Vpc  status  error, status:%s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}

}

func (VolcengineVpcService) WithResourceResponseHandlers(vpc map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if ct, ok := vpc["cluster"]; ok {
			cluster := ct.(map[string]interface{})
			vpc["cluster_name"] = cluster["cluster_name"]
		}
		return vpc, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineVpcService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateVPCInstance",
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"vpc_name": {
					TargetField: "vpc_name",
				},
				"cluster_name": {
					TargetField: "cluster_name",
				},
				"desc": {
					TargetField: "desc",
				},
				"cidr": {
					TargetField: "cidr",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				bytes, _ := json.Marshal(call.SdkParam)
				logger.Debug(logger.ReqFormat, call.Action, string(bytes))
				return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.vpc_instance.vpc_identity", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineVpcService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "SetVPCInstanceDesc",
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"desc": {
					TargetField: "desc",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["vpc_identity"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				bytes, _ := json.Marshal(call.SdkParam)
				logger.Debug(logger.ReqFormat, call.Action, string(bytes))
				return s.Client.UniversalClient.DoCall(universalPost(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVpcService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVpcService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "vpc_identity_list",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		IdField:      "vpc_identity",
		NameField:    "vpc_name",
		CollectField: "vpc_instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"vpc_identity": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineVpcService) ReadResourceId(id string) string {
	return id
}

func universalPost(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veenedge",
		Version:     "2021-04-30",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
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
