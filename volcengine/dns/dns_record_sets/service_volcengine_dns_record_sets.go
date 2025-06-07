package dns_record_sets

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineDnsRecordSetsService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewDnsRecordSetsService(c *ve.SdkClient) *VolcengineDnsRecordSetsService {
	return &VolcengineDnsRecordSetsService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineDnsRecordSetsService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineDnsRecordSetsService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRecordSets"

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
		results, err = ve.ObtainSdkValue("Result.RecordSets", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RecordSets is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineDnsRecordSetsService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineDnsRecordSetsService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return nil, "", err
		},
	}
}

func (s *VolcengineDnsRecordSetsService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineDnsRecordSetsService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineDnsRecordSetsService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineDnsRecordSetsService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineDnsRecordSetsService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"zid": {
				TargetField: "ZID",
			},
			"record_set_id": {
				TargetField: "RecordSetID",
			},
			"host": {
				TargetField: "Host",
			},
		},
		NameField:    "PQDN",
		IdField:      "ID",
		CollectField: "record_sets",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "id",
			},
			"PQDN": {
				TargetField: "pqdn",
			},
			"Host": {
				TargetField: "host",
			},
			"Type": {
				TargetField: "type",
			},
			"Line": {
				TargetField: "line",
			},
		},
	}
}

func (s *VolcengineDnsRecordSetsService) ReadResourceId(id string) string {
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
