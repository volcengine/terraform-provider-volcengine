package cloudfs_access

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

type VolcengineCloudfsAccessService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineCloudfsAccessService {
	return &VolcengineCloudfsAccessService{
		Client: c,
	}
}

func (s *VolcengineCloudfsAccessService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudfsAccessService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListAccess"
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineCloudfsAccessService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")

	req := map[string]interface{}{
		"FsName": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); ok {
			if data["AccessId"] == ids[1] {
				return data, nil
			}
		}
	}
	return data, fmt.Errorf("access not exist: %s", id)
}

func (s *VolcengineCloudfsAccessService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "ERROR", "DOWN", "CREATE_FAILED",
				"DELETE_FAILED", "ENABLE_VPC_ROUTE_FAILED", "DISABLE_VPC_ROUTE_FAILED")
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
					return nil, "", fmt.Errorf("Access  status  error, status:%s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}

}

func (VolcengineCloudfsAccessService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineCloudfsAccessService) describeTask(taskId interface{}) (string, error) {
	req := map[string]interface{}{
		"TaskId": taskId,
	}
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeTask"), &req)
	if err != nil {
		return "", err
	}
	logger.Debug(logger.RespFormat, "DescribeTask", req, *resp)
	id, err := ve.ObtainSdkValue("Result.AccessId", *resp)
	if err != nil {
		return "", err
	}
	return id.(string), nil
}

func (s *VolcengineCloudfsAccessService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableCacheAccess",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if v, ok := (*call.SdkParam)["SubnetId"]; ok {
					action := "DescribeSubnetAttributes"
					req := map[string]interface{}{
						"SubnetId": v.(string),
					}
					resp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
					if err != nil {
						return false, err
					}
					logger.Debug(logger.RespFormat, action, req, *resp)
					vpcId, err := ve.ObtainSdkValue("Result.VpcId", *resp)
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["VpcId"] = vpcId
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// 通过 taskId 获取对应的 AccessId
				taskId, err := ve.ObtainSdkValue("Result.TaskId", *resp)
				if err != nil {
					return err
				}
				accessId, err := s.describeTask(taskId)
				if err != nil {
					return err
				}

				d.SetId(fmt.Sprintf("%s:%s", d.Get("fs_name"), accessId))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"UP"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("fs_name").(string)
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineCloudfsAccessService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("vpc_route_enabled") {
		_, n := resourceData.GetChange("vpc_route_enabled")
		action := "DisableVpcRoute"
		if n.(bool) {
			action = "EnableVpcRoute"
		}
		enableCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      action,
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["AccessId"] = strings.Split(resourceData.Id(), ":")[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"UP"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("fs_name").(string)
				},
			},
		}
		callbacks = append(callbacks, enableCallback)
	}
	return callbacks
}

func (s *VolcengineCloudfsAccessService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableCacheAccess",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AccessId": strings.Split(resourceData.Id(), ":")[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("fs_name").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudfsAccessService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "AccessId",
		CollectField: "accesses",
	}
}

func (s *VolcengineCloudfsAccessService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cfs",
		Version:     "2022-02-02",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cfs",
		Version:     "2022-02-02",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
