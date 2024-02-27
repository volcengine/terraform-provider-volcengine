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
Please note that when you execute destroy, we will first take the domain name offline and then delete it.
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
			Update: schema.DefaultTimeout(30 * time.Minute),
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
				Computed: true,
				Description: "Indicates the acceleration area. The parameter can take the following values: " +
					"`chinese_mainland`: Indicates mainland China. `global`: Indicates global." +
					" `outside_chinese_mainland`: Indicates global (excluding mainland China).",
			},
			"origin_protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Configuration for origin-pull protocol. " +
					"The parameter can take the following values: " +
					"`http`: HTTP origin-pull will be used for both HTTP and HTTPS requests initiated by the user. " +
					"`https`: HTTPS origin-pull will be used for both HTTP and HTTPS requests initiated by the user. " +
					"`followclient`: HTTP origin-pull will be used for HTTP requests initiated by the user, " +
					"and HTTPS origin-pull will be used for HTTPS requests initiated by the user.",
			},
			"origin_host": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The value of the Host header for origin retrieval, " +
					"with a maximum length of 1,024 characters. " +
					"This parameter has the following specifications: " +
					"- The default value of this parameter is the same as the accelerated domain name. " +
					"- The priority of this parameter is lower than the OriginHost parameter in the origin configuration module. " +
					"- This parameter does not take effect if the InstanceType of the origin is tos.",
			},
			"domain_config": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Accelerate domain configuration. " +
					"Please convert the configuration module structure into json and pass it into a string. " +
					"You must specify the Origin module. The OriginProtocol parameter, OriginHost parameter, " +
					"and other domain configuration modules are optional.",
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
				ForceNew:    true,
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
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the domain.",
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
