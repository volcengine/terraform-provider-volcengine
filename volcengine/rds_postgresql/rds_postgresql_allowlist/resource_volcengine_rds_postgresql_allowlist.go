package rds_postgresql_allowlist

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlAllowlist can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_allowlist.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlAllowlist() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlAllowlistCreate,
		Read:   resourceVolcengineRdsPostgresqlAllowlistRead,
		Update: resourceVolcengineRdsPostgresqlAllowlistUpdate,
		Delete: resourceVolcengineRdsPostgresqlAllowlistDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the postgresql allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the postgresql allow list.",
			},
			"allow_list_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The type of IP address in the whitelist. Currently only `IPv4` addresses are supported.",
			},
			"allow_list": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 300,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Enter an IP address or a range of IP addresses in CIDR format.",
			},

			// computed fields
			"associated_instance_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of instances bound under the whitelist.",
			},
			"associated_instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of postgresql instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the postgresql instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the postgresql instance.",
						},
						"vpc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlAllowlistCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlAllowlistUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return err
}
