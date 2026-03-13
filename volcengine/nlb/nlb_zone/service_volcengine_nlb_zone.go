package nlb_zone

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbZoneService struct {
	Client *ve.SdkClient
}

func NewNlbZoneService(c *ve.SdkClient) *VolcengineNlbZoneService {
	return &VolcengineNlbZoneService{
		Client: c,
	}
}

func (s *VolcengineNlbZoneService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbZoneService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "DescribeNLBZones"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return nil, err
	}
	results, err = ve.ObtainSdkValue("Result.Zones", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return nil, errors.New("Result.Zones is not Slice")
	}
	return data, nil
}

func (s *VolcengineNlbZoneService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineNlbZoneService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineNlbZoneService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (VolcengineNlbZoneService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineNlbZoneService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbZoneService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbZoneService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "ZoneId",
		CollectField: "zones",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ZoneId": {
				TargetField: "zone_id",
			},
		},
	}
}

func (s *VolcengineNlbZoneService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Regional,
	}
}
