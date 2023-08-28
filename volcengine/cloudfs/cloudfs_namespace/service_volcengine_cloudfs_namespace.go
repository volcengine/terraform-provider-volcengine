package cloudfs_namespace

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

type VolcengineCloudfsNamespaceService struct {
	Client *ve.SdkClient
}

func NewCloudfsNamespaceService(c *ve.SdkClient) *VolcengineCloudfsNamespaceService {
	return &VolcengineCloudfsNamespaceService{
		Client: c,
	}
}

func (s *VolcengineCloudfsNamespaceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudfsNamespaceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListNs"
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

func (s *VolcengineCloudfsNamespaceService) ReadResource(resourceData *schema.ResourceData, namespaceId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if namespaceId == "" {
		namespaceId = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(namespaceId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid cloudfs namespace id")
	}
	req := map[string]interface{}{
		"FsName": ids[0],
		"NsId":   ids[1],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloudfs namespace %s not exist ", namespaceId)
	}

	// 针对无法读取的字段单独进行处理，使用用户原先的参数进行替换
	if tosAk := resourceData.Get("tos_ak"); len(tosAk.(string)) > 0 {
		data["TosAk"] = tosAk
	}
	if tosSk := resourceData.Get("tos_sk"); len(tosSk.(string)) > 0 {
		data["TosSk"] = tosSk
	}
	data["TosAccountId"] = resourceData.Get("tos_account_id")
	data["NsId"] = ids[1]

	return data, err
}

func (s *VolcengineCloudfsNamespaceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "ERROR", "DOWN", "CREATE_FAILED", "DELETE_FAILED")
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
					return nil, "", fmt.Errorf("cloudfs namespace status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return d, status.(string), err
		},
	}

}

func (VolcengineCloudfsNamespaceService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return data, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudfsNamespaceService) describeTask(taskId interface{}) (string, error) {
	req := map[string]interface{}{
		"TaskId": taskId,
	}
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeTask"), &req)
	if err != nil {
		return "", err
	}
	logger.Debug(logger.RespFormat, "DescribeTask", req, *resp)
	id, err := ve.ObtainSdkValue("Result.Id", *resp)
	if err != nil {
		return "", err
	}
	return id.(string), nil
}

func (s *VolcengineCloudfsNamespaceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNs",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// 通过 taskId 获取对应的 AccessId
				taskId, err := ve.ObtainSdkValue("Result.TaskId", *resp)
				if err != nil {
					return err
				}
				nsId, err := s.describeTask(taskId)
				if err != nil {
					return err
				}

				d.SetId(fmt.Sprintf("%s:%s", d.Get("fs_name"), nsId))
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

func (s *VolcengineCloudfsNamespaceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCloudfsNamespaceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNs",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				tmpId := d.Id()
				ids := strings.Split(tmpId, ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("error cloudfs namespace id: %s", tmpId)
				}
				(*call.SdkParam)["NsId"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading cloudfs namespace on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudfsNamespaceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "Id",
		CollectField: "namespaces",
	}
}

func (s *VolcengineCloudfsNamespaceService) ReadResourceId(id string) string {
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
