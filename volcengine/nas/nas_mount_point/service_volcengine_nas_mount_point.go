package nas_mount_point

import (
	"errors"
	"fmt"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_file_system"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/subnet"
)

type VolcengineNasMountPointService struct {
	Client *ve.SdkClient
}

func (v *VolcengineNasMountPointService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineNasMountPointService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeMountPoints"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
		} else {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)
		if err != nil {
			return data, err
		}
		results, err = ve.ObtainSdkValue("Result.MountPoints", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.MountPoints is not slice")
		}
		for index, ele := range data {
			mountPoint, ok := ele.(map[string]interface{})
			if !ok {
				return data, errors.New("Result.MountPoint is not map")
			}
			// 保证import可以读到permission group id
			if permissionGroup, ok := mountPoint["PermissionGroup"].(map[string]interface{}); ok {
				if permissionId, ok := permissionGroup["PermissionGroupId"]; ok {
					data[index].(map[string]interface{})["PermissionGroupId"] = permissionId
				}
			}
		}
		return data, err
	})
}

func (v *VolcengineNasMountPointService) ReadResource(resData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return nil, errors.New("id form error")
	}
	req := map[string]interface{}{
		"FileSystemId": ids[0],
		"MountPointId": ids[1],
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("nas mount point %s not found", ids[1])
	}
	return data, err
}

func (v *VolcengineNasMountPointService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			demo, err = v.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("nas mount point status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (v *VolcengineNasMountPointService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineNasMountPointService) CreateResource(resData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateMountPoint",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			LockId: func(d *schema.ResourceData) string {
				return d.Get("file_system_id").(string)
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				subnetId, _ := d.Get("subnet_id").(string)
				resp, err := subnet.NewSubnetService(v.Client).ReadResource(resData, subnetId)
				if err != nil {
					return false, err
				}
				vpcId := resp["VpcId"]
				(*call.SdkParam)["VpcId"] = vpcId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("Result.MountPointId", *resp)
				if err != nil {
					return err
				}
				fileSystemId, _ := d.Get("file_system_id").(string)
				d.SetId(fmt.Sprint(fileSystemId, ":", id))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				nas_file_system.NewNasFileSystemService(v.Client): {
					Target:     []string{"Running"},
					Timeout:    resData.Timeout(schema.TimeoutCreate),
					ResourceId: resData.Get("file_system_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineNasMountPointService) ModifyResource(resData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	ids := strings.Split(resData.Id(), ":")
	fileSystemId := ids[0]
	mountPointId := ids[1]
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateMountPoint",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			LockId: func(d *schema.ResourceData) string {
				return d.Get("file_system_id").(string)
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["FileSystemId"] = fileSystemId
				(*call.SdkParam)["MountPointId"] = mountPointId
				return true, nil
			},
			Convert: map[string]ve.RequestConvert{
				"permission_group_id": {
					ConvertType: ve.ConvertDefault,
				},
				"mount_point_name": {
					ConvertType: ve.ConvertDefault,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineNasMountPointService) RemoveResource(resData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resData.Id(), ":")
	fileSystemId := ids[0]
	mountPointId := ids[1]
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteMountPoint",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			LockId: func(d *schema.ResourceData) string {
				return d.Get("file_system_id").(string)
			},
			SdkParam: &map[string]interface{}{
				"FileSystemId": fileSystemId,
				"MountPointId": mountPointId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading nas mount point on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 10*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineNasMountPointService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		IdField:      "MountPointId",
		NameField:    "MountPointName",
		CollectField: "mount_points",
	}
}

func (v *VolcengineNasMountPointService) ReadResourceId(s string) string {
	return s
}

func NewVolcengineNasMountPointService(c *ve.SdkClient) *VolcengineNasMountPointService {
	return &VolcengineNasMountPointService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "FileNAS",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
