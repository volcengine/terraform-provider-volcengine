package download_url

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsDownloadTaskService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsDownloadTaskService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsDownloadTaskService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDownloadTasks"
		logger.Debug(logger.ReqFormat, action, condition)

		// Create a copy of condition to avoid modifying the original map which might be used in loop
		req := make(map[string]interface{})
		for k, v := range condition {
			req[k] = v
		}

		resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, resp)

		results, err := ve.ObtainSdkValue("RESPONSE.Tasks", *resp)
		if err != nil {
			return nil, err
		}

		if results == nil {
			return []interface{}{}, nil
		}

		taskList, ok := results.([]interface{})
		if !ok {
			return nil, fmt.Errorf("DescribeDownloadTasks response.Tasks is not a slice")
		}

		// Handle LogContextInfos transformation from Object to List
		// Also convert StartTime and EndTime from string to int64 for consistency with schema
		for _, task := range taskList {
			t, ok := task.(map[string]interface{})
			if !ok {
				continue
			}
			if info, ok := t["LogContextInfos"]; ok && info != nil {
				t["LogContextInfos"] = []interface{}{info}
			}

			// Convert timestamps
			if startTimeStr, ok := t["StartTime"].(string); ok && startTimeStr != "" {
				if parsed, err := time.Parse("2006-01-02 15:04:05", startTimeStr); err == nil {
					t["StartTime"] = parsed.Unix() * 1000
				}
			}
			if endTimeStr, ok := t["EndTime"].(string); ok && endTimeStr != "" {
				if parsed, err := time.Parse("2006-01-02 15:04:05", endTimeStr); err == nil {
					t["EndTime"] = parsed.Unix() * 1000
				}
			}

			// Get download URL if task is successful
			// Only try to get download URL if status is success or created_cut
			taskStatus, _ := t["TaskStatus"].(string)
			if taskStatus == "success" || taskStatus == "created_cut" {
				req := map[string]interface{}{
					"TaskId": t["TaskId"],
				}
				action := "DescribeDownloadUrl"
				logger.Debug(logger.ReqFormat, action, req)
				urlResp, urlErr := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.Default,
					HttpMethod:  ve.GET,
					Path:        []string{action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, &req)
				if urlErr == nil {
					logger.Debug(logger.RespFormat, action, urlResp)
					url, urlErr := ve.ObtainSdkValue("RESPONSE.DownloadUrl", *urlResp)
					if urlErr == nil && url != nil {
						t["DownloadUrl"] = url
					}
				}
			}
		}

		return taskList, nil
	})
}

func (v *VolcengineTlsDownloadTaskService) ReadDownloadUrl(m map[string]interface{}) (data []interface{}, err error) {
	req := map[string]interface{}{
		"TaskId": m["task_id"],
	}
	action := "DescribeDownloadUrl"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	url, err := ve.ObtainSdkValue("RESPONSE.DownloadUrl", *resp)
	if err != nil {
		return nil, err
	}

	if url == nil {
		return nil, fmt.Errorf("DownloadUrl not found in response")
	}

	res := map[string]interface{}{
		"task_id":      m["task_id"],
		"download_url": url,
	}
	return []interface{}{res}, nil
}

func (v *VolcengineTlsDownloadTaskService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = resourceData.Id()
	}

	req := map[string]interface{}{
		"TaskId":  id,
		"TopicId": resourceData.Get("topic_id"),
	}

	// Use DescribeDownloadTasks to find the task since there's no single Get API
	// DescribeDownloadTasks returns TaskList where StartTime/EndTime are strings (e.g. "2021-09-01 10:40:00")
	// But our schema expects Int (Unix timestamp)
	// We need to convert them or skip them if they are just display fields
	// Since we can't easily change the schema type now (it's Required), we should probably try to parse them back to int if possible
	// However, for ReadResource, we mostly care about ID and existence.
	// Let's rely on WithResourceResponseHandlers to handle data mapping if needed, or just let the default behavior happen.
	// The error "expected type 'int', got unconvertible type 'string'" comes from the SDK/Terraform trying to set the state.

	// Actually, the issue is that DescribeDownloadTasks returns string time, but schema defines int.
	// We should convert the string time from response to int timestamp before returning data.

	taskList, err := v.ReadResources(req)
	if err != nil {
		return nil, err
	}

	for _, task := range taskList {
		taskMap, ok := task.(map[string]interface{})
		if !ok {
			continue
		}

		if taskMap["TaskId"].(string) == id {
			// Convert StartTime and EndTime from string to int64
			if startTimeStr, ok := taskMap["StartTime"].(string); ok && startTimeStr != "" {
				// Try to parse using RFC3339 format first (e.g. 2021-09-01 10:40:00)
				if t, err := time.Parse("2006-01-02 15:04:05", startTimeStr); err == nil {
					// Convert back to int64 timestamp
					taskMap["StartTime"] = t.Unix() * 1000
				}
			} else if startTimeInt, ok := taskMap["StartTime"].(float64); ok {
				taskMap["StartTime"] = int64(startTimeInt)
			}

			if endTimeStr, ok := taskMap["EndTime"].(string); ok && endTimeStr != "" {
				if t, err := time.Parse("2006-01-02 15:04:05", endTimeStr); err == nil {
					// Convert back to int64 timestamp
					taskMap["EndTime"] = t.Unix() * 1000
				}
			} else if endTimeInt, ok := taskMap["EndTime"].(float64); ok {
				taskMap["EndTime"] = int64(endTimeInt)
			}

			return taskMap, nil
		}
	}

	return nil, fmt.Errorf("tls download task %s not exist", id)
}

func (v *VolcengineTlsDownloadTaskService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsDownloadTaskService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsDownloadTaskService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "CreateDownloadTask",
				ConvertMode: ve.RequestConvertAll,
				ContentType: ve.ContentTypeJson,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					// Handle LogContextInfos transformation from List to Object
					if infos, ok := (*call.SdkParam)["LogContextInfos"].([]interface{}); ok && len(infos) > 0 {
						(*call.SdkParam)["LogContextInfos"] = infos[0]
					}

					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
						ContentType: ve.ApplicationJSON,
						HttpMethod:  ve.POST,
						Path:        []string{call.Action},
						Client:      v.Client.BypassSvcClient.NewTlsClient(),
					}, call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp)

					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					id, _ := ve.ObtainSdkValue("RESPONSE.TaskId", *resp)
					d.SetId(id.(string))
					return nil
				},
			},
		},
	}
}

func (v *VolcengineTlsDownloadTaskService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// No Modify API available for download tasks
	return []ve.Callback{}
}

func (v *VolcengineTlsDownloadTaskService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{
		{
			Call: ve.SdkCall{
				Action:      "CancelDownloadTask",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"TaskId": data.Id(),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {

					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
						ContentType: ve.ApplicationJSON,
						HttpMethod:  ve.POST,
						Path:        []string{call.Action},
						Client:      v.Client.BypassSvcClient.NewTlsClient(),
					}, call.SdkParam)
				},
			},
		},
	}
}

func (v *VolcengineTlsDownloadTaskService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "download_tasks",
		IdField:      "TaskId",
		NameField:    "TaskName",
	}
}

func (v *VolcengineTlsDownloadTaskService) ReadResourceId(s string) string {
	return s
}

func NewTlsDownloadTaskService(c *ve.SdkClient) *VolcengineTlsDownloadTaskService {
	return &VolcengineTlsDownloadTaskService{
		Client: c,
	}
}
