package rds_postgresql_allowlist_version_upgrade

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlAllowlistVersionUpgradeService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlAllowlistVersionUpgradeService(c *ve.SdkClient) *VolcengineRdsPostgresqlAllowlistVersionUpgradeService {
	return &VolcengineRdsPostgresqlAllowlistVersionUpgradeService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return []interface{}{}, nil
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// No remote readable state; keep empty
	return map[string]interface{}{}, nil
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpgradeAllowListVersion",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id": {TargetField: "InstanceId"},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(d.Get("instance_id").(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	// No remote delete; just remove state
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// No modify operation for upgrade-only resource
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsPostgresqlAllowlistVersionUpgradeService) ReadResourceId(id string) string {
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
