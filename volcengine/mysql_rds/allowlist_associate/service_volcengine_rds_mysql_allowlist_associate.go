package allowlist_associate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMysqlAllowListAssociateService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func (s *VolcengineRdsMysqlAllowListAssociateService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlAllowListAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRdsMysqlAllowListAssociateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRdsMysqlAllowListAssociateService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineRdsMysqlAllowListAssociateService) WithResourceResponseHandlers(m map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return m, map[string]volc.ResponseConvert{}, nil
	}
	return []volc.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlAllowListAssociateService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	instanceId := data.Get("instance_id").(string)
	allowListId := data.Get("allow_list_id").(string)
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "AssociateAllowList",
			ContentType: volc.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceIds":  []string{instanceId},
				"AllowListIds": []string{allowListId},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				d.SetId(fmt.Sprint(instanceId, ":", allowListId))
				return nil
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListAssociateService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	return []volc.Callback{}
}

func (s *VolcengineRdsMysqlAllowListAssociateService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	instanceId := data.Get("instance_id").(string)
	allowListId := data.Get("allow_list_id").(string)
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DisassociateAllowList",
			ContentType: volc.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceIds":  []string{instanceId},
				"AllowListIds": []string{allowListId},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListAssociateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{}
}

func (s *VolcengineRdsMysqlAllowListAssociateService) ReadResourceId(id string) string {
	return id
}

func NewRdsMysqlAllowListAssociateService(client *volc.SdkClient) *VolcengineRdsMysqlAllowListAssociateService {
	return &VolcengineRdsMysqlAllowListAssociateService{
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
