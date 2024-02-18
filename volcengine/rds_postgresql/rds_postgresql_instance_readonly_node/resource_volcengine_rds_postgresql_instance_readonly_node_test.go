package rds_postgresql_instance_readonly_node_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance_readonly_node"
)

const testAccVolcengineRdsPostgresqlInstanceReadonlyNodeCreateConfig = `
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
    vpc_name   = "acc-test-project1"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-subnet-test-2"
    cidr_block  = "172.16.0.0/24"
    zone_id     = data.volcengine_zones.foo.zones[0].id
    vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_rds_postgresql_instance" "foo" {
    db_engine_version = "PostgreSQL_12"
    node_spec = "rds.postgres.1c2g"
    primary_zone_id        = data.volcengine_zones.foo.zones[0].id
    secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
    storage_space          = 40
    subnet_id = volcengine_subnet.foo.id
    instance_name          = "acc-test-1"
    charge_info {
        charge_type = "PostPaid"
    }
    project_name = "default"
    tags {
        key   = "tfk1"
        value = "tfv1"
    }
    parameters {
        name  = "auto_explain.log_analyze"
        value = "off"
    }
    parameters {
        name  = "auto_explain.log_format"
        value = "text"
    }
}

resource "volcengine_rds_postgresql_instance_readonly_node" "foo" {
    instance_id = volcengine_rds_postgresql_instance.foo.id
    node_spec = "rds.postgres.1c2g"
    zone_id = data.volcengine_zones.foo.zones[0].id
}
`

const testAccVolcengineRdsPostgresqlInstanceReadonlyNodeUpdateConfig = `
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
    vpc_name   = "acc-test-project1"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-subnet-test-2"
    cidr_block  = "172.16.0.0/24"
    zone_id     = data.volcengine_zones.foo.zones[0].id
    vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_rds_postgresql_instance" "foo" {
    db_engine_version = "PostgreSQL_12"
    node_spec = "rds.postgres.1c2g"
    primary_zone_id        = data.volcengine_zones.foo.zones[0].id
    secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
    storage_space          = 40
    subnet_id = volcengine_subnet.foo.id
    instance_name          = "acc-test-1"
    charge_info {
        charge_type = "PostPaid"
    }
    project_name = "default"
    tags {
        key   = "tfk1"
        value = "tfv1"
    }
    parameters {
        name  = "auto_explain.log_analyze"
        value = "off"
    }
    parameters {
        name  = "auto_explain.log_format"
        value = "text"
    }
}

resource "volcengine_rds_postgresql_instance_readonly_node" "foo" {
    instance_id = volcengine_rds_postgresql_instance.foo.id
    node_spec = "rds.postgres.2c4g"
    zone_id = data.volcengine_zones.foo.zones[0].id
}
`

func TestAccVolcengineRdsPostgresqlInstanceReadonlyNodeResource_Basic(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_instance_readonly_node.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_instance_readonly_node.NewRdsPostgresqlInstanceReadonlyNodeService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsPostgresqlInstanceReadonlyNodeCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_spec", "rds.postgres.1c2g"),
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

func TestAccVolcengineRdsPostgresqlInstanceReadonlyNodeResource_Update(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_instance_readonly_node.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_instance_readonly_node.NewRdsPostgresqlInstanceReadonlyNodeService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsPostgresqlInstanceReadonlyNodeCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_spec", "rds.postgres.1c2g"),
				),
			},
			{
				Config: testAccVolcengineRdsPostgresqlInstanceReadonlyNodeUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_spec", "rds.postgres.2c4g"),
				),
			},
			{
				Config:             testAccVolcengineRdsPostgresqlInstanceReadonlyNodeUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
