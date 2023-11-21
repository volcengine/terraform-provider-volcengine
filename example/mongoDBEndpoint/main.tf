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

resource "volcengine_eip_address" "foo" {
    billing_type = "PostPaidByBandwidth"
    bandwidth = 1
    isp = "ChinaUnicom"
    name = "acc-eip-${count.index}"
    description = "acc-test"
    project_name = "default"
    count = 2
}

resource "volcengine_mongodb_instance" "replica-set"{
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

resource "volcengine_mongodb_endpoint" "replica-set-public-endpoint"{
    instance_id = volcengine_mongodb_instance.replica-set.id
    network_type = "Public"
    eip_ids = volcengine_eip_address.foo[*].id
}

resource "volcengine_mongodb_instance" "sharded-cluster"{
    db_engine_version = "MongoDB_4_0"
    instance_type="ShardedCluster"
    super_account_password="@acc-test-123"
    node_spec="mongo.shard.1c2g"
    mongos_node_spec="mongo.mongos.1c2g"
    instance_name="acc-test-mongo-shard"
    charge_type="PostPaid"
    project_name = "default"
    mongos_node_number = 2
    shard_number=2
    storage_space_gb=20
    subnet_id=volcengine_subnet.foo.id
    zone_id= data.volcengine_zones.foo.zones[0].id
    tags {
        key = "k1"
        value = "v1"
    }
}

resource "volcengine_mongodb_endpoint" "sharded-cluster-private-endpoint"{
    instance_id = volcengine_mongodb_instance.sharded-cluster.id
    network_type = "Private"
    object_id = volcengine_mongodb_instance.sharded-cluster.shards[0].shard_id
}
