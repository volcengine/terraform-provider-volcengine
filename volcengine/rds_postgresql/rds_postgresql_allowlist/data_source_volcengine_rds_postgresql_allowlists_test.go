package rds_postgresql_allowlist_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_allowlist"
)

const testAccVolcengineRdsPostgresqlAllowlistsDSConfig = `
data "volcengine_zones" "foo" {}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project-allowlist-ds"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-allowlist-ds"
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
  instance_name       = "acc-test-allowlist-ds"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
}

resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name     = "acc-test-allowlist-ds"
  allow_list_desc     = "acc test allowlist ds"
  allow_list_category = "Ordinary"
  user_allow_list     = ["10.1.1.1", "10.2.3.0/24"]
}

data "volcengine_rds_postgresql_allowlists" "foo" {
  name_regex = ".*allowlist.*"
}
`

func TestAccVolcengineRdsPostgresqlAllowlistsDataSource_Basic(t *testing.T) {
	dsName := "data.volcengine_rds_postgresql_allowlists.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: dsName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_allowlist.NewRdsPostgresqlAllowlistService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { volcengine.AccTestPreCheck(t) },
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsPostgresqlAllowlistsDSConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "total_count", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "postgresql_allow_lists.#", "1"),
				),
			},
		},
	})
}
