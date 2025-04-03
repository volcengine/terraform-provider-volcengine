#data "volcengine_zones" "foo"{
#}
#
#resource "volcengine_vpc" "foo" {
#    vpc_name = "acc-test-project1"
#    cidr_block = "172.16.0.0/16"
#}
#
#resource "volcengine_subnet" "foo" {
#    subnet_name = "acc-subnet-test-2"
#    cidr_block = "172.16.0.0/24"
#    zone_id = data.volcengine_zones.foo.zones[0].id
#    vpc_id = volcengine_vpc.foo.id
#}
#
#resource "volcengine_security_group" "foo" {
#    vpc_id = volcengine_vpc.foo.id
#    security_group_name = "acc-test-security-group2"
#}
#
#resource "volcengine_veecp_cluster" "foo" {
#    name = "acc-test-1"
#    description = "created by terraform"
#    delete_protection_enabled = false
#    profile = "Edge"
#    cluster_config {
#        subnet_ids = [volcengine_subnet.foo.id]
#        api_server_public_access_enabled = true
#        api_server_public_access_config {
#            public_access_network_config {
#                billing_type = "PostPaidByBandwidth"
#                bandwidth = 1
#            }
#        }
#        resource_public_access_default_enabled = true
#    }
#    pods_config {
#        pod_network_mode = "Flannel"
#        flannel_config {
#            pod_cidrs = ["172.22.224.0/20"]
#            max_pods_per_node = 64
#        }
#    }
#    services_config {
#        service_cidrsv4 = ["172.30.0.0/18"]
#    }
#}

resource "volcengine_veecp_edge_node_pool" "foo" {
    cluster_id = "ccvmb0c66t101fnob3dhg"
    name = "acc-test-tf"
    node_pool_type = "edge-machine-pool"
    vpc_id = "vpc-l9sz9qlf2t"
#    billing_configs {
#
#    }
    elastic_config {
        cloud_server_identity = "cloudserver-47vz7k929cp9xqb"
        auto_scale_config {
            enabled = true
            max_replicas = 2
            desired_replicas = 0
            min_replicas = 0
            priority = 10
        }
        instance_area {
            cluster_name = "bdcdn-zzcu02"
            vpc_identity = "vpc-l9sz9qlf2t"
        }
    }
}