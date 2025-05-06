package dns_record

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
DnsRecord can be imported using the id, e.g.
```
$ terraform import volcengine_dns_record.default ZID:recordId
```

*/

func ResourceVolcengineDnsRecord() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineDnsRecordCreate,
		Read:   resourceVolcengineDnsRecordRead,
		Update: resourceVolcengineDnsRecordUpdate,
		Delete: resourceVolcengineDnsRecordDelete,
		Importer: &schema.ResourceImporter{
			State: dnsRecordImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"zid": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the domain to which you want to add a DNS record.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The host record, which is the domain prefix of the subdomain.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The record type.",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the DNS record.",
			},
			"line": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The value of the DNS record.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The remark for the DNS record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The Time-To-Live (TTL) of the DNS record, in seconds.",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The weight of the DNS record.",
			},
			// computed
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the domain.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the domain.",
			},
			"enable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the DNS record is enabled.",
			},
			"operators": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The account ID that called this API.",
			},
			"pqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account ID that called this API.",
			},
			"record_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the DNS record.",
			},
			"record_set_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the record set where the DNS record is located.",
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The tag information of the DNS record.",
			},
		},
	}
	return resource
}

func resourceVolcengineDnsRecordCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineDnsRecord())
	if err != nil {
		return fmt.Errorf("error on creating dns_record %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsRecordRead(d, meta)
}

func resourceVolcengineDnsRecordRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineDnsRecord())
	if err != nil {
		return fmt.Errorf("error on reading dns_record %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineDnsRecordUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineDnsRecord())
	if err != nil {
		return fmt.Errorf("error on updating dns_record %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsRecordRead(d, meta)
}

func resourceVolcengineDnsRecordDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineDnsRecord())
	if err != nil {
		return fmt.Errorf("error on deleting dns_record %q, %s", d.Id(), err)
	}
	return err
}
