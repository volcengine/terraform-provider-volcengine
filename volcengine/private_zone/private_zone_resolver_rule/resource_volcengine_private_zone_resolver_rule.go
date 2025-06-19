package private_zone_resolver_rule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZoneResolverRule can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_resolver_rule.default resource_id
```

*/

func ResourceVolcenginePrivateZoneResolverRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneResolverRuleCreate,
		Read:   resourceVolcenginePrivateZoneResolverRuleRead,
		Update: resourceVolcenginePrivateZoneResolverRuleUpdate,
		Delete: resourceVolcenginePrivateZoneResolverRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the rule.",
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Forwarding rule types. " +
					"OUTBOUND: Forward to external DNS servers. " +
					"LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS," +
					" and customize the operator's exit IP address for the recursive DNS server.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the private zone resolver rule.",
			},
			"tags": ve.TagsSchema(),
			"vpc_trns": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "The vpc trns of the private zone resolver rule. Format：trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"endpoint_trn": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "The endpoint trn of the private zone resolver rule. Format：trn:private_zone::accountId:endpoint/endpointId. This field is only effected when creating resource. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"zone_name": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("type").(string) != "OUTBOUND"
				},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Domain names associated with forwarding rules. " +
					"You can enter one or more domain names. Up to 500 domain names are supported. " +
					"This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.",
			},
			"endpoint_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("type").(string) != "OUTBOUND"
				},
				Description: "Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.",
			},
			"forward_ips": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Description: "IP address and port of external DNS server. You can add up to 10 IP addresses. " +
					"This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Required: true,
							Description: "IP address of the external DNS server. " +
								"This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.",
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							Description: "The port of the external DNS server. Default is 53. " +
								"This parameter is only valid and optional when the Type parameter is OUTBOUND.",
						},
					},
				},
			},
			"line": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("type").(string) != "LINE"
				},
				Description: "The operator of the exit IP address of the recursive DNS server. " +
					"This parameter is only valid when the Type parameter is LINE and is a required parameter. " +
					"MOBILE, TELECOM, UNICOM.",
			},
			"vpcs": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Description: "The parameter name <region> is a variable that represents the region where the VPC is located, " +
					"such as cn-beijing. The parameter value can include one or more VPC IDs, " +
					"such as vpc-2750bd1. " +
					"For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, " +
					"the parameter would be \"cn-beijing\":[\"vpc-2d6si87atfh1c58ozfd0nzq8k\"]. " +
					"You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The region of the bind vpc. The default value is the region of the default provider config.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the bind vpc.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneResolverRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on creating private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneResolverRuleRead(d, meta)
}

func resourceVolcenginePrivateZoneResolverRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on reading private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneResolverRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on updating private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneResolverRuleRead(d, meta)
}

func resourceVolcenginePrivateZoneResolverRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZoneResolverRule())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone_resolver_rule %q, %s", d.Id(), err)
	}
	return err
}
