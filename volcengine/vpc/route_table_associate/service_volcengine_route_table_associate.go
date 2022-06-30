package route_table_associate

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

type VolcengineRouteTableAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRouteTableAssociateService(c *ve.SdkClient) *VolcengineRouteTableAssociateService {
	return &VolcengineRouteTableAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRouteTableAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRouteTableAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		vpcClient := s.Client.VpcClient
		action := "DescribeRouteTableList"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = vpcClient.DescribeRouteTableListCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = vpcClient.DescribeRouteTableListCommon(&condition)
			if err != nil {
				return data, err
			}
		}

		results, err = ve.ObtainSdkValue("Result.RouterTableList", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RouterTableList is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRouteTableAssociateService) ReadResource(resourceData *schema.ResourceData, associateId string) (data map[string]interface{}, err error) {
	var (
		results        []interface{}
		ok             bool
		associate      bool
		subnetIds      interface{}
		tmpSubnetIds   []interface{}
		routeTableId   string
		targetSubnetId string
		ids            []string
	)

	if associateId == "" {
		associateId = s.ReadResourceId(resourceData.Id())
	}

	ids = strings.Split(associateId, ":")
	routeTableId = ids[0]
	targetSubnetId = ids[1]

	req := map[string]interface{}{
		"RouteTableId": routeTableId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("route table %s not exist ", routeTableId)
	}
	subnetIds, err = ve.ObtainSdkValue("SubnetIds", data)
	if err != nil {
		return data, err
	}
	if subnetIds == nil {
		return data, errors.New("not associate")
	}
	tmpSubnetIds, ok = subnetIds.([]interface{})
	if !ok {
		return data, errors.New("subnet ids is not string slice")
	}
	for _, subnetId := range tmpSubnetIds {
		if subnetId.(string) == targetSubnetId {
			associate = true
			break
		}
	}
	if !associate {
		return data, errors.New("not associate")
	}
	return data, err
}

func (s *VolcengineRouteTableAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status = "Associate"
			)
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				if !strings.Contains(err.Error(), "not associate") {
					return nil, "", err
				}
				status = "Available"
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status, nil
		},
	}
}

func (VolcengineRouteTableAssociateService) WithResourceResponseHandlers(routeTables map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return routeTables, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineRouteTableAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateRouteTable",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.AssociateRouteTableCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["RouteTableId"], ":", (*call.SdkParam)["SubnetId"]))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Associate"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRouteTableAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineRouteTableAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(s.ReadResourceId(resourceData.Id()), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateRouteTable",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"RouteTableId": ids[0],
				"SubnetId":     ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.DisassociateRouteTableCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutDelete),
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading route table associate on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRouteTableAssociateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRouteTableAssociateService) ReadResourceId(id string) string {
	return id
}
