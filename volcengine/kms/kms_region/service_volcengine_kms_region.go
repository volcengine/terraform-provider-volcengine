package kms_region

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsRegionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsRegionService(c *ve.SdkClient) *VolcengineKmsRegionService {
	return &VolcengineKmsRegionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsRegionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsRegionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "DescribeRegions"
	bytes, _ := json.Marshal(m)
	logger.Debug(logger.ReqFormat, action, string(bytes))

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
	if err != nil {
		return data, err
	}

	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, m, string(respBytes))

	results, err = ve.ObtainSdkValue("Result.Regions", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		return []interface{}{}, nil
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.Regions is not Slice")
	}
	return data, nil
}

func (s *VolcengineKmsRegionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsRegionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return nil, "", err
		},
	}
}

func (s *VolcengineKmsRegionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (VolcengineKmsRegionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsRegionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineKmsRegionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineKmsRegionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "RegionId",
		IdField:      "RegionId",
		CollectField: "regions",
	}
}

func (s *VolcengineKmsRegionService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
