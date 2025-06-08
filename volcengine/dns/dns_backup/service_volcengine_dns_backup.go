package dns_backup

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineDnsBackupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewDnsBackupService(c *ve.SdkClient) *VolcengineDnsBackupService {
	return &VolcengineDnsBackupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineDnsBackupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineDnsBackupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListUserZoneBackups"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.BackupInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.BackupInfos is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineDnsBackupService) ReadResource(resourceData *schema.ResourceData, id string) (result map[string]interface{}, err error) {
	var (
		data     map[string]interface{}
		results  []interface{}
		ok       bool
		zid      string
		bucketId string
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("format of dns backup resource id is invalid,%s", id)
	}
	zid = parts[0]
	bucketId = parts[1]
	zidInt, err := strconv.Atoi(zid)
	if err != nil {
		return data, fmt.Errorf(" ZID cannot convert to int ")
	}
	req := map[string]interface{}{
		"ZID": zidInt,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}

		if data["BackupID"] == bucketId {
			result = data
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("dns_backup %s not exist ", id)
	}
	return result, err
}

func (s *VolcengineDnsBackupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
	}
}

func (s *VolcengineDnsBackupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUserZoneBackup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"zid": {
					TargetField: "ZID",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var (
					data            map[string]interface{}
					ok              bool
					result          map[string]interface{}
					createRspResult map[string]interface{}
				)
				zid := d.Get("zid")
				req := map[string]interface{}{
					"ZID": zid,
				}
				results, err := s.ReadResources(req)
				if err != nil {
					return err
				}

				for _, v := range results {
					if data, ok = v.(map[string]interface{}); !ok {
						return errors.New("CreateResource AfterCall Value is not map ")
					}

					if createRspResult, ok = (*resp)["Result"].(map[string]interface{}); !ok {
						return errors.New("create result is not map ")
					}

					parts := strings.Split(createRspResult["BackupTime"].(string), "_")
					// 这里 dns 后面要改 去掉下划线 兼容一下
					//if len(parts) != 2 {
					//	return errors.New("BackupTime is not map")
					//}

					if data["BackupTime"] == parts[0] {
						result = data
					}
				}

				if len(result) == 0 {
					return fmt.Errorf("dns_backup %s not exist ", zid)
				}

				d.SetId(fmt.Sprintf("%s:%s", strconv.Itoa(zid.(int)), result["BackupID"].(string)))

				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineDnsBackupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ZID": {
				TargetField: "zid",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineDnsBackupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineDnsBackupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var (
		zid      string
		backupId string
	)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteUserZoneBackup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := s.ReadResourceId(resourceData.Id())
				parts := strings.Split(id, ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("the id format must be 'ZID:BackupID'")
				}
				zid = parts[0]
				backupId = parts[1]
				zidInt, err := strconv.Atoi(zid)
				if err != nil {
					return false, fmt.Errorf(" ZID cannot convert to int ")
				}
				(*call.SdkParam)["ZID"] = zidInt
				(*call.SdkParam)["BackupID"] = backupId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				if strings.Contains(baseErr.Error(), "ErrDBNotFound") {
					logger.DebugInfo(fmt.Sprintf("error: %s.\nmsg: %s",
						baseErr.Error(), "The resource to be operated does not exist."))
					return nil
				}
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading zone on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineDnsBackupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"zid": {
				TargetField: "ZID",
			},
		},
		//NameField:    "Name",
		IdField:      "BackupID",
		CollectField: "backup_infos",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BackupID": {
				TargetField: "backup_id",
			},
			"BackupTime": {
				TargetField: "backup_time",
			},
		},
	}
}

func (s *VolcengineDnsBackupService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "DNS",
		Version:     "2018-08-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "DNS",
		Version:     "2018-08-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}
