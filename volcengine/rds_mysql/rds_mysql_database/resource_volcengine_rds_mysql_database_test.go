package rds_mysql_database_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_database"
)

const testAccVolcengineRdsMysqlDatabaseCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
    vpc_name = "acc-test-project1"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-subnet-test-2"
    cidr_block = "172.16.0.0/24"
    zone_id = data.volcengine_zones.foo.zones[0].id
    vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version = "MySQL_5_7"
  node_spec = "rds.mysql.1c2g"
  primary_zone_id = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
  storage_space = 80
  subnet_id = volcengine_subnet.foo.id
  instance_name = "acc-test"
  lower_case_table_names = "1"

  charge_info {
    charge_type = "PostPaid"
  }

  parameters {
    parameter_name = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name = "auto_increment_offset"
    parameter_value = "4"
  }
}

resource "volcengine_rds_mysql_database" "foo" {
    db_name = "acc-test"
    instance_id = volcengine_rds_mysql_instance.foo.id
}
`

func TestAccVolcengineRdsMysqlDatabaseResource_Basic(t *testing.T) {
	resourceName := "volcengine_rds_mysql_database.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return rds_mysql_database.NewRdsMysqlDatabaseService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		// CheckDestroy 此处可能存在问题，业务方接口查询已删除实例的数据库会报内部错误，预计于2023-09修复
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsMysqlDatabaseCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "character_set_name", "utf8mb4"),
					resource.TestCheckResourceAttr(acc.ResourceId, "db_name", "acc-test"),
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
