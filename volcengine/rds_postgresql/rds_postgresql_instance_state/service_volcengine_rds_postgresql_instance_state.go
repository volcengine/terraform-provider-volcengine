package rds_postgresql_instance_state

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	rds_postgresql_instance "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
)

type VolcengineRdsPostgresqlInstanceStateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceStateService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceStateService {
	return &VolcengineRdsPostgresqlInstanceStateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceStateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceStateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		return []interface{}{}, nil
	})
}

func (s *VolcengineRdsPostgresqlInstanceStateService) ReadResource(*schema.ResourceData, string) (map[string]interface{}, error) {
	return nil, nil
}

func (s *VolcengineRdsPostgresqlInstanceStateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)
			rid := id
			if rid == "" {
				if v, ok := resourceData.Get("instance_id").(string); ok && v != "" {
					rid = v
				}
			}
			instSvc := rds_postgresql_instance.NewRdsPostgresqlInstanceService(s.Client)
			demo, err = instSvc.ReadResource(resourceData, rid)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("instance_status", demo)
			if err != nil {
				return nil, "", err
			}
			return demo, status.(string), err
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceStateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	restart := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RestartDBInstance",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"action": {
					Ignore: true,
				},
				"instance_id": {
					TargetField: "InstanceId",
					ForceGet:    true,
				},
				"apply_scope": {
					TargetField: "ApplyScope",
				},
				"custom_node_ids": {
					TargetField: "CustomNodeIds",
					ConvertType: ve.ConvertWithN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				d.SetId("state:" + instanceId)
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rds_postgresql_instance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{restart}
}

func (VolcengineRdsPostgresqlInstanceStateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceStateService) ModifyResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceStateService) RemoveResource(*schema.ResourceData, *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceStateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsPostgresqlInstanceStateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
