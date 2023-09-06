package rds_mysql_database_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_database"
)

const testAccVolcengineRdsMysqlDatabasesDatasourceConfig = `
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
data "volcengine_rds_mysql_databases" "foo"{
    db_name = "acc-test"
    instance_id = volcengine_rds_mysql_instance.foo.id
}
`

func TestAccVolcengineRdsMysqlDatabasesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_rds_mysql_databases.foo"

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
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRdsMysqlDatabasesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "databases.#", "1"),
				),
			},
		},
	})
}
