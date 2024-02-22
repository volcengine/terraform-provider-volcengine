package cdn_domain

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CdnDomain can be imported using the domain, e.g.
```
$ terraform import volcengine_cdn_domain.default www.volcengine.com
```

*/

func ResourceVolcengineCdnDomain() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCdnDomainCreate,
		Read:   resourceVolcengineCdnDomainRead,
		Update: resourceVolcengineCdnDomainUpdate,
		Delete: resourceVolcengineCdnDomainDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "You need to add a domain. The main account can add up to 200 accelerated domains.",
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The business type of the domain name is indicated by this parameter. " +
					"The possible values are: `download`: for file downloads. `web`: for web pages. " +
					"`video`: for audio and video on demand.",
			},
			"service_region": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "Indicates the acceleration area. The parameter can take the following values: " +
					"`chinese_mainland`: Indicates mainland China. `global`: Indicates global." +
					" `outside_chinese_mainland`: Indicates global (excluding mainland China).",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "default",
				Description: "The project to which this domain name belongs. Default is `default`.",
			},
			"resource_tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Indicate the tags you have set for this domain name. You can set up to 10 tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the tag.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the tag.",
						},
					},
				},
			},
			"shared_cname": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Configuration for sharing CNAME.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Specify whether to enable shared CNAME.",
						},
						"cname": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Assign a CNAME to the accelerated domain.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineCdnDomainCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCdnDomain())
	if err != nil {
		return fmt.Errorf("error on creating cdn_domain %q, %s", d.Id(), err)
	}
	return resourceVolcengineCdnDomainRead(d, meta)
}

func resourceVolcengineCdnDomainRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCdnDomain())
	if err != nil {
		return fmt.Errorf("error on reading cdn_domain %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCdnDomainUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCdnDomain())
	if err != nil {
		return fmt.Errorf("error on updating cdn_domain %q, %s", d.Id(), err)
	}
	return resourceVolcengineCdnDomainRead(d, meta)
}

func resourceVolcengineCdnDomainDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCdnDomain())
	if err != nil {
		return fmt.Errorf("error on deleting cdn_domain %q, %s", d.Id(), err)
	}
	return err
}
