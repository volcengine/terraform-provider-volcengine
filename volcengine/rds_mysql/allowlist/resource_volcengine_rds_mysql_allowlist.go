package allowlist

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS AllowList can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_allowlist.default acl-d1fd76693bd54e658912e7337d5b****
```

*/

func ResourceVolcengineRdsMysqlAllowlist() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsMysqlAllowlistCreate,
		Read:   resourceVolcengineRdsMysqlAllowlistRead,
		Update: resourceVolcengineRdsMysqlAllowlistUpdate,
		Delete: resourceVolcengineRdsMysqlAllowlistDelete,
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
}

func resourceVolcengineRdsMysqlAllowlistCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error creating RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsMysqlAllowlistRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error reading RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlAllowlistUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error updating RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsMysqlAllowlistDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error deleting RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return err
}
