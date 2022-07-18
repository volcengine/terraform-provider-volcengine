package rds_instance

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_instance.default mysql-42b38c769c4b
```

*/

func ResourceVolcengineRdsInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsInstanceCreate,
		Read:   resourceVolcengineRdsInstanceRead,
		Delete: resourceVolcengineRdsInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the RDS instance.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region of the RDS instance.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone of the RDS instance.",
			},
			"db_engine": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The engine of the RDS instance.",
			},
			"db_engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The engine version of the RDS instance.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the RDS instance.",
			},
			"instance_spec_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The spec name of the RDS instance.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The storage type of the RDS instance.",
			},
			"storage_space_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The storage space(GB) of the RDS instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The vpc ID of the RDS instance.",
			},
			"number": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The purchase number of the RDS instance.",
			},
			"instance_category": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The category of the RDS instance.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The charge type of the RDS instance.",
			},
			"auto_renew": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to automatically renew.",
			},
			"prepaid_period": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Purchase cycle in prepaid scenarios.",
			},
			"used_time": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The purchase time of RDS instance.",
			},
			"connection_info": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The connection info ot the RDS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internal_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The internal domain of the RDS instance.",
						},
						"internal_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The interval port of the RDS instance.",
						},
						"public_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public domain of the RDS instance.",
						},
						"public_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public port of the RDS instance.",
						},
						"enable_read_write_splitting": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether read-write separation is enabled.",
						},
						"enable_read_only": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether global read-only is enabled.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineRdsInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Create(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on creating RDS instance %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsInstanceRead(d, meta)
}

func resourceVolcengineRdsInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Read(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on reading RDS instance %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Delete(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on deleting RDS instance %q, %w", d.Id(), err)
	}
	return err
}
