package rds_postgresql_allowlist_associate

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
)

type VolcengineRdsPostgresqlAllowlistAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlAllowlistAssociateService(c *ve.SdkClient) *VolcengineRdsPostgresqlAllowlistAssociateService {
	return &VolcengineRdsPostgresqlAllowlistAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("invalid postgresql allowlist associate id: %v", id)
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
		return resultsMap, fmt.Errorf("Rds Postgresql allowlist %s not exist ", ids[1])
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
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rds_postgresql_allowlist_associate %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineRdsPostgresqlAllowlistAssociateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateAllowList",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				instanceId := d.Get("instance_id").(string)
				allowListId := d.Get("allow_list_id").(string)
				(*call.SdkParam)["InstanceIds"] = []string{instanceId}
				(*call.SdkParam)["AllowListIds"] = []string{allowListId}

				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				allowListId := d.Get("allow_list_id").(string)
				d.SetId(fmt.Sprint(instanceId, ":", allowListId))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rds_postgresql_instance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateAllowList",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return nil, fmt.Errorf("invalid postgresql allowlist associate id: %v", d.Id())
				}
				(*call.SdkParam)["InstanceIds"] = []string{ids[0]}
				(*call.SdkParam)["AllowListIds"] = []string{ids[1]}

				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rds_postgresql_instance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsPostgresqlAllowlistAssociateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
