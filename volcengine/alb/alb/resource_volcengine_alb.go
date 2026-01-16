package alb

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Alb can be imported using the id, e.g.
```
$ terraform import volcengine_alb.default resource_id
```

*/

func ResourceVolcengineAlb() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbCreate,
		Read:   resourceVolcengineAlbRead,
		Update: resourceVolcengineAlbUpdate,
		Delete: resourceVolcengineAlbDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"address_ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "IPv4",
				ValidateFunc: validation.StringInSlice([]string{"IPv4", "DualStack"}, false),
				Description:  "The address ip version of the Alb. Valid values: `IPv4`, `DualStack`. Default is `ipv4`.",
			},
			// ModifyLoadBalancerType 允许变更实例的网络类型
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"public", "private"}, false),
				Description:  "The type of the Alb. Valid values: `public`, `private`.",
			},
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the Alb.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the Alb.",
			},
			// ModifyLoadBalancerZones-变更 ALB 实例可用区
			"subnet_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The id of the Subnet.",
			},
			"allocation_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The ID of the public IP. This field is only valid when the type field changes from private to public.",
			},
			// CloneLoadBalancer-复制 ALB 实例时使用
			"source_load_balancer_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The source ALB instance ID for cloning. If specified, the ALB instance will be cloned from this source.",
			},
			//"bandwidth_package_id": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	ForceNew:    true,
			//	Description: "The bandwidth package id of the Eip which automatically associated to the Alb. This field is valid when the type of the Alb is `public`.",
			//},
			"delete_protection": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "off",
				ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
				Description:  "Whether to enable the delete protection function of the Alb. Valid values: `on`, `off`. Default is `off`.",
			},
			// 新增字段
			"modification_protection_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"NonProtection", "ConsoleProtection"}, false),
				Description: "Whether to enable the modification protection function of the Alb. Valid values: `NonProtection`, `ConsoleProtection`. Default is `NonProtection`. " +
					"NonProtection: Instance modification protection is not enabled. ConsoleProtection: Instance modification protection is enabled; you cannot modify the instance configuration through the ALB console, and can only modify the instance configuration by calling the API.",
			},
			"modification_protection_reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The reason for enabling instance modification protection. This parameter is valid when the modification_protection_status is `ConsoleProtection`.",
			},
			"load_balancer_edition": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"Basic", "Standard"}, false),
				Description:  "The version of the ALB instance. Basic: Basic Edition. Standard: Standard Edition. Default is `Basic`.",
			},

			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the Alb.",
			},
			"tags": ve.TagsSchema(),

			"eip_billing_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Description: "The billing configuration of the EIP which automatically associated to the Alb. This field is valid when the type of the Alb is `public`." +
					"When the type of the Alb is `private`, suggest using a combination of resource `volcengine_eip_address` and `volcengine_eip_associate` to achieve public network access function.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isp": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"BGP"}, false),
							Description:  "The ISP of the EIP which automatically associated to the Alb, the value can be `BGP`.",
						},
						"eip_billing_type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
							Description:  "The billing type of the EIP which automatically assigned to the Alb. Valid values: `PostPaidByBandwidth`, `PostPaidByTraffic`.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "The peek bandwidth of the EIP which automatically assigned to the Alb. Unit: Mbps.",
						},
					},
				},
			},
			"ipv6_eip_billing_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Description: "The billing configuration of the Ipv6 EIP which automatically associated to the Alb. This field is required when the type of the Alb is `public`." +
					"When the type of the Alb is `private`, suggest using a combination of resource `volcengine_vpc_ipv6_gateway` and `volcengine_vpc_ipv6_address_bandwidth` to achieve ipv6 public network access function.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isp": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"BGP"}, false),
							Description:  "The ISP of the Ipv6 EIP which automatically associated to the Alb, the value can be `BGP`.",
						},
						"billing_type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
							Description:  "The billing type of the Tpv6 EIP which automatically assigned to the Alb. Valid values: `PostPaidByBandwidth`, `PostPaidByTraffic`.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "The peek bandwidth of the Ipv6 EIP which automatically assigned to the Alb. Unit: Mbps.",
						},
					},
				},
			},
			// ModifyLoadBalancerAttributes-修改 ALB 实例属性时使用
			"waf_protection_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"off", "on"}, false),
				Description:  "Whether to enable the WAF protection function of the Alb. Valid values: `off`, `on`. Default is `off`.",
			},
			"waf_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID of the WAF instance to be associated with the Alb. This field is valid when the value of the `waf_protection_enabled` is `on`.",
			},
			"waf_protected_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain name of the WAF protected Alb. This field is valid when the value of the `waf_protection_enabled` is `on`.",
			},
			"global_accelerator": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The global accelerator configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accelerator_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The global accelerator id.",
						},
						"accelerator_listener_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The global accelerator listener id.",
						},
						"endpoint_group_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The global accelerator endpoint group id.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The traffic distribution weight of the endpoint. The value range is: 1 - 100.",
						},
					},
				},
			},
			"proxy_protocol_enabled": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"off", "on"}, false),
				Description:  "ALB can support the Proxy Protocol and record the real IP of the client.",
			},
			// computed fields
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vpc id of the Alb.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the Alb.",
			},
			"dns_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DNS name.",
			},
			"local_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The local addresses of the Alb.",
			},
			"zone_mappings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configuration information of the Alb instance in different Availability Zones.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The availability zone id of the Alb.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the Alb in this availability zone.",
						},
						"load_balancer_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The IP address information of the Alb in this availability zone.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eni_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Eni address of the Alb in this availability zone.",
									},
									"eni_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Eni id of the Alb in this availability zone.",
									},
									"eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Eip address of the Alb in this availability zone.",
									},
									"eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Eip id of alb instance in this availability zone.",
									},
									"eni_ipv6_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Eni Ipv6 address of the Alb in this availability zone.",
									},
									"ipv6_eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Ipv6 Eip id of alb instance in this availability zone.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineAlbCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlb())
	if err != nil {
		return fmt.Errorf("error on creating alb %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbRead(d, meta)
}

func resourceVolcengineAlbRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlb())
	if err != nil {
		return fmt.Errorf("error on reading alb %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlb())
	if err != nil {
		return fmt.Errorf("error on updating alb %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbRead(d, meta)
}

func resourceVolcengineAlbDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlb())
	if err != nil {
		return fmt.Errorf("error on deleting alb %q, %s", d.Id(), err)
	}
	return err
}
