package instance_parameter_log

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoDBInstanceParameterLogService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewMongoDBInstanceParameterLogService(c *ve.SdkClient) *VolcengineMongoDBInstanceParameterLogService {
	return &VolcengineMongoDBInstanceParameterLogService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineMongoDBInstanceParameterLogService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongoDBInstanceParameterLogService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 100, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceParametersLog"
		logger.Debug(logger.ReqFormat, action, condition)
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
		logger.Debug(logger.RespFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("Result.ParameterChangeLog", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		data, ok = results.([]interface{})
		if !ok {
			return data, fmt.Errorf("DescribeDBInstanceParametersLog response is not a slice")
		}
		return data, nil
	})
}

func (s *VolcengineMongoDBInstanceParameterLogService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineMongoDBInstanceParameterLogService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineMongoDBInstanceParameterLogService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBInstanceParameterLogService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoDBInstanceParameterLogService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoDBInstanceParameterLogService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoDBInstanceParameterLogService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "parameter_change_logs",
		ContentType:  ve.ContentTypeJson,
	}
}

func (s *VolcengineMongoDBInstanceParameterLogService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "mongodb",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
