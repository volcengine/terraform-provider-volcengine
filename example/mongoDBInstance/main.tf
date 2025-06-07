# query available zones in current region
data "volcengine_zones" "foo" {
}

# create vpc
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  dns_servers  = ["8.8.8.8", "114.114.114.114"]
  project_name = "default"
}

# create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

# create mongodb ReplicaSet instance
resource "volcengine_mongodb_instance" "foo-replica" {
  zone_ids               = [data.volcengine_zones.foo.zones[0].id]
  db_engine_version      = "MongoDB_4_0"
  instance_type          = "ReplicaSet"
  node_spec              = "mongo.2c4g"
  storage_space_gb       = 100
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test-mongodb-replica"
  charge_type            = "PostPaid"
  super_account_password = "93f0cb0614Aab12"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  node_availability_zone {
    zone_id     = data.volcengine_zones.foo.zones[0].id
    node_number = 2
  }
}

# create mongodb ShardedCluster instance
resource "volcengine_mongodb_instance" "foo-sharded" {
  zone_ids                       = [data.volcengine_zones.foo.zones[0].id]
  db_engine_version              = "MongoDB_4_0"
  instance_type                  = "ShardedCluster"
  node_spec                      = "mongo.shard.2c4g"
  mongos_node_spec               = "mongo.mongos.2c4g"
  mongos_node_number             = 3
  shard_number                   = 3
  config_server_node_spec        = "mongo.config.2c4g"
  config_server_storage_space_gb = 30
  storage_space_gb               = 100
  subnet_id                      = volcengine_subnet.foo.id
  instance_name                  = "acc-test-mongodb-sharded"
  charge_type                    = "PostPaid"
  super_account_password         = "93f0cb0614Aab12"
  project_name                   = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  node_availability_zone {
    zone_id     = data.volcengine_zones.foo.zones[0].id
    node_number = 2
  }
}
