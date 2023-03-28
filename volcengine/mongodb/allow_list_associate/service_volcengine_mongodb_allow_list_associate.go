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
	Client *ve.SdkClient
}

const (
	ActionAssociateAllowList      = "AssociateAllowList"
	ActionDisassociateAllowList   = "DisassociateAllowList"
	ActionDescribeAllowListDetail = "DescribeAllowListDetail"
)

func NewMongodbAllowListAssociateService(c *ve.SdkClient) *VolcengineMongodbAllowListAssociateService {
	return &VolcengineMongodbAllowListAssociateService{
		Client: c,
	}
}

func (s *VolcengineMongodbAllowListAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongodbAllowListAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineMongodbAllowListAssociateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		targetInstanceId string
		allowListId      string
		output           *map[string]interface{}
		resultsMap       = make(map[string]interface{})
		instanceMap      = make(map[string]interface{})
		results          interface{}
		ok               bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid id")
	}
	targetInstanceId = ids[0]
	allowListId = ids[1]
	req := map[string]interface{}{
		"AllowListId": allowListId,
	}
	logger.Debug(logger.ReqFormat, ActionDescribeAllowListDetail, req)
	output, err = s.Client.UniversalClient.DoCall(getUniversalInfo(ActionDescribeAllowListDetail), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, ActionDescribeAllowListDetail, req, *output)
	results, err = ve.ObtainSdkValue("Result", *output)
	if err != nil {
		return data, err
	}
	if resultsMap, ok = results.(map[string]interface{}); !ok {
		return resultsMap, errors.New("Value is not map ")
	}
	if len(resultsMap) == 0 {
		return resultsMap, fmt.Errorf("MongoDB allowlist %s not exist ", allowListId)
	}
	instances := resultsMap["AssociatedInstances"].([]interface{})
	for _, instance := range instances {
		if instanceMap, ok = instance.(map[string]interface{}); !ok {
			return data, errors.New("instance is not map ")
		}
		if instanceMap["InstanceId"].(string) == targetInstanceId {
			data = resultsMap
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("MongoDB allowlist associate %s not associate ", id)
	}
	return data, err
}

func (s *VolcengineMongodbAllowListAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      1 * time.Second,
		Pending:    []string{},
		Target:     target,
		Timeout:    timeout,
		MinTimeout: 1 * time.Second,

		Refresh: func() (result interface{}, state string, err error) {
			logger.DebugInfo("Refreshing")
			output, err := s.ReadResource(resourceData, id)
			if err != nil {
				if strings.Contains(err.Error(), "not associate") {
					return output, "UnAttached", nil
				}
				return nil, "", err
			}
			return output, "Attached", nil
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
				Target:  []string{"Attached"},
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
				Target:  []string{"UnAttached"},
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
