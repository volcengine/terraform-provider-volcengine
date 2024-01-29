package rds_mssql_zone

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineZoneService struct {
	Client *ve.SdkClient
}

func NewZoneService(c *ve.SdkClient) *VolcengineZoneService {
	return &VolcengineZoneService{
		Client: c,
	}
}

func (s *VolcengineZoneService) GetClient() *ve.SdkClient {
	return s.Client
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_mssql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineZoneService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		err     error
		data    []interface{}
	)
	action := "DescribeAvailabilityZones"
	logger.Debug(logger.ReqFormat, action, condition)
	if condition == nil {
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
	} else {
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
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

func (s *VolcengineZoneService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineZoneService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineZoneService) WithResourceResponseHandlers(zone map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return zone, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineZoneService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineZoneService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {
	return callbacks
}

func (s *VolcengineZoneService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineZoneService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "ZoneName",
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

func (s *VolcengineZoneService) ReadResourceId(id string) string {
	return id
}
