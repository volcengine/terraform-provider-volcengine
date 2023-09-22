package ecs_key_pair_associate

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEcsKeyPairAssociateService struct {
	Client *ve.SdkClient
}

func NewEcsKeyPairAssociateService(c *ve.SdkClient) *VolcengineEcsKeyPairAssociateService {
	return &VolcengineEcsKeyPairAssociateService{
		Client: c,
	}
}

func (s *VolcengineEcsKeyPairAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsKeyPairAssociateService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	return nil, nil
}

func (s *VolcengineEcsKeyPairAssociateService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		resp             *map[string]interface{}
		results          interface{}
		ok               bool
		keyPairId        string
		targetInstanceId string
		ids              []string
		dep              []interface{}
	)

	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}

	ids = strings.Split(tmpId, ":")
	keyPairId = ids[0]
	targetInstanceId = ids[1]

	req := map[string]interface{}{
		"KeyPairIds.1": keyPairId,
	}
	client := s.Client.UniversalClient
	action := "DescribeKeyPairs"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = client.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}

	results, err = ve.ObtainSdkValue("Result.KeyPairs", *resp)
	if err != nil {
		return data, err
	}
	if dep, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.KeyPairs is not Slice")
	}
	if len(dep) == 0 {
		return data, fmt.Errorf("Ecs KeyPairs %s not exist ", keyPairId)
	}
	keyPairName := dep[0].(map[string]interface{})["KeyPairName"].(string)

	insReq := map[string]interface{}{
		"InstanceIds.1": targetInstanceId,
	}
	action = "DescribeInstances"
	logger.Debug(logger.ReqFormat, action, insReq)
	resp, err = client.DoCall(getUniversalInfo(action), &insReq)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Result.Instances", *resp)
	if err != nil {
		return data, err
	}
	if dep, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.Instances is not Slice")
	}
	if len(dep) == 0 {
		return data, fmt.Errorf("Ecs Instances %s not exist ", keyPairId)
	}
	if dep[0].(map[string]interface{})["KeyPairName"].(string) != keyPairName {
		return data, errors.New("not associate")
	}

	data = make(map[string]interface{})
	data["KeyPairId"] = keyPairId
	data["InstanceId"] = targetInstanceId
	return data, err

}

func (s *VolcengineEcsKeyPairAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data map[string]interface{}
			)

			if err = resource.Retry(3*time.Minute, func() *resource.RetryError {
				data, err = s.ReadResource(resourceData, id)
				if err != nil {
					return resource.NonRetryableError(err)
				}
				if len(data) == 0 {
					return resource.RetryableError(fmt.Errorf("Retry "))
				}
				return nil
			}); err != nil {
				return nil, "error", err
			}
			return data, "success", err
		},
	}
}

func (s *VolcengineEcsKeyPairAssociateService) WithResourceResponseHandlers(deploymentSet map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return deploymentSet, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsKeyPairAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachKeyPair",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"KeyPairId":     resourceData.Get("key_pair_id"),
				"InstanceIds.1": resourceData.Get("instance_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				if err != nil {
					return resp, err
				}
				logger.Debug(logger.RespFormat, call.Action, *resp)
				operationDetails, err := ve.ObtainSdkValue("Result.OperationDetails", *resp)
				if err != nil {
					return resp, fmt.Errorf("get Result.OperationDetails failed")
				}
				if _, ok := operationDetails.([]interface{}); !ok || len(operationDetails.([]interface{})) == 0 {
					return resp, fmt.Errorf("get Result.OperationDetails is not a valid slice")
				}
				operationDetail, ok := operationDetails.([]interface{})[0].(map[string]interface{})
				if !ok {
					return resp, fmt.Errorf("operation detail is not a map")
				}
				errStruct, ok := operationDetail["Error"]
				if !ok || errStruct == nil {
					return resp, nil
				}
				errMap, ok := operationDetail["Error"].(map[string]interface{})
				if !ok {
					return resp, fmt.Errorf("operation detail Error is not a map")
				}
				code := errMap["Code"].(string)
				message := errMap["Message"].(string)
				return resp, fmt.Errorf(code + ": " + message)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)["KeyPairId"].(string) + ":" + (*call.SdkParam)["InstanceIds.1"].(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsKeyPairAssociateService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEcsKeyPairAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(s.ReadResourceId(resourceData.Id()), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachKeyPair",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"KeyPairId":     ids[0],
				"InstanceIds.1": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading key pair associate on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsKeyPairAssociateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineEcsKeyPairAssociateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
