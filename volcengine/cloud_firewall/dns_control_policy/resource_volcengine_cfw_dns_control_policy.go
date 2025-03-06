package dns_control_policy

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
DnsControlPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_dns_control_policy.default resource_id
```

*/

func ResourceVolcengineDnsControlPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineDnsControlPolicyCreate,
		Read:   resourceVolcengineDnsControlPolicyRead,
		Update: resourceVolcengineDnsControlPolicyUpdate,
		Delete: resourceVolcengineDnsControlPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"destination_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination type of the dns control policy. Valid values: `group`, `domain`.",
			},
			"destination": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination of the dns control policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the dns control policy.",
			},
			"internet_firewall_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The internet firewall id of the control policy.",
			},
			"source": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "The source vpc list of the dns control policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the source vpc.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The region of the source vpc.",
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "Whether to enable the dns control policy.",
			},

			// computed fields
			"hit_cnt": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The hit count of the dns control policy.",
			},
			"use_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The use count of the dns control policy.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the dns control policy.",
			},
			"last_hit_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The last hit time of the dns control policy. Unix timestamp.",
			},
		},
	}
	return resource
}

func resourceVolcengineDnsControlPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineDnsControlPolicy())
	if err != nil {
		return fmt.Errorf("error on creating dns_control_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsControlPolicyRead(d, meta)
}

func resourceVolcengineDnsControlPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineDnsControlPolicy())
	if err != nil {
		return fmt.Errorf("error on reading dns_control_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineDnsControlPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineDnsControlPolicy())
	if err != nil {
		return fmt.Errorf("error on updating dns_control_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsControlPolicyRead(d, meta)
}

func resourceVolcengineDnsControlPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineDnsControlPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting dns_control_policy %q, %s", d.Id(), err)
	}
	return err
}
