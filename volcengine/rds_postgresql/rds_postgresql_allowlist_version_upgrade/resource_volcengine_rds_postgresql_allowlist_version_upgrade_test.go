package rds_postgresql_allowlist_version_upgrade_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	upgrade "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_allowlist_version_upgrade"
)

const testAccVolcengineRdsPostgresqlAllowlistVersionUpgradeConfig = `
data "volcengine_zones" "foo" {}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project-allowlist-upgrade"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-allowlist-upgrade"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rds_postgresql_instance" "foo" {
  db_engine_version   = "PostgreSQL_12"
  node_spec           = "rds.postgres.1c2g"
  primary_zone_id     = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id   = data.volcengine_zones.foo.zones[0].id
  storage_space       = 40
  subnet_id           = volcengine_subnet.foo.id
  instance_name       = "acc-test-allowlist-upgrade"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
}

resource "volcengine_rds_postgresql_allowlist_version_upgrade" "upgrade" {
  instance_id = volcengine_rds_postgresql_instance.foo.id
}
`

func TestAccVolcengineRdsPostgresqlAllowlistVersionUpgrade_Basic(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_allowlist_version_upgrade.upgrade"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return upgrade.NewRdsPostgresqlAllowlistVersionUpgradeService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { volcengine.AccTestPreCheck(t) },
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsPostgresqlAllowlistVersionUpgradeConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
