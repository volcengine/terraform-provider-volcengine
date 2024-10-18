package vedb_mysql_allowlist_associate

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vedb_mysql/vedb_mysql_instance"
)

type VolcengineVedbMysqlAllowlistAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVedbMysqlAllowlistAssociateService(c *ve.SdkClient) *VolcengineVedbMysqlAllowlistAssociateService {
	return &VolcengineVedbMysqlAllowlistAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results     interface{}
		resultsMap  map[string]interface{}
		instanceMap map[string]interface{}
		instances   []interface{}
		ok          bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, err
	}
	req := map[string]interface{}{
		"AllowListId": ids[1],
	}
	action := "DescribeAllowListDetail"
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if resultsMap, ok = results.(map[string]interface{}); !ok {
		return resultsMap, errors.New("Value is not map ")
	}
	if len(resultsMap) == 0 {
		return resultsMap, fmt.Errorf("veDB allowlist %s not exist ", ids[1])
	}
	logger.Debug(logger.ReqFormat, action, resultsMap)
	instances = resultsMap["AssociatedInstances"].([]interface{})
	logger.Debug(logger.ReqFormat, action, instances)
	for _, instance := range instances {
		if instanceMap, ok = instance.(map[string]interface{}); !ok {
			return data, errors.New("instance is not map ")
		}
		if len(instanceMap) == 0 {
			continue
		}
		if instanceMap["InstanceId"].(string) == ids[0] {
			data = resultsMap
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("veDB allowlist associate %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vedb_mysql_allowlist_associate status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	allowListId := resourceData.Get("allow_list_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateAllowList",
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceIds":  []string{instanceId},
				"AllowListIds": []string{allowListId},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint(instanceId, ":", allowListId))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVedbMysqlAllowlistAssociateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	allowListId := resourceData.Get("allow_list_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateAllowList",
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceIds":  []string{instanceId},
				"AllowListIds": []string{allowListId},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				err := ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
				return err
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: instanceId,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineVedbMysqlAllowlistAssociateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vedbm",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
