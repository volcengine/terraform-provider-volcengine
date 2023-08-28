package cloudfs_ns_quota

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudfsNsQuotaService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineCloudfsNsQuotaService {
	return &VolcengineCloudfsNsQuotaService{
		Client: c,
	}
}

func (s *VolcengineCloudfsNsQuotaService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudfsNsQuotaService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "GetNsQuota"
		logger.Debug(logger.ReqFormat, action, condition)

		fs, ok := condition["FsNames"]
		if !ok {
			return data, nil
		}
		for _, fsName := range fs.([]interface{}) {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
				"FsName": fsName,
			})
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, resp)
			res, err := ve.ObtainSdkValue("Result", *resp)
			if err != nil {
				return data, err
			}
			data = append(data, res)
		}
		return data, nil
	})
}

func (s *VolcengineCloudfsNsQuotaService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineCloudfsNsQuotaService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineCloudfsNsQuotaService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineCloudfsNsQuotaService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}

}

func (s *VolcengineCloudfsNsQuotaService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCloudfsNsQuotaService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCloudfsNsQuotaService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"fs_names": {
				TargetField: "FsNames",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		CollectField: "quotas",
	}
}

func (s *VolcengineCloudfsNsQuotaService) ReadResourceId(id string) string {
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
