package rds_postgresql_parameter_template

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlParameterTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_parameter_template.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlParameterTemplate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlParameterTemplateCreate,
		Read:   resourceVolcengineRdsPostgresqlParameterTemplateRead,
		Update: resourceVolcengineRdsPostgresqlParameterTemplateUpdate,
		Delete: resourceVolcengineRdsPostgresqlParameterTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"template_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Parameter template ID.",
			},
			"template_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Parameter template name.",
			},
			"template_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "PostgreSQL",
				Description: "The type of the parameter template. The current value can only be PostgreSQL.",
			},
			"template_type_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The version of PostgreSQL supported by the parameter template. The current value can be PostgreSQL_11/12/13/14/15/16/17.",
			},
			"template_params": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Parameter configuration of the parameter template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the parameter.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the parameter.",
						},
					},
				},
			},
			"template_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the parameter template. The maximum length is 200 characters.",
			},
			"src_template_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of the source parameter template to clone. If set, the parameter template will be cloned from the source template.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of the instance to export the current parameters as a parameter template. If set, the parameter template will be created based on the current parameters of the instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlParameterTemplateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_parameter_template %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlParameterTemplateRead(d, meta)
}

func resourceVolcengineRdsPostgresqlParameterTemplateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_parameter_template %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlParameterTemplateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_parameter_template %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlParameterTemplateRead(d, meta)
}

func resourceVolcengineRdsPostgresqlParameterTemplateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_parameter_template %q, %s", d.Id(), err)
	}
	return err
}
