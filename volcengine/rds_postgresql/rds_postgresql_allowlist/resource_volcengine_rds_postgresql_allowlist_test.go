package rds_postgresql_allowlist_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_allowlist"
)

const testAccVolcengineRdsPostgresqlAllowlistCreateConfig = `
data "volcengine_zones" "foo" {}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project-allowlist"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-allowlist"
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
  instance_name       = "acc-test-allowlist"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
}

resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name     = "acc-test-allowlist"
  allow_list_desc     = "acc test allowlist"
  allow_list_category = "Ordinary"
  user_allow_list     = ["10.1.1.1", "10.2.3.0/24"]
}

resource "volcengine_rds_postgresql_allowlist" "unify" {
  allow_list_name  = "acc-test-unify-allowlist"
  allow_list_desc  = "unify from instances"
  instance_ids     = [volcengine_rds_postgresql_instance.foo.id]
}
`

const testAccVolcengineRdsPostgresqlAllowlistUpdateConfig = `
data "volcengine_zones" "foo" {}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project-allowlist"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-allowlist"
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
  instance_name       = "acc-test-allowlist"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
}

resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name     = "acc-test-allowlist"
  allow_list_desc     = "acc test allowlist updated"
  allow_list_category = "Ordinary"
  modify_mode         = "Cover"
  update_security_group = false
  user_allow_list     = ["10.1.1.1", "10.2.3.0/24"]
}
`

func TestAccVolcengineRdsPostgresqlAllowlistResource_Basic(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_allowlist.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_allowlist.NewRdsPostgresqlAllowlistService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { volcengine.AccTestPreCheck(t) },
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsPostgresqlAllowlistCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_category", "Ordinary"),
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

func TestAccVolcengineRdsPostgresqlAllowlistResource_Update(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_allowlist.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_allowlist.NewRdsPostgresqlAllowlistService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { volcengine.AccTestPreCheck(t) },
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsPostgresqlAllowlistCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc test allowlist"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_category", "Ordinary"),
				),
			},
			{
				Config: testAccVolcengineRdsPostgresqlAllowlistUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc test allowlist updated"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modify_mode", "Cover"),
				),
			},
			{
				Config:             testAccVolcengineRdsPostgresqlAllowlistUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
