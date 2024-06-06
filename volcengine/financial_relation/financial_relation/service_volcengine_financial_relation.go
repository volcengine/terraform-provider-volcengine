package financial_relation

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineFinancialRelationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewFinancialRelationService(c *ve.SdkClient) *VolcengineFinancialRelationService {
	return &VolcengineFinancialRelationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineFinancialRelationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineFinancialRelationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListFinancialRelation"

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

		results, err = ve.ObtainSdkValue("Result.List", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.List is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineFinancialRelationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")
	if len(ids) != 3 {
		return map[string]interface{}{}, fmt.Errorf("invalid financial relation id: %v", id)
	}
	subAccountId := ids[0]
	relation := ids[1]

	req := map[string]interface{}{
		"AccountIDSearchList": []string{subAccountId},
		"Relation":            []string{relation},
		"Status":              []string{"200"},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("financial_relation %s not exist ", id)
	}

	if authInfo, exist := data["AuthInfo"]; exist {
		if authArr, ok := authInfo.([]interface{}); ok && len(authArr) > 0 {
			authMap, ok := authArr[0].(map[string]interface{})
			if !ok {
				return data, errors.New("Auth info value is not map ")
			}
			data["AuthList"] = authMap["AuthList"]
		}
	}
	return data, err
}

func (s *VolcengineFinancialRelationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineFinancialRelationService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"RelationID": {
				TargetField: "relation_id",
			},
			"SubAccountID": {
				TargetField: "sub_account_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineFinancialRelationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateFinancialRelation",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"sub_account_id": {
					TargetField: "SubAccountID",
				},
				"account_alias": {
					TargetField: "AccountAlias",
				},
				"relation": {
					TargetField: "Relation",
				},
				"auth_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 处理 AuthListStr，逗号分离
				if authList, exist := d.GetOk("auth_list"); exist {
					authSet, ok := authList.(*schema.Set)
					if !ok {
						return false, fmt.Errorf(" AuthList is not set ")
					}
					authArr := make([]string, 0)
					for _, auth := range authSet.List() {
						authArr = append(authArr, strconv.Itoa(auth.(int)))
					}
					(*call.SdkParam)["AuthListStr"] = strings.Join(authArr, ",")
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// query relation id
				time.Sleep(5 * time.Second)
				financialRelation := make(map[string]interface{})
				req := map[string]interface{}{
					"AccountIDSearchList": []string{strconv.Itoa((*call.SdkParam)["SubAccountID"].(int))},
					"Relation":            []string{strconv.Itoa((*call.SdkParam)["Relation"].(int))},
					"Status":              []string{"200"},
				}
				results, err := s.ReadResources(req)
				if err != nil {
					return err
				}
				for _, v := range results {
					data, ok := v.(map[string]interface{})
					if !ok {
						return errors.New("Value is not map ")
					}
					financialRelation = data
				}
				if len(financialRelation) == 0 {
					return fmt.Errorf("financial_relation %v:%v not exist ", (*call.SdkParam)["SubAccountID"], (*call.SdkParam)["Relation"])
				}

				d.SetId(fmt.Sprintf("%v:%v:%v", (*call.SdkParam)["SubAccountID"], (*call.SdkParam)["Relation"], financialRelation["RelationID"]))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineFinancialRelationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("auth_list") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateAuth",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					// 处理 AuthListStr，逗号分离
					if authList, exist := d.GetOk("auth_list"); exist {
						authSet, ok := authList.(*schema.Set)
						if !ok {
							return false, fmt.Errorf(" AuthList is not set ")
						}
						authArr := make([]string, 0)
						for _, auth := range authSet.List() {
							authArr = append(authArr, strconv.Itoa(auth.(int)))
						}
						(*call.SdkParam)["AuthListStr"] = strings.Join(authArr, ",")
					}

					(*call.SdkParam)["RelationID"] = d.Get("relation_id")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func (s *VolcengineFinancialRelationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteFinancialRelation",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"SubAccountID": resourceData.Get("sub_account_id"),
				"Relation":     resourceData.Get("relation"),
				"RelationID":   resourceData.Get("relation_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading financial relation on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineFinancialRelationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"account_ids": {
				TargetField: "AccountIDSearchList",
				ConvertType: ve.ConvertJsonArray,
			},
			"relation": {
				TargetField: "Relation",
				ConvertType: ve.ConvertJsonArray,
			},
			"status": {
				TargetField: "Status",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		CollectField: "financial_relations",
		ResponseConverts: map[string]ve.ResponseConvert{
			"RelationID": {
				TargetField: "relation_id",
			},
			"MajorAccountID": {
				TargetField: "major_account_id",
			},
			"SubAccountID": {
				TargetField: "sub_account_id",
			},
			"AuthID": {
				TargetField: "auth_id",
			},
		},
	}
}

func (s *VolcengineFinancialRelationService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "billing",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
