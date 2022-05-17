package zone

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackZoneService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewZoneService(c *ve.SdkClient) *VestackZoneService {
	return &VestackZoneService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackZoneService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackZoneService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		err     error
		data    []interface{}
	)
	ecs := s.Client.EcsClient
	action := "DescribeZones"
	logger.Debug(logger.ReqFormat, action, condition)
	if condition == nil {
		resp, err = ecs.DescribeZonesCommon(nil)
	} else {
		resp, err = ecs.DescribeZonesCommon(&condition)
	}
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.RespFormat, action, condition, *resp)

	results, err = ve.ObtainSdkValue("Result.Zones", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = make([]interface{}, 0)
	}

	if data, ok = results.([]interface{}); !ok {
		return nil, errors.New("Result.Zones is not Slice")
	}

	return data, nil
}

func (s *VestackZoneService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VestackZoneService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VestackZoneService) WithResourceResponseHandlers(zone map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return zone, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VestackZoneService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VestackZoneService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {
	return callbacks
}

func (s *VestackZoneService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VestackZoneService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ZoneIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "ZoneId",
		IdField:      "ZoneId",
		CollectField: "zones",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ZoneId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VestackZoneService) ReadResourceId(id string) string {
	return id
}
