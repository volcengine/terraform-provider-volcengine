package rds_postgresql_allowlist_associate

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlAllowlistAssociate can be imported using the instance_id:allow_list_id, e.g.
```
$ terraform import volcengine_rds_postgresql_allowlist_associate.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlAllowlistAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlAllowlistAssociateCreate,
		Read:   resourceVolcengineRdsPostgresqlAllowlistAssociateRead,
		Delete: resourceVolcengineRdsPostgresqlAllowlistAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: allowListAssociateImporter,
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
				Description: "The id of the postgresql instance.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the postgresql allow list.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlAllowlistAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_allowlist_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistAssociateRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_allowlist_associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlAllowlistAssociateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_allowlist_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistAssociateRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlAllowlistAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_allowlist_associate %q, %s", d.Id(), err)
	}
	return err
}

func allowListAssociateImporter(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
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
