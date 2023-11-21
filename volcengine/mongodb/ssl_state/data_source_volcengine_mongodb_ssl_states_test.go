package ssl_state_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/ssl_state"
)

const testAccVolcengineMongodbSslStatesDatasourceConfig = `
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
    mongos_node_number = 2
    shard_number=3
    storage_space_gb=20
    subnet_id=volcengine_subnet.foo.id
    zone_id= data.volcengine_zones.foo.zones[0].id
    tags {
        key = "k1"
        value = "v1"
    }
}

resource "volcengine_mongodb_ssl_state" "foo" {
    instance_id = volcengine_mongodb_instance.foo.id
}

data "volcengine_mongodb_ssl_states" "foo"{
    instance_id = volcengine_mongodb_instance.foo.id
}
`

func TestAccVolcengineMongodbSslStatesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_mongodb_ssl_states.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ssl_state.NewMongoDBSSLStateService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineMongodbSslStatesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "ssl_state.#", "1"),
				),
			},
		},
	})
}
