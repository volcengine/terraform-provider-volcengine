package cr_authorization_token

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCrAuthorizationTokenService struct {
	Client     *ve.SdkClient
}

func NewCrAuthorizationTokenService(c *ve.SdkClient) *VolcengineCrAuthorizationTokenService {
	return &VolcengineCrAuthorizationTokenService{
		Client:     c,
	}
}

func (s *VolcengineCrAuthorizationTokenService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCrAuthorizationTokenService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)

	action := "GetAuthorizationToken"
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

	logger.Debug(logger.RespFormat, action, resp)
	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		return data, fmt.Errorf("GetAuthorizationToken return an empty result")
	}

	token, err := ve.ObtainSdkValue("Result.Token", *resp)
	if err != nil {
		return data, err
	}
	username, err := ve.ObtainSdkValue("Result.Username", *resp)
	if err != nil {
		return data, err
	}
	expireTime, err := ve.ObtainSdkValue("Result.ExpireTime", *resp)
	if err != nil {
		return data, err
	}

	user := map[string]interface{}{
		"Token":      token,
		"Username":   username,
		"ExpireTime": expireTime,
	}

	return []interface{}{user}, err
}

func (s *VolcengineCrAuthorizationTokenService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineCrAuthorizationTokenService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineCrAuthorizationTokenService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCrAuthorizationTokenService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCrAuthorizationTokenService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return s.CreateResource(resourceData, resource)
}

func (s *VolcengineCrAuthorizationTokenService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCrAuthorizationTokenService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		CollectField: "tokens",
	}
}

func (s *VolcengineCrAuthorizationTokenService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cr",
		Version:     "2022-05-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}