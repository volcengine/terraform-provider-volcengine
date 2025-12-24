package rds_postgresql_allowlist_version_upgrade

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func ResourceVolcengineRdsPostgresqlAllowlistVersionUpgrade() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeCreate,
		Read:   resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeRead,
		Delete: resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Description: "The id of the postgresql instance to upgrade allowlist version.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistVersionUpgradeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlAllowlistVersionUpgrade())
	if err != nil {
		return fmt.Errorf("error on upgrading rds_postgresql allowlist version for instance %q, %s", d.Get("instance_id").(string), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeRead(d *schema.ResourceData, meta interface{}) (err error) {
	// Upgrade is one-shot; no remote state to read. Keep ID as instance_id.
	return nil
}

func resourceVolcengineRdsPostgresqlAllowlistVersionUpgradeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	// No remote delete; removing from state is enough.
	return nil
}
