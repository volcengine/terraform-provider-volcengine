package security_group_rule

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackSecurityGroupRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

type Direction string

const (
	DirectionIngress = Direction("ingress")
	DirectionEgress  = Direction("egress")
)

func NewSecurityGroupRuleService(c *ve.SdkClient) *VestackSecurityGroupRuleService {
	return &VestackSecurityGroupRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackSecurityGroupRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (VestackSecurityGroupRuleService) WithResourceResponseHandlers(rule map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return rule, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackSecurityGroupRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var action string
	direction := resourceData.Get("direction").(string)
	if direction == string(DirectionEgress) {
		action = "AuthorizeSecurityGroupEgress"
	} else {
		action = "AuthorizeSecurityGroupIngress"
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if resourceData.Get("cidr_ip") == "" && resourceData.Get("source_group_id") == "" {
					return false, fmt.Errorf("At least one of cidr_ip and source_group_id exists. ")
				}
				protocol := resourceData.Get("protocol").(string)
				start := resourceData.Get("port_start").(int)
				end := resourceData.Get("port_end").(int)
				if err := validateProtocol(protocol, start, end); err != nil {
					return false, err
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				if direction == string(DirectionEgress) {
					return s.Client.VpcClient.AuthorizeSecurityGroupEgressCommon(call.SdkParam)
				} else {
					return s.Client.VpcClient.AuthorizeSecurityGroupIngressCommon(call.SdkParam)
				}
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var (
					securityGroupId = resourceData.Get("security_group_id").(string)
					cidrIp          = resourceData.Get("cidr_ip").(string)
					protocol        = resourceData.Get("protocol").(string)
					portStart       = resourceData.Get("port_start").(int)
					portEnd         = resourceData.Get("port_end").(int)
				)
				id, _ := ve.ObtainSdkValue("Result.RuleId", securityGroupId+":"+protocol+":"+strconv.Itoa(portStart)+":"+strconv.Itoa(portEnd)+":"+cidrIp)
				d.SetId(id.(string))
				return nil
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VestackSecurityGroupRuleService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VestackSecurityGroupRuleService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VestackSecurityGroupRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VestackSecurityGroupRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var action string
	direction := resourceData.Get("direction").(string)
	if direction == string(DirectionEgress) {
		action = "ModifySecurityGroupRuleDescriptionsEgress"
	} else {
		action = "ModifySecurityGroupRuleDescriptionsIngress"
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				items := strings.Split(d.Id(), ":")
				if len(items) != 5 {
					return false, fmt.Errorf("import id must be of the form SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp")
				}

				securityGroupId := items[0]
				protocol := items[1]
				portStart := items[2]
				portEnd := items[3]
				cidrIp := items[4]
				ruleId := securityGroupId + ":" + protocol + ":" + portStart + ":" + portEnd + ":" + cidrIp

				(*call.SdkParam)["Protocol"] = protocol
				(*call.SdkParam)["PortStart"] = portStart
				(*call.SdkParam)["PortEnd"] = portEnd
				(*call.SdkParam)["CidrIp"] = cidrIp
				(*call.SdkParam)["SecurityGroupId"] = securityGroupId
				(*call.SdkParam)["RuleId"] = ruleId

				start, _ := strconv.Atoi((*call.SdkParam)["PortStart"].(string))
				end, _ := strconv.Atoi((*call.SdkParam)["PortEnd"].(string))

				if err := validateProtocol(protocol, start, end); err != nil {
					return false, err
				}

				d.SetId(ruleId)

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				if direction == string(DirectionEgress) {
					return s.Client.VpcClient.ModifySecurityGroupRuleDescriptionsEgressCommon(call.SdkParam)
				} else {
					return s.Client.VpcClient.ModifySecurityGroupRuleDescriptionsIngressCommon(call.SdkParam)
				}

			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VestackSecurityGroupRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var action string
	direction := resourceData.Get("direction").(string)
	if direction == string(DirectionEgress) {
		action = "RevokeSecurityGroupEgress"
	} else {
		action = "RevokeSecurityGroupIngress"
	}

	items := strings.Split(resourceData.Id(), ":")
	if len(items) != 5 {
		return []ve.Callback{}
	}

	securityGroupId := items[0]
	protocol := items[1]
	portStart := items[2]
	portEnd := items[3]
	cidrIp := items[4]

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"RuleId":          resourceData.Id(),
				"PortStart":       portStart,
				"PortEnd":         portEnd,
				"Protocol":        protocol,
				"SecurityGroupId": securityGroupId,
				"CidrIp":          cidrIp,
				"Priority":        resourceData.Get("priority"),
				"Policy":          resourceData.Get("policy"),
				"SourceGroupId":   resourceData.Get("source_group_id"),
			},

			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				if direction == string(DirectionEgress) {
					return s.Client.VpcClient.RevokeSecurityGroupEgressCommon(call.SdkParam)
				} else {
					return s.Client.VpcClient.RevokeSecurityGroupIngressCommon(call.SdkParam)
				}

			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VestackSecurityGroupRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func importSecurityGroupRule(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(d.Id(), ":")
	if len(items) != 5 {
		return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp")
	}
	err = d.Set("security_group_id", items[0])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}
	err = d.Set("protocol", items[1])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	err = d.Set("port_start", items[2])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	err = d.Set("port_end", items[3])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	err = d.Set("cidr_ip", items[4])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	return []*schema.ResourceData{d}, nil
}

func (s *VestackSecurityGroupRuleService) ReadResourceId(id string) string {
	items := strings.Split(id, ":")
	return items[0]
}

func validateProtocol(protocol string, start, end int) error {
	switch protocol {
	case "tcp":
		if start < 1 || end < 1 {
			return fmt.Errorf("Protocol is tcp,Port start or end must between 1-65535. ")
		}
	case "udp":
		if start < 1 || end < 1 {
			return fmt.Errorf("Protocol is udp,Port start or end must between 1-65535. ")
		}
	case "icmp":
		if start != -1 || end != -1 {
			return fmt.Errorf("Protocol is icmp,Port start or end must -1. ")
		}
	case "all":
		if start != -1 || end != -1 {
			return fmt.Errorf("Protocol is all,Port start or end must -1. ")
		}
	}
	return nil
}
