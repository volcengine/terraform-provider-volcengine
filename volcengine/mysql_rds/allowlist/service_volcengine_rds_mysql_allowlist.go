package allowlist

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMysqlAllowListService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func (s *VolcengineRdsMysqlAllowListService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlAllowListService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return volc.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
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
		results, err = volc.ObtainSdkValue("Result.AllowLists", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AllowLists is not slice ")
		}
		return data, err
	})
}

func (s *VolcengineRdsMysqlAllowListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"AllowListId": id,
	}
	action := "DescribeAllowListDetail"
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	results, err = volc.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = results.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Rds instance %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsMysqlAllowListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineRdsMysqlAllowListService) WithResourceResponseHandlers(m map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return m, map[string]volc.ResponseConvert{}, nil
	}
	return []volc.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlAllowListService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateAllowList",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id, _ := volc.ObtainSdkValue("Result.AllowListId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "ModifyAllowList",
			ConvertMode: volc.RequestConvertInConvert,
			ContentType: volc.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"AllowListId":   data.Id(),
				"AllowListName": data.Get("allow_list_name").(string),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteAllowList",
			ConvertMode: volc.RequestConvertIgnore,
			ContentType: volc.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"AllowListId": data.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "AllowListName",
		IdField:      "AllowListId",
		CollectField: "allow_lists",
		ResponseConverts: map[string]volc.ResponseConvert{
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
		},
	}
}

func (s *VolcengineRdsMysqlAllowListService) ReadResourceId(id string) string {
	return id
}

func NewRdsMysqlAllowListService(client *volc.SdkClient) *VolcengineRdsMysqlAllowListService {
	return &VolcengineRdsMysqlAllowListService{
		Client:     client,
		Dispatcher: &volc.Dispatcher{},
	}
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}
