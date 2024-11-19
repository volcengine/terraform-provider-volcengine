package image_import

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineImageImportService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewImageImportService(c *ve.SdkClient) *VolcengineImageImportService {
	return &VolcengineImageImportService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineImageImportService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineImageImportService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 20, nil, func(condition map[string]interface{}) (data []interface{}, next string, err error) {
		action := "DescribeImages"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, next, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Images", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}

		if _, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.Images is not Slice")
		}

		return results.([]interface{}), next, err
	})
}

func (s *VolcengineImageImportService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ImageIds.1": id,
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
		return data, fmt.Errorf("image_import %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineImageImportService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

			if err = resource.Retry(5*time.Minute, func() *resource.RetryError {
				d, err = s.describeTasks(id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}

			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("image_import status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineImageImportService) describeTasks(resourceId string) (task map[string]interface{}, err error) {
	action := "DescribeTasks"
	req := map[string]interface{}{
		"ResourceIds.1": resourceId,
	}

	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return task, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	results, err := ve.ObtainSdkValue("Result.Tasks", *resp)
	if err != nil {
		return task, err
	}
	if results == nil {
		results = []interface{}{}
	}
	tasks, ok := results.([]interface{})
	if !ok {
		return task, errors.New("Result.Tasks is not Slice")
	}
	for _, v := range tasks {
		if task, ok = v.(map[string]interface{}); !ok {
			return nil, errors.New("The value of Result.Tasks is not map ")
		}
	}
	if len(task) == 0 {
		return task, fmt.Errorf("task with resource %s not exist", resourceId)
	}
	return task, nil
}

func (VolcengineImageImportService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineImageImportService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ImportImage",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.ImageId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Succeeded"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineImageImportService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyImageAttribute",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"image_name": {
					TargetField: "ImageName",
				},
				"description": {
					TargetField: "Description",
				},
				"boot_mode": {
					TargetField: "BootMode",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["ImageId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "CreateTags", "DeleteTags", "image", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineImageImportService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteImages",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ImageIds.1": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading image on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineImageImportService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineImageImportService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineImageImportService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "ecs",
		ResourceType:         "image",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
