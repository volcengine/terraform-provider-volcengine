package allow_list_associate

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

type VolcengineMongodbAllowListAssociateService struct {
	Client     *ve.SdkClient
}

const (
	ActionAssociateAllowList       = "AssociateAllowList"
	ActionDisassociateAllowList    = "DisassociateAllowList"
	ActionDescribeDBInstanceDetail = "DescribeDBInstanceDetail"
)

func NewMongodbAllowListAssociateService(c *ve.SdkClient) *VolcengineMongodbAllowListAssociateService {
	return &VolcengineMongodbAllowListAssociateService{
		Client:     c,
	}
}

func (s *VolcengineMongodbAllowListAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongodbAllowListAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineMongodbAllowListAssociateService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		ids        []string
		instanceId string
		req        map[string]interface{}
		output     *map[string]interface{}
		results    interface{}
		ok         bool
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	ids = strings.Split(tmpId, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid id")
	}
	instanceId = ids[0]
	req = map[string]interface{}{
		"InstanceId": instanceId,
	}

	logger.Debug(logger.ReqFormat, ActionDescribeDBInstanceDetail, req)
	output, err = s.Client.UniversalClient.DoCall(getUniversalInfo(ActionDescribeDBInstanceDetail), &req)
	logger.Debug(logger.RespFormat, ActionDescribeDBInstanceDetail, req, *output)

	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Result.DBInstance", *output)
	if err != nil {
		return data, err
	}
	if data, ok = results.(map[string]interface{}); !ok {
		return data, errors.New("value is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("instance(%v) is not existed", instanceId)
	}
	return data, nil
}

func (s *VolcengineMongodbAllowListAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      time.Second,
		Pending:    []string{},
		Target:     target,
		Timeout:    timeout,
		MinTimeout: time.Second,

		Refresh: func() (result interface{}, state string, err error) {
			logger.DebugInfo("Refresh status", 0)
			var failStatus []string
			failStatus = append(failStatus, "CreateFailed")

			output, err := s.ReadResource(resourceData, id)
			if err != nil {
				logger.DebugInfo("ActionDescribeDBInstanceDetail failed", 0)
				return nil, "", err
			}
			var status interface{}
			status, err = ve.ObtainSdkValue("InstanceStatus", output)
			if err != nil {
				logger.DebugInfo("ObtainSdkValue InstanceStatus failed", 0)
				return nil, "", err
			}
			statusStr, ok := status.(string)
			if !ok {
				logger.DebugInfo("Type of InstanceStatus is not string", 0)
				return nil, "", fmt.Errorf("type of status if not string")
			}
			for _, v := range failStatus {
				if v == statusStr {
					return nil, "", fmt.Errorf("instance status error,status %s", status.(string))
				}
			}
			return output, statusStr, nil
		},
	}
}

func (s *VolcengineMongodbAllowListAssociateService) WithResourceResponseHandlers(association map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return association, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongodbAllowListAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionAssociateAllowList,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				var (
					output           *map[string]interface{}
					req              map[string]interface{}
					err              error
					ok               bool
					instanceIdInter  interface{}
					allowListIdInter interface{}
					instanceId       string
					allowListId      string
				)
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				instanceIdInter, ok = (*call.SdkParam)["InstanceId"]
				if !ok {
					return output, fmt.Errorf("please input instance_id")
				}
				instanceId, ok = instanceIdInter.(string)
				if !ok {
					return output, fmt.Errorf("type of instanceIdInter is not string")
				}
				allowListIdInter, ok = (*call.SdkParam)["AllowListId"]
				if !ok {
					return output, fmt.Errorf("please input allow_list_id")
				}
				allowListId, ok = allowListIdInter.(string)
				if !ok {
					return output, fmt.Errorf("type of allowListIdInter is not string")
				}
				req = make(map[string]interface{})
				req["InstanceIds"] = []string{instanceId}
				req["AllowListIds"] = []string{allowListId}
				output, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &req)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["InstanceId"], ":", (*call.SdkParam)["AllowListId"]))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongodbAllowListAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongodbAllowListAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionDisassociateAllowList,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := s.ReadResourceId(d.Id())
				ids := strings.Split(id, ":")
				instanceId := ids[0]
				allowListId := ids[1]
				(*call.SdkParam)["InstanceIds"] = []string{instanceId}
				(*call.SdkParam)["AllowListIds"] = []string{allowListId}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutDelete),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongodbAllowListAssociateService) DatasourceResources(data *schema.ResourceData, resource2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType: ve.ContentTypeJson,
	}
}

func (s *VolcengineMongodbAllowListAssociateService) ReadResourceId(id string) string {
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