resource "volcengine_vpc" "foo" {
  vpc_name = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block = "172.16.0.0/24"
  zone_id = "cn-beijing-a"
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  vpc_id = volcengine_vpc.foo.id
  security_group_name = "acc-test-security-group2"
}

resource "volcengine_ecs_instance" "foo" {
  image_id = "image-ybqi99s7yq8rx7mnk44b"
  instance_type = "ecs.g1ie.large"
  instance_name = "acc-test-ecs-name2"
  password = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type = "ESSD_PL0"
  system_volume_size = 40
  subnet_id = volcengine_subnet.foo.id
  security_group_ids = [volcengine_security_group.foo.id]
  lifecycle {
    ignore_changes = [security_group_ids, instance_name]
  }
  project_name = "default"
}

resource "volcengine_vke_cluster" "foo" {
  name = "acc-test-1"
  description = "created by terraform"
  delete_protection_enabled = false
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
    pod_network_mode = "VpcCniShared"
    vpc_cni_config {
      subnet_ids = [volcengine_subnet.foo.id]
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/18"]
  }
  tags {
    key = "tf-k1"
    value = "tf-v1"
  }
}

resource "volcengine_vke_node_pool" "foo" {
  cluster_id = volcengine_vke_cluster.foo.id
  name = "acc-tf-test"
  node_config {
    instance_type_ids = ["ecs.g1ie.large"]
    subnet_ids = [volcengine_subnet.foo.id]
    security {
      login {
        #      ssh_key_pair_name = "ssh-6fbl66fxqm"
        password = "UHdkMTIzNDU2"
      }
      security_group_ids = [volcengine_security_group.foo.id]
    }
    instance_charge_type = "PostPaid"
    period = 1
  }
  kubernetes_config {
    labels {
      key   = "aa"
      value = "bb"
    }
    labels {
      key   = "cccc"
      value = "dddd"
    }
    cordon = false
  }
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_vke_node" "foo" {
  cluster_id = volcengine_vke_cluster.foo.id
  instance_id = volcengine_ecs_instance.foo.id
  keep_instance_name = true
  additional_container_storage_enabled = false
  container_storage_path = ""
  node_pool_id = volcengine_vke_node_pool.foo.id
  kubernetes_config {
    labels {
      key   = "tf-key1"
      value = "tf-value1"
    }
    labels {
      key   = "tf-key2"
      value = "tf-value2"
    }
    taints {
      key = "tf-key3"
      value = "tf-value3"
      effect = "NoSchedule"
    }
    taints {
      key = "tf-key4"
      value = "tf-value4"
      effect = "NoSchedule"
    }
    cordon = true
  }
}