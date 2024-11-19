package image_share_permission

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

type VolcengineImageSharePermissionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewImageSharePermissionService(c *ve.SdkClient) *VolcengineImageSharePermissionService {
	return &VolcengineImageSharePermissionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineImageSharePermissionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineImageSharePermissionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 20, nil, func(condition map[string]interface{}) (data []interface{}, next string, err error) {
		action := "DescribeImageSharePermission"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, next, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Accounts", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}

		if _, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.Images is not Slice")
		}

		return results.([]interface{}), next, err
	})
}

func (s *VolcengineImageSharePermissionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid image_share_permission id: %s", id)
	}

	req := map[string]interface{}{
		"ImageId": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		var account map[string]interface{}
		if account, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if account["AccountId"] == ids[1] {
			data = account
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("image_share_permission %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineImageSharePermissionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineImageSharePermissionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineImageSharePermissionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyImageSharePermission",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"account_id": {
					TargetField: "AddAccounts.1",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				imageId := d.Get("image_id").(string)
				accountId := d.Get("account_id").(string)
				d.SetId(imageId + ":" + accountId)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineImageSharePermissionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineImageSharePermissionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyImageSharePermission",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid image_share_permission id: %s", d.Id())
				}
				(*call.SdkParam)["ImageId"] = ids[0]
				(*call.SdkParam)["RemoveAccounts.1"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineImageSharePermissionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "accounts",
	}
}

func (s *VolcengineImageSharePermissionService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
