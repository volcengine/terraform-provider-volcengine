---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_cluster"
sidebar_current: "docs-volcengine-resource-vke_cluster"
description: |-
  Provides a resource to manage vke cluster
---
# volcengine_vke_cluster
Provides a resource to manage vke cluster
## Example Usage
```hcl
# query available zones in current region
data "volcengine_zones" "foo" {
}

# create vpc
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

# create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

# create security group
resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

# create vke cluster
resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-1"
  description               = "created by terraform"
  project_name              = "default"
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

# query the image_id which match the specified image_name
data "volcengine_images" "foo" {
  name_regex = "veLinux 1.0 CentOS Compatible 64 bit"
}

# create vke node pool
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
    image_id          = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS Compatible 64 bit"][0]
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
    project_name                         = "default"
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
    auto_sync_disabled = false
  }
  tags {
    key   = "node-pool-k1"
    value = "node-pool-v1"
  }
}

# create ecs instance
resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs"
  host_name            = "tf-acc-test"
  image_id             = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS Compatible 64 bit"][0]
  instance_type        = "ecs.g1ie.xlarge"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 50
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  lifecycle {
    ignore_changes = [security_group_ids, tags, instance_name]
  }
}

# add the ecs instance to the vke node pool
resource "volcengine_vke_node" "foo" {
  cluster_id   = volcengine_vke_cluster.foo.id
  instance_id  = volcengine_ecs_instance.foo.id
  node_pool_id = volcengine_vke_node_pool.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `cluster_config` - (Required) The config of the cluster.
* `name` - (Required) The name of the cluster.
* `pods_config` - (Required) The config of the pods.
* `services_config` - (Required, ForceNew) The config of the services.
* `client_token` - (Optional) ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.
* `delete_protection_enabled` - (Optional) The delete protection of the cluster, the value is `true` or `false`.
* `description` - (Optional) The description of the cluster.
* `kubernetes_version` - (Optional, ForceNew) The version of Kubernetes specified when creating a VKE cluster (specified to patch version), with an example value of `1.24`. If not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.
* `logging_config` - (Optional) Cluster log configuration information.
* `project_name` - (Optional) The project name of the cluster.
* `tags` - (Optional) Tags.

The `api_server_public_access_config` object supports the following:

* `public_access_network_config` - (Optional, ForceNew) Public network access network configuration.

The `cluster_config` object supports the following:

* `subnet_ids` - (Required) The subnet ID for the cluster control plane to communicate within the private network.
Up to 3 subnets can be selected from each available zone, and a maximum of 2 subnets can be added to each available zone.
Cannot support deleting configured subnets.
* `api_server_public_access_config` - (Optional) Cluster API Server public network access configuration.
* `api_server_public_access_enabled` - (Optional) Cluster API Server public network access configuration, the value is `true` or `false`.
* `resource_public_access_default_enabled` - (Optional, ForceNew) Node public network access configuration, the value is `true` or `false`.

The `flannel_config` object supports the following:

* `max_pods_per_node` - (Optional, ForceNew) The maximum number of single-node Pod instances for a Flannel container network, the value can be `16` or `32` or `64` or `128` or `256`.
* `pod_cidrs` - (Optional, ForceNew) Pod CIDR for the Flannel container network.

The `log_setups` object supports the following:

* `log_type` - (Required) The current types of logs that can be enabled are:
Audit: Cluster audit logs.
KubeApiServer: kube-apiserver component logs.
KubeScheduler: kube-scheduler component logs.
KubeControllerManager: kube-controller-manager component logs.
* `enabled` - (Optional) Whether to enable the log option, true means enable, false means not enable, the default is false. When Enabled is changed from false to true, a new Topic will be created.
* `log_ttl` - (Optional) The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.

The `logging_config` object supports the following:

* `log_project_id` - (Optional) The TLS log item ID of the collection target.
* `log_setups` - (Optional) Cluster logging options. This structure can only be modified and added, and cannot be deleted. When encountering a `cannot be deleted` error, please query the log setups of the current cluster and fill in the current `tf` file.

The `pods_config` object supports the following:

* `pod_network_mode` - (Required, ForceNew) The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
* `flannel_config` - (Optional, ForceNew) Flannel network configuration.
* `vpc_cni_config` - (Optional) VPC-CNI network configuration.

The `public_access_network_config` object supports the following:

* `bandwidth` - (Optional) The peak bandwidth of the public IP, unit: Mbps.
* `billing_type` - (Optional) Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.

The `services_config` object supports the following:

* `service_cidrsv4` - (Required, ForceNew) The IPv4 private network address exposed by the service.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `vpc_cni_config` object supports the following:

* `subnet_ids` - (Optional) A list of Pod subnet IDs for the VPC-CNI container network.
* `vpc_id` - (Optional, ForceNew) The private network where the cluster control plane network resides.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `eip_allocation_id` - Eip allocation Id.
* `kubeconfig_private` - Kubeconfig data with private network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
* `kubeconfig_public` - Kubeconfig data with public network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.


## Import
VkeCluster can be imported using the id, e.g.
```
$ terraform import volcengine_vke_cluster.default cc9l74mvqtofjnoj5****
```

