package cloud_identity_user_attachment

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudIdentityUserAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudIdentityUserAttachmentService(c *ve.SdkClient) *VolcengineCloudIdentityUserAttachmentService {
	return &VolcengineCloudIdentityUserAttachmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudIdentityUserAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudIdentityUserAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineCloudIdentityUserAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results  []interface{}
		ok       bool
		attached bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid user attachment id: %v", id)
	}
	groupId := ids[0]
	userId := ids[1]

	req := map[string]interface{}{
		"GroupId": groupId,
	}
	// query user info in group
	results, err = ve.WithPageNumberQuery(req, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListGroupMembers"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return results, err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		members, err := ve.ObtainSdkValue("Result.Members", *resp)
		if err != nil {
			return results, err
		}
		if members == nil {
			members = []interface{}{}
		}
		if results, ok = members.([]interface{}); !ok {
			return results, errors.New("Result.Members is not Slice")
		}
		return results, err
	})
	if err != nil {
		return data, err
	}
	for _, v := range results {
		var user map[string]interface{}
		user, ok = v.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		}
		if user["UserId"].(string) == userId {
			attached = true
			break
		}
	}
	if !attached {
		return data, fmt.Errorf("The user %s is not associate to the group %s ", userId, groupId)
	}
	return data, err
}

func (s *VolcengineCloudIdentityUserAttachmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineCloudIdentityUserAttachmentService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudIdentityUserAttachmentService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	groupId := resourceData.Get("group_id").(string)
	userId := resourceData.Get("user_id").(string)

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddUserToGroup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"GroupId": groupId,
				"UserId":  userId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%v:%v", (*call.SdkParam)["GroupId"], (*call.SdkParam)["UserId"]))
				return nil
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudIdentityUserAttachmentService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCloudIdentityUserAttachmentService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	groupId := resourceData.Get("group_id").(string)
	userId := resourceData.Get("user_id").(string)

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveUserFromGroup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"GroupId": groupId,
				"UserId":  userId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading cloud identity user attacnment on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudIdentityUserAttachmentService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineCloudIdentityUserAttachmentService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cloudidentity",
		Version:     "2023-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cloudidentity",
		Version:     "2023-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
