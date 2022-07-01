package vpc

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVpcService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVpcService(c *ve.SdkClient) *VolcengineVpcService {
	return &VolcengineVpcService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVpcService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVpcService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		vpcClient := s.Client.VpcClient
		action := "DescribeVpcs"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = vpcClient.DescribeVpcsCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = vpcClient.DescribeVpcsCommon(&condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.Vpcs", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Vpcs is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVpcService) ReadResource(resourceData *schema.ResourceData, vpcId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if vpcId == "" {
		vpcId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"VpcIds.1": vpcId,
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
		return data, fmt.Errorf("Vpc %s not exist ", vpcId)
	}
	return data, err
}

func (s *VolcengineVpcService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("Vpc  status  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineVpcService) WithResourceResponseHandlers(vpc map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return vpc, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineVpcService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateVpc",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"dns_servers": {
					TargetField: "DnsServers",
					ConvertType: ve.ConvertWithN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建vpc
				return s.Client.VpcClient.CreateVpcCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.VpcId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineVpcService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyVpcAttributes",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"dns_servers": {
					TargetField: "DnsServers",
					ConvertType: ve.ConvertWithN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["VpcId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改vpc属性
				return s.Client.VpcClient.ModifyVpcAttributesCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVpcService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteVpc",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VpcId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除VPC
				return s.Client.VpcClient.DeleteVpcCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading vpc on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVpcService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "VpcIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "VpcName",
		IdField:      "VpcId",
		CollectField: "vpcs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"VpcId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineVpcService) ReadResourceId(id string) string {
	return id
}
