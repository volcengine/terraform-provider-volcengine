package ipv6_address

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIpv6AddressService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIpv6AddressService(c *ve.SdkClient) *VolcengineIpv6AddressService {
	return &VolcengineIpv6AddressService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIpv6AddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIpv6AddressService) ReadResources(condition map[string]interface{}) (ipv6Addresses []interface{}, err error) {
	var (
		resp               *map[string]interface{}
		data               []interface{}
		results            interface{}
		next               string
		ok                 bool
		ecsInstance        map[string]interface{}
		networkInterfaces  []interface{}
		networkInterfaceId string
	)
	data, err = ve.WithNextTokenQuery(condition, "MaxResults", "NextToken", 20, nil, func(m map[string]interface{}) ([]interface{}, string, error) {
		action := "DescribeInstances"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getEcsUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getEcsUniversalInfo(action), &condition)
			if err != nil {
				return data, next, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}

		if data, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.Instances is not Slice")
		}
		return data, next, err
	})

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return ipv6Addresses, nil
	}
	if ecsInstance, ok = data[0].(map[string]interface{}); !ok {
		return ipv6Addresses, errors.New("Value is not map ")
	} else {
		// query primary network interface info of the ecs instance
		if networkInterfaces, ok = ecsInstance["NetworkInterfaces"].([]interface{}); !ok {
			return ipv6Addresses, errors.New("Instances.NetworkInterfaces is not Slice")
		}
		for _, networkInterface := range networkInterfaces {
			if networkInterfaceMap, ok := networkInterface.(map[string]interface{}); ok &&
				networkInterfaceMap["Type"] == "primary" {
				networkInterfaceId = networkInterfaceMap["NetworkInterfaceId"].(string)
			}
		}

		action := "DescribeNetworkInterfaces"
		req := map[string]interface{}{
			"NetworkInterfaceIds.1": networkInterfaceId,
		}
		logger.Debug(logger.ReqFormat, action, req)
		res, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
		if err != nil {
			logger.Info("DescribeNetworkInterfaces error:", err)
			return ipv6Addresses, err
		}
		logger.Debug(logger.RespFormat, action, condition, *res)

		networkInterfaceInfos, err := ve.ObtainSdkValue("Result.NetworkInterfaceSets", *res)
		if err != nil {
			logger.Info("ObtainSdkValue Result.NetworkInterfaceSets error:", err)
			return ipv6Addresses, err
		}
		if ipv6Sets, ok := networkInterfaceInfos.([]interface{})[0].(map[string]interface{})["IPv6Sets"].([]interface{}); ok {
			for _, ipv6Address := range ipv6Sets {
				ipv6AddressMap := make(map[string]interface{})
				ipv6AddressMap["Ipv6Address"] = ipv6Address
				ipv6Addresses = append(ipv6Addresses, ipv6AddressMap)
			}
		}
	}

	return ipv6Addresses, err
}

func (s *VolcengineIpv6AddressService) ReadResource(resourceData *schema.ResourceData, allocationId string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineIpv6AddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineIpv6AddressService) WithResourceResponseHandlers(ipv6Address map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return ipv6Address, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineIpv6AddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIpv6AddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIpv6AddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineIpv6AddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"associated_instance_id": {
				TargetField: "InstanceIds.1",
			},
		},
		//IdField:      "AllocationId",
		CollectField: "ipv6_addresses",
	}
}

func (s *VolcengineIpv6AddressService) ReadResourceId(id string) string {
	return id
}

func getEcsUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
