data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  name_regex = "veLinux 1.0 CentOS兼容版 64位"
}

resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-cluster"
  description               = "created by terraform"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids                       = [volcengine_subnet.foo.id]
    api_server_public_access_enabled = true
    api_server_public_access_config {
      public_access_network_config {
        billing_type = "PostPaidByBandwidth"
        bandwidth    = 1
      }
    }
    resource_public_access_default_enabled = true
  }
  pods_config {
    pod_network_mode = "VpcCniShared"
    vpc_cni_config {
      subnet_ids = [volcengine_subnet.foo.id]
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/18"]
  }
  tags {
    key   = "tf-k1"
    value = "tf-v1"
  }
}

resource "volcengine_vke_node_pool" "foo" {
  cluster_id = volcengine_vke_cluster.foo.id
  name       = "acc-test-node-pool"
  auto_scaling {
    enabled          = true
    min_replicas     = 0
    max_replicas     = 5
    desired_replicas = 0
    priority         = 5
    subnet_policy    = "ZoneBalance"
  }
  node_config {
    instance_type_ids = ["ecs.g1ie.xlarge"]
    subnet_ids        = [volcengine_subnet.foo.id]
    image_id          = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
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
    cordon = true
  }
  tags {
    key   = "node-pool-k1"
    value = "node-pool-v1"
  }
}

// add existing instances to a custom node pool
resource "volcengine_ecs_instance" "foo" {
    instance_name = "acc-test-ecs-${count.index}"
    host_name = "tf-acc-test"
    image_id = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
    instance_type = "ecs.g1ie.xlarge"
    password = "93f0cb0614Aab12"
    instance_charge_type = "PostPaid"
    system_volume_type = "ESSD_PL0"
    system_volume_size = 50
    data_volumes {
        volume_type = "ESSD_PL0"
        size = 50
        delete_with_instance = true
    }
    subnet_id = volcengine_subnet.foo.id
    security_group_ids = [volcengine_security_group.foo.id]
    project_name = "default"
    tags {
        key = "k1"
        value = "v1"
    }
    lifecycle {
        ignore_changes = [security_group_ids, tags]
    }
    count = 2
}

resource "volcengine_vke_node_pool" "foo1" {
    cluster_id = volcengine_vke_cluster.foo.id
    name       = "acc-test-node-pool"
    instance_ids = volcengine_ecs_instance.foo[*].id
    keep_instance_name = true
    node_config {
        instance_type_ids = ["ecs.g1ie.xlarge"]
        subnet_ids        = [volcengine_subnet.foo.id]
        image_id          = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
      system_volume {
        type = "ESSD_PL0"
        size = "50"
      }
      data_volumes {
        type = "ESSD_PL0"
        size = "50"
        mount_point = "/tf"
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
        cordon = true
    }
    tags {
        key   = "node-pool-k1"
        value = "node-pool-v1"
    }
}
