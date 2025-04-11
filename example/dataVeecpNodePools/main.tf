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

resource "volcengine_security_group" "foo" {
    vpc_id = volcengine_vpc.foo.id
    security_group_name = "acc-test-security-group2"
}

resource "volcengine_veecp_cluster" "foo" {
    name = "acc-test-1"
    description = "created by terraform"
    delete_protection_enabled = false
    profile = "Edge"
    cluster_config {
        subnet_ids = [volcengine_subnet.foo.id]
        api_server_public_access_enabled = true
        api_server_public_access_config {
            public_access_network_config {
                billing_type = "PostPaidByBandwidth"
                bandwidth = 1
            }
        }
        resource_public_access_default_enabled = true
    }
    pods_config {
        pod_network_mode = "Flannel"
        flannel_config {
            pod_cidrs = ["172.22.224.0/20"]
            max_pods_per_node = 64
        }
    }
    services_config {
        service_cidrsv4 = ["172.30.0.0/18"]
    }
}

resource "volcengine_veecp_node_pool" "foo" {
    cluster_id = volcengine_veecp_cluster.foo.id
    name       = "acc-test-node-pool-9505"
    client_token = "FGAHIxa23412FGAIOHioj"
    auto_scaling {
        enabled          = true
        min_replicas     = 0
        max_replicas     = 5
        desired_replicas = 0
        priority         = 5
        subnet_policy    = "ZoneBalance"
    }
    node_config {
        instance_type_ids = ["ecs.c1ie.xlarge"]
        subnet_ids        = [volcengine_subnet.foo.id]
        image_id          = ""
        system_volume {
            type = "ESSD_PL0"
            size = 80
        }
        data_volumes {
            type        = "ESSD_PL0"
            size        = 80
            mount_point = "/tf1"
        }
        data_volumes {
            type        = "ESSD_PL0"
            size        = 60
            mount_point = "/tf2"
        }
        initialize_script = "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"
        security {
            login {
                password = "UHdkMTIzNDU2"
            }
            security_strategies = ["Hids"]
            security_group_ids  = [volcengine_security_group.foo.id]
        }
        additional_container_storage_enabled = false
        instance_charge_type                 = "PostPaid"
        name_prefix                          = "acc-test"
        ecs_tags {
            key   = "ecs_k1"
            value = "ecs_v1"
        }
    }
    kubernetes_config {
        labels {
            key   = "label1"
            value = "value1"
        }
        taints {
            key    = "taint-key/node-type"
            value  = "taint-value"
            effect = "NoSchedule"
        }
        cordon             = true
        #auto_sync_disabled = false
    }
}

data "volcengine_veecp_node_pools" "foo"{
    ids = [volcengine_veecp_node_pool.foo.id]
}