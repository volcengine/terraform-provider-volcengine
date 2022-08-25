package allow_list

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoDBAllowListService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewMongoDBAllowListService(c *ve.SdkClient) *VolcengineMongoDBAllowListService {
	return &VolcengineMongoDBAllowListService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineMongoDBAllowListService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongoDBAllowListService) readAllowListDetails(allowListId string) (allowList interface{}, err error) {
	var (
		resp *map[string]interface{}
		//ok   bool
	)
	action := "DescribeAllowListDetail"
	cond := map[string]interface{}{
		"AllowListId": allowListId,
	}
	logger.Debug(logger.RespFormat, action, cond)
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &cond)
	if err != nil {
		return allowList, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	allowList, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return allowList, err
	}
	return allowList, err
}

func (s *VolcengineMongoDBAllowListService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "AllowListId",
		NameField:    "AllowListName",
		CollectField: "allow_lists",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
			"VPC": {
				TargetField: "vpc",
			},
		},
	}
}

func (s *VolcengineMongoDBAllowListService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	action := "DescribeAllowLists"
	logger.Debug(logger.ReqFormat, action, condition)
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
	logger.Debug(logger.RespFormat, action, condition, *resp)
	results, err = ve.ObtainSdkValue("Result.AllowLists", *resp)
	if err != nil {
		logger.DebugInfo("ve.ObtainSdkValue return :%v", err)
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	allowLists, ok := results.([]interface{})
	if !ok {
		return data, fmt.Errorf("DescribeAllowLists responsed instances is not a slice")
	}

	for _, ele := range allowLists {
		allowList := ele.(map[string]interface{})
		id := allowList["AllowListId"].(string)

		detail, err := s.readAllowListDetails(id)
		if err != nil {
			logger.DebugInfo("read allow list %s detail failed,err:%v.", id, err)
			data = append(data, ele)
			continue
		}
		allowList["AllowList"] = detail.(map[string]interface{})["AllowList"]
		allowList["AssociatedInstances"] = detail.(map[string]interface{})["AssociatedInstances"]

		logger.DebugInfo("ins:   %v", allowList)
		data = append(data, allowList)
	}
	return data, nil
}

func (s *VolcengineMongoDBAllowListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("the format of import id must be 'allow_list:instanceId'")
	}
	regionId := resourceData.Get("region_id")
	req := map[string]interface{}{
		"RegionId":   regionId,
		"InstanceId": parts[1],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("allowlist %s is not exist", id)
	}
	return data, err
}

func (s *VolcengineMongoDBAllowListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineMongoDBAllowListService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBAllowListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAllowList",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				id, _ := ve.ObtainSdkValue("Result.AllowListId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBAllowListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	if resourceData.HasChange("allow_list_name") || resourceData.HasChange("allow_list_desc") ||
		resourceData.HasChange("allow_list") || resourceData.HasChange("apply_instance_num") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyAllowList",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["AllowListId"] = d.Id()
					if resourceData.HasChange("allow_list_name") {
						(*call.SdkParam)["AllowListName"] = resourceData.Get("allow_list_name")
					}
					if resourceData.HasChange("allow_list_desc") {
						(*call.SdkParam)["AllowListDesc"] = resourceData.Get("AllowListDesc")
					}
					if resourceData.HasChange("allow_list") {
						(*call.SdkParam)["AllowList"] = resourceData.Get("allow_list")
					}
					if resourceData.HasChange("apply_instance_num") {
						(*call.SdkParam)["ApplyInstanceNum"] = resourceData.Get("apply_instance_num")
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func (s *VolcengineMongoDBAllowListService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAllowList",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["AllowListId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBAllowListService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "mongodb",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
