package dns_zone

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Zone can be imported using the id, e.g.
```
$ terraform import volcengine_zone.default resource_id
```

*/

func ResourceVolcengineZone() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineZoneCreate,
		Read:   resourceVolcengineZoneRead,
		Update: resourceVolcengineZoneUpdate,
		Delete: resourceVolcengineZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"tags": ve.TagsSchema(),
			"zone_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The domain to be created. The domain must be a second-level domain and cannot be a wildcard domain.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The remark for the domain.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project to which the domain name belongs. The default value is default.",
			},
			// computed
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the domain.",
			},
			"dns_security": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of DNS DDoS protection service.",
			},
			"expired_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The expiration time of the domain.",
			},
			"is_sub_domain": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the domain is a subdomain.",
			},
			"record_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of DNS records under the domain.",
			},
			"trade_code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The edition of the domain.",
			},
			"zid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the domain.",
			},
			"allocate_dns_server_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of DNS servers allocated to the domain by BytePlus DNS.",
			},
			"auto_renew": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether automatic domain renewal is enabled.",
			},
			"instance_no": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the instance. For free edition, the value of this field is null.",
			},
			"is_ns_correct": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the configuration of NS servers is correct. If the configuration is correct, the status of the domain in BytePlus DNS is Active.",
			},
			"real_dns_server_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of DNS servers actually used by the domain.",
			},
			"stage": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status of the domain.",
			},
			"sub_domain_host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The domain prefix of the subdomain. If the domain is not a subdomain, this parameter is null.",
			},
		},
	}
	return resource
}

func resourceVolcengineZoneCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineZone())
	if err != nil {
		return fmt.Errorf("error on creating zone %q, %s", d.Id(), err)
	}
	return resourceVolcengineZoneRead(d, meta)
}

func resourceVolcengineZoneRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineZone())
	if err != nil {
		return fmt.Errorf("error on reading zone %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineZoneUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineZone())
	if err != nil {
		return fmt.Errorf("error on updating zone %q, %s", d.Id(), err)
	}
	return resourceVolcengineZoneRead(d, meta)
}

func resourceVolcengineZoneDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineZone())
	if err != nil {
		return fmt.Errorf("error on deleting zone %q, %s", d.Id(), err)
	}
	return err
}
