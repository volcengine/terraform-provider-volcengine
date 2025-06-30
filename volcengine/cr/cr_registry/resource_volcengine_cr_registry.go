package cr_registry

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR Registry can be imported using the name, e.g.
```
$ terraform import volcengine_cr_registry.default enterprise-x
```

*/

func ResourceVolcengineCrRegistry() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCrRegistryCreate,
		Read:   resourceVolcengineCrRegistryRead,
		Update: resourceVolcengineCrRegistryUpdate,
		Delete: resourceVolcengineCrRegistryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of registry.",
			},
			"delete_immediately": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether delete registry immediately. Only effected in delete action.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of registry user.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the cr registry.",
			},
			"resource_tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The type of registry. Valid values: `Enterprise`, `Micro`. Default is `Enterprise`.",
			},
			"proxy_cache_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Whether to enable proxy cache.",
			},
			"proxy_cache": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "The proxy cache of registry. This field is valid when proxy_cache_enabled is true.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The type of proxy cache. Valid values: `DockerHub`, `DockerRegistry`.",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The endpoint of proxy cache.",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Sensitive:   true,
							Description: "The password of proxy cache.",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The username of proxy cache.",
						},
						"skip_ssl_verify": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Whether to skip ssl verify.",
						},
					},
				},
			},
		},
	}
	dataSource := DataSourceVolcengineCrRegistries().Schema["registries"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrRegistryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("error on creating CrRegistry %q,%s", d.Id(), err)
	}
	return resourceVolcengineCrRegistryRead(d, meta)
}

func resourceVolcengineCrRegistryUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("error on updating CrRegistry  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrRegistryRead(d, meta)
}

func resourceVolcengineCrRegistryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("error on deleting CrRegistry %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrRegistryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCrRegistry())
	if err != nil {
		return fmt.Errorf("Error on reading CrRegistry %q,%s", d.Id(), err)
	}
	return err
}
