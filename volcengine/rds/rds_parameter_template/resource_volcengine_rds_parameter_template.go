package rds_parameter_template

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_parameter_template.default mysql-sys-80bb93aa14be22d0
```

*/

func ResourceVolcengineRdsParameterTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsParameterTemplateCreate,
		Read:   resourceVolcengineRdsParameterTemplateRead,
		Update: resourceVolcengineRdsParameterTemplateUpdate,
		Delete: resourceVolcengineRdsParameterTemplateDelete,
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
				Optional:    true,
				ForceNew:    true,
				Description: "Parameter template database type, range of values:\nMySQL - MySQL database. (Defaults).",
			},
			"template_type_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Parameter template database version, value range:\nMySQL_Community_5_7 - MySQL 5.7 (default)\nMySQL_8_0 - MySQL 8.0.",
			},
			"template_params": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Template parameters. InstanceParam only needs to pass Name and RunningValue.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Parameter name.",
						},
						"running_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Parameter running value.",
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
}

func resourceVolcengineRdsParameterTemplateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsParameterTemplateService := NewRdsParameterTemplateService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Create(rdsParameterTemplateService, d, ResourceVolcengineRdsParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on creating RDS parameter template %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsParameterTemplateRead(d, meta)
}

func resourceVolcengineRdsParameterTemplateRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsParameterTemplateService := NewRdsParameterTemplateService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Read(rdsParameterTemplateService, d, ResourceVolcengineRdsParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on reading RDS parameter template %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsParameterTemplateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsParameterTemplateService := NewRdsParameterTemplateService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Update(rdsParameterTemplateService, d, ResourceVolcengineRdsParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on updating RDS parameter template %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsParameterTemplateRead(d, meta)
}

func resourceVolcengineRdsParameterTemplateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsParameterTemplateService := NewRdsParameterTemplateService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Delete(rdsParameterTemplateService, d, ResourceVolcengineRdsParameterTemplate())
	if err != nil {
		return fmt.Errorf("error on deleting RDS parameter template %q, %w", d.Id(), err)
	}
	return err
}