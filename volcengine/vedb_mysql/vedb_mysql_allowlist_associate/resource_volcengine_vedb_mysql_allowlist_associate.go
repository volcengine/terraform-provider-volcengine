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
VedbMysqlAllowlistAssociate can be imported using the instance id and allow list id, e.g.
```
$ terraform import volcengine_vedb_mysql_allowlist_associate.default vedbm-iqnh3a7z****:acl-d1fd76693bd54e658912e7337d5b****
```

*/

func ResourceVolcengineVedbMysqlAllowlistAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlAllowlistAssociateCreate,
		Read:   resourceVolcengineVedbMysqlAllowlistAssociateRead,
		Delete: resourceVolcengineVedbMysqlAllowlistAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: importAllowListAssociate,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the mysql instance.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the allow list.",
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

func importAllowListAssociate(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:AllowListId")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("allow_list_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
