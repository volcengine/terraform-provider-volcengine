package vedb_mysql_allowlist_associate

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlAllowlistAssociate can be imported using the instance id, endpoint id and the eip id, e.g.
```
$ terraform import volcengine_vedb_mysql_allowlist_associate.default vedbm-iqnh3a7z****:vedbm-2pf2xk5v****-Custom-50yv:eip-xxxx
```

*/

func ResourceVolcengineVedbMysqlAllowlistAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlAllowlistAssociateCreate,
		Read:   resourceVolcengineVedbMysqlAllowlistAssociateRead,
		Delete: resourceVolcengineVedbMysqlAllowlistAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: vedbMysqlEndpointAssociateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The instance id.",
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The endpoint id.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "EIP ID that needs to be bound to the instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineVedbMysqlAllowlistAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_allowlist_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlAllowlistAssociateRead(d, meta)
}

func resourceVolcengineVedbMysqlAllowlistAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_allowlist_associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlAllowlistAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_allowlist_associate %q, %s", d.Id(), err)
	}
	return err
}

var vedbMysqlEndpointAssociateImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("endpoint_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("eip_id", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
