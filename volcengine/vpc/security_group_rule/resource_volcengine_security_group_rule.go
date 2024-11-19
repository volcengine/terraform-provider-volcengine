package security_group_rule

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
SecurityGroupRule can be imported using the id, e.g.
```
$ terraform import volcengine_security_group_rule.default ID is a string concatenated with colons(SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp:SourceGroupId:Direction:Policy:Priority)
```

*/

func ResourceVolcengineSecurityGroupRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineSecurityGroupRuleCreate,
		Read:   resourceVolcengineSecurityGroupRuleRead,
		Update: resourceVolcengineSecurityGroupRuleUpdate,
		Delete: resourceVolcengineSecurityGroupRuleDelete,
		Importer: &schema.ResourceImporter{
			State: importSecurityGroupRule,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"direction": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ingress",
					"egress",
				}, false),
				Description: "Direction of rule, ingress (inbound) or egress (outbound).",
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"tcp",
					"udp",
					"icmp",
					"all",
					"icmpv6",
				}, false),
				Description: "Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.",
			},
			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of SecurityGroup.",
			},
			"port_start": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Port start of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.",
			},
			"port_end": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Port end of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.",
			},
			"cidr_ip": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"source_group_id"},
				Description:   "Cidr ip of egress/ingress Rule.",
			},
			"source_group_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cidr_ip"},
				Description:   "ID of the source security group whose access permission you want to set.",
			},
			"policy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"accept",
					"drop",
				}, false),
				Default:     "accept",
				Description: "Access strategy. Valid values: `accept`, `drop`. Default is `accept`.",
			},
			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(1, 100),
				Description:  "Priority of a security group rule. Valid value range: 1~100. Default is 1.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "description of a egress rule.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of SecurityGroup.",
			},
		},
	}
}

func importSecurityGroupRule(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	var (
		err    error
		cidrIp string
	)
	items := strings.Split(d.Id(), ":")
	itemsLength := len(items)
	if itemsLength < 9 {
		return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form " +
			"SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp:SourceGroupId:Direction:Policy:Priority")
	}
	err = d.Set("security_group_id", items[0])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}
	err = d.Set("protocol", items[1])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	if len(items[2]) > 0 {
		ps, err := strconv.Atoi(items[2])
		if err != nil {
			return []*schema.ResourceData{d}, err
		}
		err = d.Set("port_start", ps)
		if err != nil {
			return []*schema.ResourceData{d}, err
		}
	}

	if len(items[3]) > 0 {
		pn, err := strconv.Atoi(items[3])
		if err != nil {
			return []*schema.ResourceData{d}, err
		}
		err = d.Set("port_end", pn)
		if err != nil {
			return []*schema.ResourceData{d}, err
		}
	}
	if itemsLength == 9 {
		// ipv4
		cidrIp = items[4]
	} else {
		// ipv6
		strArr := make([]string, 0)
		for i := 4; i < itemsLength-4; i++ {
			strArr = append(strArr, items[i])
		}
		cidrIp = strings.Join(strArr, ":")
	}
	err = d.Set("cidr_ip", cidrIp)
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	err = d.Set("source_group_id", items[itemsLength-4])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	err = d.Set("direction", items[itemsLength-3])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	err = d.Set("policy", items[itemsLength-2])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	if len(items[itemsLength-1]) > 0 {
		pr, err := strconv.Atoi(items[itemsLength-1])
		if err != nil {
			return []*schema.ResourceData{d}, err
		}
		err = d.Set("priority", pr)
		if err != nil {
			return []*schema.ResourceData{d}, err
		}
	}
	return []*schema.ResourceData{d}, nil
}

func resourceVolcengineSecurityGroupRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupRuleService := NewSecurityGroupRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(securityGroupRuleService, d, ResourceVolcengineSecurityGroupRule())
	if err != nil {
		return fmt.Errorf("error on creating securityGroupRuleService  %q, %w", d.Id(), err)
	}
	return resourceVolcengineSecurityGroupRuleRead(d, meta)
}

func resourceVolcengineSecurityGroupRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupRuleService := NewSecurityGroupRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(securityGroupRuleService, d, ResourceVolcengineSecurityGroupRule())
	if err != nil {
		return fmt.Errorf("error on reading securityGroupRuleService %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineSecurityGroupRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupRuleService := NewSecurityGroupRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(securityGroupRuleService, d, ResourceVolcengineSecurityGroupRule())
	if err != nil {
		return fmt.Errorf("error on updating securityGroupRuleService  %q, %w", d.Id(), err)
	}
	return resourceVolcengineSecurityGroupRuleRead(d, meta)
}

func resourceVolcengineSecurityGroupRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupRuleService := NewSecurityGroupRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(securityGroupRuleService, d, ResourceVolcengineSecurityGroupRule())
	if err != nil {
		return fmt.Errorf("error on deleting securityGroupRuleService %q, %w", d.Id(), err)
	}
	return err
}
