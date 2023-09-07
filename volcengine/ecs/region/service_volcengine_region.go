package region

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRegionService struct {
	Client *ve.SdkClient
}

func (v *VolcengineRegionService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineRegionService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "DescribeRegions"
	logger.Debug(logger.ReqFormat, action, condition)
	if condition == nil {
		resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
	} else {
		resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
	}
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.RespFormat, action, condition, *resp)

	results, err = ve.ObtainSdkValue("Result.Regions", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = make([]interface{}, 0)
	}

	if data, ok = results.([]interface{}); !ok {
		return nil, errors.New("Result.Regions is not Slice")
	}

	return data, nil
}

func (v *VolcengineRegionService) ReadResource(data *schema.ResourceData, s string) (map[string]interface{}, error) {
	return nil, nil
}

func (v *VolcengineRegionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (v *VolcengineRegionService) WithResourceResponseHandlers(region map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return region, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineRegionService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineRegionService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineRegionService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineRegionService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "RegionIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "RegionId",
		IdField:      "RegionId",
		CollectField: "regions",
		ResponseConverts: map[string]ve.ResponseConvert{
			"RegionId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (v *VolcengineRegionService) ReadResourceId(s string) string {
	return s
}

func NewRegionService(c *ve.SdkClient) *VolcengineRegionService {
	return &VolcengineRegionService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
