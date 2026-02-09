package iam_access_key_last_used

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamAccessKeyLastUsedService struct {
	Client *ve.SdkClient
}

func NewIamAccessKeyLastUsedService(c *ve.SdkClient) *VolcengineIamAccessKeyLastUsedService {
	return &VolcengineIamAccessKeyLastUsedService{
		Client: c,
	}
}

func (s *VolcengineIamAccessKeyLastUsedService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamAccessKeyLastUsedService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	action := "GetAccessKeyLastUsed"
	logger.Debug(logger.ReqFormat, action, m)
	if m == nil {
		return data, errors.New("missing params")
	}

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}

	result, err := ve.ObtainSdkValue("Result.AccessKeyLastUsed", *resp)
	if err != nil {
		return data, err
	}
	if result == nil {
		return []interface{}{}, nil
	}

	return []interface{}{result}, nil
}

func (s *VolcengineIamAccessKeyLastUsedService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"access_key_id": {TargetField: "AccessKeyId"},
			"user_name":     {TargetField: "UserName"},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"Region": {
				TargetField: "region",
			},
			"Service": {
				TargetField: "service",
			},
			"RequestTime": {
				TargetField: "request_time",
			},
		},
		CollectField: "access_key_last_useds",
	}
}

func (s *VolcengineIamAccessKeyLastUsedService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineIamAccessKeyLastUsedService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamAccessKeyLastUsedService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineIamAccessKeyLastUsedService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIamAccessKeyLastUsedService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIamAccessKeyLastUsedService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIamAccessKeyLastUsedService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}
