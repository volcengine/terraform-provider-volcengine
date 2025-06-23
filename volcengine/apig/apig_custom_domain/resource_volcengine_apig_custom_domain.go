package apig_custom_domain

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigCustomDomain can be imported using the id, e.g.
```
$ terraform import volcengine_apig_custom_domain.default resource_id
```

*/

func ResourceVolcengineApigCustomDomain() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigCustomDomainCreate,
		Read:   resourceVolcengineApigCustomDomainRead,
		Update: resourceVolcengineApigCustomDomainUpdate,
		Delete: resourceVolcengineApigCustomDomainDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the api gateway service.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The custom domain of the api gateway service.",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the certificate.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The resource type of domain. Valid values: `Console`, `Ingress`.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The comments of the custom domain.",
			},
			"ssl_redirect": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to redirect https.",
			},
			"protocol": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The protocol of the custom domain.",
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the custom domain.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the custom domain.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the custom domain.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the domain.",
			},
		},
	}
	return resource
}

func resourceVolcengineApigCustomDomainCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigCustomDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigCustomDomain())
	if err != nil {
		return fmt.Errorf("error on creating apig_custom_domain %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigCustomDomainRead(d, meta)
}

func resourceVolcengineApigCustomDomainRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigCustomDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigCustomDomain())
	if err != nil {
		return fmt.Errorf("error on reading apig_custom_domain %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigCustomDomainUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigCustomDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigCustomDomain())
	if err != nil {
		return fmt.Errorf("error on updating apig_custom_domain %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigCustomDomainRead(d, meta)
}

func resourceVolcengineApigCustomDomainDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigCustomDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigCustomDomain())
	if err != nil {
		return fmt.Errorf("error on deleting apig_custom_domain %q, %s", d.Id(), err)
	}
	return err
}
