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

resource "volcengine_mongodb_allow_list" "foo" {
    allow_list_name="acc-test"
    allow_list_desc="acc-test"
    allow_list_type="IPv4"
    allow_list="10.1.1.3,10.2.3.0/24,10.1.1.1"
}

resource "volcengine_mongodb_allow_list_associate" "foo" {
    allow_list_id = volcengine_mongodb_allow_list.foo.id
    instance_id = volcengine_mongodb_instance.foo.id
}