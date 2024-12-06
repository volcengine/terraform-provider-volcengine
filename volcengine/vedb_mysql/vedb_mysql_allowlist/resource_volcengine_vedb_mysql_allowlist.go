package vedb_mysql_allowlist

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlAllowlist can be imported using the id, e.g.
```
$ terraform import volcengine_vedb_mysql_allowlist.default resource_id
```

*/

func ResourceVolcengineVedbMysqlAllowlist() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlAllowlistCreate,
		Read:   resourceVolcengineVedbMysqlAllowlistRead,
		Update: resourceVolcengineVedbMysqlAllowlistUpdate,
		Delete: resourceVolcengineVedbMysqlAllowlistDelete,
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
				Description: "The name of the allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the allow list.",
			},
			"allow_list_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The type of IP address in the whitelist. Currently only IPv4 addresses are supported.",
			},
			"allow_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Enter an IP address or a range of IP addresses in CIDR format.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the allow list.",
			},
		},
	}
	return resource
}

func resourceVolcengineVedbMysqlAllowlistCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_allowlist %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlAllowlistRead(d, meta)
}

func resourceVolcengineVedbMysqlAllowlistRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_allowlist %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlAllowlistUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVedbMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on updating vedb_mysql_allowlist %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlAllowlistRead(d, meta)
}

func resourceVolcengineVedbMysqlAllowlistDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_allowlist %q, %s", d.Id(), err)
	}
	return err
}
