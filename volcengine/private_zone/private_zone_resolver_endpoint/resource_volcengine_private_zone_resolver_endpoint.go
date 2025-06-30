package private_zone_resolver_endpoint

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZoneResolverEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_resolver_endpoint.default resource_id
```

*/

func ResourceVolcenginePrivateZoneResolverEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneResolverEndpointCreate,
		Read:   resourceVolcenginePrivateZoneResolverEndpointRead,
		Update: resourceVolcenginePrivateZoneResolverEndpointUpdate,
		Delete: resourceVolcenginePrivateZoneResolverEndpointDelete,
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
				Description: "The name of the private zone resolver endpoint.",
			},
			"direction": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "OUTBOUND",
				Description: "DNS request forwarding direction for terminal nodes. " +
					"OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. " +
					"INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC ID of the endpoint.",
			},
			"vpc_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC region of the endpoint.",
			},
			"security_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The security group ID of the endpoint.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the private zone resolver endpoint.",
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
				Description: "The vpc trns of the private zone resolver endpoint. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"ip_configs": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				//Set: func(i interface{}) int {
				//	if i == nil {
				//		return hashcode.String("")
				//	}
				//	m := i.(map[string]interface{})
				//	var (
				//		buf bytes.Buffer
				//	)
				//	buf.WriteString(fmt.Sprintf("%v#%v#%v", m["az_id"], m["subnet_id"], m["ip"]))
				//	return hashcode.String(buf.String())
				//},
				Description: "Availability zones, subnets, and IP configurations of terminal nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"az_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Id of the availability zone.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Id of the subnet.",
						},
						"ip": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Source IP address of traffic. You can add up to 6 IP addresses at most. " +
								"To ensure high availability, you must add at least two IP addresses.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneResolverEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZoneResolverEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating private_zone_resolver_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneResolverEndpointRead(d, meta)
}

func resourceVolcenginePrivateZoneResolverEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZoneResolverEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading private_zone_resolver_endpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneResolverEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZoneResolverEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating private_zone_resolver_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneResolverEndpointRead(d, meta)
}

func resourceVolcenginePrivateZoneResolverEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneResolverEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZoneResolverEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone_resolver_endpoint %q, %s", d.Id(), err)
	}
	return err
}
