package spec

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoSpecService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewSpecService(c *ve.SdkClient) *VolcengineMongoSpecService {
	return &VolcengineMongoSpecService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineMongoSpecService) GetClient() *ve.SdkClient {
	return s.Client
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "mongodb",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineMongoSpecService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		err     error
		data    []interface{}
	)
	action := "DescribeNodeSpecs"
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

	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = make([]interface{}, 0)
	}
	data = append(data, results)
	return data, nil
}

func (s *VolcengineMongoSpecService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineMongoSpecService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineMongoSpecService) WithResourceResponseHandlers(zone map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineMongoSpecService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoSpecService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {
	return callbacks
}

func (s *VolcengineMongoSpecService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoSpecService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "specs",
	}
}

func (s *VolcengineMongoSpecService) ReadResourceId(id string) string {
	return id
}
