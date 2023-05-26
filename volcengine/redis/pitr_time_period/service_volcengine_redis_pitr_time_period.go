package pitr_time_period

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRedisPitrTimePeriodService struct {
	Client *ve.SdkClient
}

func (v *VolcengineRedisPitrTimePeriodService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineRedisPitrTimePeriodService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		ids       []interface{}
		resp      *map[string]interface{}
		result    interface{}
		results   []interface{}
		resultMap map[string]interface{}
	)
	action := "DescribePitrTimeWindow"
	ids = m["Ids"].(*schema.Set).List()
	if len(ids) == 0 {
		return data, nil
	}
	for _, id := range ids {
		instanceId, ok := id.(string)
		if !ok {
			return data, errors.New("err instance id")
		}
		req := map[string]interface{}{
			"InstanceId": instanceId,
		}
		logger.Debug(logger.ReqFormat, action, req)
		resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, action, req, *resp)
		result, err = ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if resultMap, ok = result.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
		// 加个ID，方便对照
		resultMap["InstanceId"] = instanceId
		results = append(results, resultMap)
	}
	return results, nil
}

func (v *VolcengineRedisPitrTimePeriodService) ReadResource(data *schema.ResourceData, s string) (map[string]interface{}, error) {
	return nil, nil
}

func (v *VolcengineRedisPitrTimePeriodService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return nil
}

func (v *VolcengineRedisPitrTimePeriodService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineRedisPitrTimePeriodService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineRedisPitrTimePeriodService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineRedisPitrTimePeriodService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineRedisPitrTimePeriodService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		CollectField: "periods",
	}
}

func (v *VolcengineRedisPitrTimePeriodService) ReadResourceId(s string) string {
	return s
}

func NewVolcengineRedisPitrTimeWindowService(c *ve.SdkClient) *VolcengineRedisPitrTimePeriodService {
	return &VolcengineRedisPitrTimePeriodService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Redis",
		Action:      actionName,
		Version:     "2020-12-07",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
