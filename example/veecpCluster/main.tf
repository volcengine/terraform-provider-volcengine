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
    name = "acc-test-2"
    description = "created by terraform"
    delete_protection_enabled = false
    kubernetes_version = "v1.24.15-veecp.1"
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
