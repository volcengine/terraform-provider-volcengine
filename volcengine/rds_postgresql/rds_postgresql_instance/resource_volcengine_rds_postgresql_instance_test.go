package rds_postgresql_instance_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
)

const testAccVolcengineRdsPostgresqlInstanceCreateConfig = `
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
`

const testAccVolcengineRdsPostgresqlInstanceUpdateConfig = `
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
    storage_space          = 20
    subnet_id = volcengine_subnet.foo.id
    instance_name          = "acc-test-2"
    charge_info {
        charge_type = "PostPaid"
    }
    project_name = "default"
    tags {
        key   = "tfk2"
        value = "tfv2"
    }
    parameters {
        name  = "auto_explain.log_analyze"
        value = "on"
    }
    parameters {
        name  = "auto_explain.log_format"
        value = "xml"
    }
}
`

func TestAccVolcengineRdsPostgresqlInstanceResource_Basic(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_instance.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_instance.NewRdsPostgresqlInstanceService(client)
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
				Config: testAccVolcengineRdsPostgresqlInstanceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_info.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_info.0.charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "db_engine_version", "PostgreSQL_12"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_spec", "rds.postgres.1c2g"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameters.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "storage_space", "40"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tfk1",
						"value": "tfv1",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "parameters.*", map[string]string{
						"name":  "auto_explain.log_analyze",
						"value": "off",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "parameters.*", map[string]string{
						"name":  "auto_explain.log_format",
						"value": "text",
					}),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameters"},
			},
		},
	})
}

func TestAccVolcengineRdsPostgresqlInstanceResource_Update(t *testing.T) {
	resourceName := "volcengine_rds_postgresql_instance.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_postgresql_instance.NewRdsPostgresqlInstanceService(client)
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
				Config: testAccVolcengineRdsPostgresqlInstanceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_info.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_info.0.charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "db_engine_version", "PostgreSQL_12"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_spec", "rds.postgres.1c2g"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameters.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "storage_space", "40"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tfk1",
						"value": "tfv1",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "parameters.*", map[string]string{
						"name":  "auto_explain.log_analyze",
						"value": "off",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "parameters.*", map[string]string{
						"name":  "auto_explain.log_format",
						"value": "text",
					}),
				),
			},
			{
				Config: testAccVolcengineRdsPostgresqlInstanceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_info.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_info.0.charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "db_engine_version", "PostgreSQL_12"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_name", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_spec", "rds.postgres.1c2g"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameters.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "storage_space", "20"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "tfk2",
						"value": "tfv2",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "parameters.*", map[string]string{
						"name":  "auto_explain.log_analyze",
						"value": "on",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "parameters.*", map[string]string{
						"name":  "auto_explain.log_format",
						"value": "xml",
					}),
				),
			},
			{
				Config:             testAccVolcengineRdsPostgresqlInstanceUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
