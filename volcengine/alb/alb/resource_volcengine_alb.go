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
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
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
			"subnet_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The id of the Subnet.",
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
