package rds_mysql_parameter_template

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsMysqlParameterTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_parameter_template.default resource_id
```

*/

func ResourceVolcengineRdsMysqlParameterTemplate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMysqlParameterTemplateCreate,
		Read:   resourceVolcengineRdsMysqlParameterTemplateRead,
		Update: resourceVolcengineRdsMysqlParameterTemplateUpdate,
		Delete: resourceVolcengineRdsMysqlParameterTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Parameter template name.",
			},
			"template_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Database type of parameter template. The default value is Mysql.",
			},
			"template_type_version": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Database version of parameter template. " +
					"Value range:\nMySQL_5_7: Default value. MySQL 5.7 version.\nMySQL_8_0: MySQL 8.0 version.",
			},
			"template_params": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Parameters contained in the parameter template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Instance parameter name.\n" +
								"Description: When using CreateParameterTemplate and ModifyParameterTemplate as request parameters," +
								" only Name and RunningValue need to be passed in.",
						},
						"running_value": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Parameter running value.\n" +
								"Description: When making request parameters in CreateParameterTemplate and ModifyParameterTemplate," +
								" only Name and RunningValue need to be passed in.",
						},
					},
				},
			},
			"template_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parameter template description.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsMysqlParameterTemplateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on creating rds_mysql_parameter_template %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlParameterTemplateRead(d, meta)
}

func resourceVolcengineRdsMysqlParameterTemplateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on reading rds_mysql_parameter_template %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlParameterTemplateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMysqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on updating rds_mysql_parameter_template %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlParameterTemplateRead(d, meta)
}

func resourceVolcengineRdsMysqlParameterTemplateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlParameterTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on deleting rds_mysql_parameter_template %q, %s", d.Id(), err)
	}
	return err
}
