package rds_ip_list

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDSIPList can be imported using the id, e.g.
```
$ terraform import volcengine_rds_ip_list.default mysql-42b38c769c4b:group_name
```

*/

func ResourceVolcengineRdsIpList() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsIpListCreate,
		Read:   resourceVolcengineRdsIpListRead,
		Update: resourceVolcengineRdsIpListUpdate,
		Delete: resourceVolcengineRdsIpListDelete,
		Importer: &schema.ResourceImporter{
			State: rdsIPListImporter,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the RDS instance.",
			},
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the RDS ip list.",
			},
			"ip_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of IP address.",
			},
		},
	}
}

func resourceVolcengineRdsIpListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsIpListService := NewRdsIpListService(meta.(*volc.SdkClient))
	err = rdsIpListService.Dispatcher.Create(rdsIpListService, d, ResourceVolcengineRdsIpList())
	if err != nil {
		return fmt.Errorf("error on creating RDS ip list %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsIpListRead(d, meta)
}

func resourceVolcengineRdsIpListRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsIpListService := NewRdsIpListService(meta.(*volc.SdkClient))
	err = rdsIpListService.Dispatcher.Read(rdsIpListService, d, ResourceVolcengineRdsIpList())
	if err != nil {
		return fmt.Errorf("error on reading RDS ip list %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsIpListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsIpListService := NewRdsIpListService(meta.(*volc.SdkClient))
	err = rdsIpListService.Dispatcher.Update(rdsIpListService, d, ResourceVolcengineRdsIpList())
	if err != nil {
		return fmt.Errorf("error on reading RDS ip list %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsIpListRead(d, meta)
}

func resourceVolcengineRdsIpListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsIpListService := NewRdsIpListService(meta.(*volc.SdkClient))
	err = rdsIpListService.Dispatcher.Delete(rdsIpListService, d, ResourceVolcengineRdsIpList())
	if err != nil {
		return fmt.Errorf("error on deleting RDS ip list %q, %w", d.Id(), err)
	}
	return err
}
