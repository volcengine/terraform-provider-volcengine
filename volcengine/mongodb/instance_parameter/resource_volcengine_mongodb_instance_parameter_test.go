package instance_parameter_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance_parameter"
)

const testAccVolcengineMongodbInstanceParameterCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_mongodb_instance" "foo"{
    db_engine_version = "MongoDB_4_0"
    instance_type="ReplicaSet"
    super_account_password="@acc-test-123"
    node_spec="mongo.2c4g"
    mongos_node_spec="mongo.mongos.2c4g"
    instance_name="acc-test-mongo-replica"
    charge_type="PostPaid"
    project_name = "default"
    mongos_node_number = 32
    shard_number=3
    storage_space_gb=20
    subnet_id=volcengine_subnet.foo.id
    zone_id= data.volcengine_zones.foo.zones[0].id
    tags {
        key = "k1"
        value = "v1"
    }
}

resource "volcengine_mongodb_instance_parameter" "foo" {
    instance_id = volcengine_mongodb_instance.foo.id
    parameter_name = "cursorTimeoutMillis"
    parameter_role = "Node"
    parameter_value = "600001"
}
`

const testAccVolcengineMongodbInstanceParameterUpdateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_mongodb_instance" "foo"{
    db_engine_version = "MongoDB_4_0"
    instance_type="ReplicaSet"
    super_account_password="@acc-test-123"
    node_spec="mongo.2c4g"
    mongos_node_spec="mongo.mongos.2c4g"
    instance_name="acc-test-mongo-replica"
    charge_type="PostPaid"
    project_name = "default"
    mongos_node_number = 32
    shard_number=3
    storage_space_gb=20
    subnet_id=volcengine_subnet.foo.id
    zone_id= data.volcengine_zones.foo.zones[0].id
    tags {
        key = "k1"
        value = "v1"
    }
}

resource "volcengine_mongodb_instance_parameter" "foo" {
    instance_id = volcengine_mongodb_instance.foo.id
    parameter_name = "cursorTimeoutMillis"
    parameter_role = "Node"
    parameter_value = "600111"
}
`

func TestAccVolcengineMongodbInstanceParameterResource_Basic(t *testing.T) {
	resourceName := "volcengine_mongodb_instance_parameter.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return instance_parameter.NewMongoDBInstanceParameterService(client)
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
				Config: testAccVolcengineMongodbInstanceParameterCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_name", "cursorTimeoutMillis"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_role", "Node"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_value", "600001"),
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

func TestAccVolcengineMongodbInstanceParameterResource_Update(t *testing.T) {
	resourceName := "volcengine_mongodb_instance_parameter.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return instance_parameter.NewMongoDBInstanceParameterService(client)
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
				Config: testAccVolcengineMongodbInstanceParameterCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_name", "cursorTimeoutMillis"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_role", "Node"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_value", "600001"),
				),
			},
			{
				Config: testAccVolcengineMongodbInstanceParameterUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_name", "cursorTimeoutMillis"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_role", "Node"),
					resource.TestCheckResourceAttr(acc.ResourceId, "parameter_value", "600111"),
				),
			},
			{
				Config:             testAccVolcengineMongodbInstanceParameterUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
