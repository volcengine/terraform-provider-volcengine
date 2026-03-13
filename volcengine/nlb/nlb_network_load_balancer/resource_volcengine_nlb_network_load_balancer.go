package network_load_balancer

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NlbNetworkLoadBalancer can be imported using the NLB instance ID, e.g.
```
$ terraform import volcengine_nlb_network_load_balancer.foo nlb-1xxx
```

*/

func ResourceVolcengineNlbNetworkLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNlbNetworkLoadBalancerCreate,
		Read:   resourceVolcengineNlbNetworkLoadBalancerRead,
		Update: resourceVolcengineNlbNetworkLoadBalancerUpdate,
		Delete: resourceVolcengineNlbNetworkLoadBalancerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the NLB instance.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the NLB instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the VPC where the NLB instance is deployed.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region of the NLB instance.",
			},
			"network_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"internet", "intranet"}, false),
				Description:  "The network type of the NLB instance. Valid values: `internet`, `intranet`.\n`internet`: The NLB instance is an internet-facing instance.\n`intranet`: The NLB instance is an internal-facing instance.",
			},
			"ip_address_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "ipv4",
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "dualstack"}, false),
				Description:  "The IP address version of the NLB instance. Valid values: `ipv4`, `dualstack`.\n`ipv4`: The NLB instance supports IPv4.\n`dualstack`: The NLB instance supports both IPv4 and IPv6.",
			},
			"cross_zone_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable cross-zone load balancing for the NLB instance. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the NLB instance.",
			},
			"modification_protection_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"NonProtection", "ConsoleProtection"}, false),
				Description:  "The modification protection status of the NLB instance. Valid values: `NonProtection`, `ConsoleProtection`.\n`NonProtection`: No protection.\n`ConsoleProtection`: Console protection.",
			},
			"modification_protection_reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The reason for modification protection.",
			},
			"security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The security group IDs of the NLB instance.",
			},
			"ipv4_bandwidth_package_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the IPv4 bandwidth package.",
			},
			"ipv6_bandwidth_package_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the IPv6 bandwidth package.",
			},
			"tags": ve.TagsSchema(),
			"zone_mappings": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The zone mappings of the NLB instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of the zone.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of the subnet.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IPv4 address of the NLB instance.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IPv6 address of the NLB instance.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The ID of the EIP.",
						},
						"eni_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the ENI.",
						},
						"ipv4_eip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public IPv4 address.",
						},
						"ipv4_eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the EIP.",
						},
						"ipv6_eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the IPv6 EIP.",
						},
						"ipv4_hc_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv4 health check status.",
						},
						"ipv6_hc_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv6 health check status.",
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the NLB instance.",
			},
			"dns_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DNS name of the NLB instance.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the NLB instance.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the NLB instance.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account ID of the NLB instance.",
			},
			"billing_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The billing status of the NLB instance.",
			},
			"overdue_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The overdue time of the NLB instance.",
			},
			"reclaimed_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reclaimed time of the NLB instance.",
			},
			"expected_overdue_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expected overdue time of the NLB instance.",
			},
			"managed_security_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The managed security group ID of the NLB instance.",
			},
			"ipv4_network_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IPv4 network type of the NLB instance.",
			},
			"ipv6_network_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IPv6 network type of the NLB instance.",
			},
			"access_log_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The access log configuration of the NLB instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable access logging. Valid values: `true`, `false`.\n`true`: Enable.\n`false`: Disable.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The project name where the access log topic resides.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The topic name of the access log.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The topic ID of the access log.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineNlbNetworkLoadBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbNetworkLoadBalancerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNlbNetworkLoadBalancer())
	if err != nil {
		return fmt.Errorf("error on creating nlb %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbNetworkLoadBalancerRead(d, meta)
}

func resourceVolcengineNlbNetworkLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbNetworkLoadBalancerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNlbNetworkLoadBalancer())
	if err != nil {
		return fmt.Errorf("error on reading nlb %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNlbNetworkLoadBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbNetworkLoadBalancerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNlbNetworkLoadBalancer())
	if err != nil {
		return fmt.Errorf("error on updating nlb %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbNetworkLoadBalancerRead(d, meta)
}

func resourceVolcengineNlbNetworkLoadBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbNetworkLoadBalancerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNlbNetworkLoadBalancer())
	if err != nil {
		return fmt.Errorf("error on deleting nlb %q, %w", d.Id(), err)
	}
	return err
}
