package allowlist_associate

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_instance"
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
	results, err = volc.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if resultsMap, ok = results.(map[string]interface{}); !ok {
		return resultsMap, errors.New("Value is not map ")
	}
	if len(resultsMap) == 0 {
		return resultsMap, fmt.Errorf("Rds allowlist %s not exist ", ids[1])
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
		return data, fmt.Errorf("Rds allowlist associate %s not exist ", id)
	}
	return data, err
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
				// 规避 创建成功查询不到
				time.Sleep(5 * time.Second)
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
			ExtraRefresh: map[volc.ResourceService]*volc.StateRefresh{
				rds_mysql_instance.NewRdsMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    data.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
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
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				err := volc.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
				// 规避 解绑后删除实例OperationDenied: 无法执行该操作。
				time.Sleep(5 * time.Second)
				return err
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
			ExtraRefresh: map[volc.ResourceService]*volc.StateRefresh{
				rds_mysql_instance.NewRdsMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    data.Timeout(schema.TimeoutDelete),
					ResourceId: instanceId,
				},
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
