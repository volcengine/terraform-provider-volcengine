package rds_postgresql_instance_ssl

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlInstanceSsl can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance_ssl.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlInstanceSsl() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlInstanceSslCreate,
		Read:   resourceVolcengineRdsPostgresqlInstanceSslRead,
		Update: resourceVolcengineRdsPostgresqlInstanceSslUpdate,
		Delete: resourceVolcengineRdsPostgresqlInstanceSslDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the postgresql Instance.",
			},
			"ssl_enable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable SSL.",
			},
			"force_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Whether to enable force encryption. " +
					"This only takes effect when the SSL encryption function of the instance is enabled.",
			},
			"reload_ssl_certificate": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Update the validity period of the SSL certificate. " +
					"This only takes effect when the SSL encryption function of the instance is enabled. " +
					"It is not supported to pass in reload_ssl_certificate and ssl_enable at the same time.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlInstanceSslCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceSslService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlInstanceSsl())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_instance_ssl %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlInstanceSslRead(d, meta)
}

func resourceVolcengineRdsPostgresqlInstanceSslRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceSslService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlInstanceSsl())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_instance_ssl %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlInstanceSslUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceSslService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlInstanceSsl())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_instance_ssl %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlInstanceSslRead(d, meta)
}

func resourceVolcengineRdsPostgresqlInstanceSslDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceSslService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlInstanceSsl())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_instance_ssl %q, %s", d.Id(), err)
	}
	return err
}
