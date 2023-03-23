package security_group_rule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
SecurityGroupRule can be imported using the id, e.g.
```
$ terraform import volcengine_security_group_rule.default ID is a string concatenated with colons(SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp)
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
				}, false),
				Description: "Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.",
			},
			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of SecurityGroup.",
			},
			"port_start": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(-1, 65535),
				Description:  "Port start of egress/ingress Rule.",
			},
			"port_end": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(-1, 65535),
				Description:  "Port end of egress/ingress Rule.",
			},
			"cidr_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Cidr ip of egress/ingress Rule.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of SecurityGroup.",
			},
			"source_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the source security group whose access permission you want to set.",
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
				Description: "Access strategy.",
			},
			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 100),
				Description:  "Priority of a security group rule.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "description of a egress rule.",
			},
		},
	}
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