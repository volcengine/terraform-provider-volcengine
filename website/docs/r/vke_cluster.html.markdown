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
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  vpc_id              = volcengine_vpc.foo.id
  security_group_name = "acc-test-security-group2"
}

resource "volcengine_ecs_instance" "foo" {
  image_id             = "image-ybqi99s7yq8rx7mnk44b"
  instance_type        = "ecs.g1ie.large"
  instance_name        = "acc-test-ecs-name2"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  lifecycle {
    ignore_changes = [security_group_ids, instance_name]
  }
}

resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-1"
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
* `kubernetes_version` - (Optional, ForceNew) The version of Kubernetes specified when creating a VKE cluster (specified to patch version), if not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.
* `logging_config` - (Optional) Cluster log configuration information.
* `tags` - (Optional) Tags.

The `api_server_public_access_config` object supports the following:

* `public_access_network_config` - (Optional, ForceNew) Public network access network configuration.

The `cluster_config` object supports the following:

* `subnet_ids` - (Required, ForceNew) The subnet ID for the cluster control plane to communicate within the private network.
* `api_server_public_access_config` - (Optional) Cluster API Server public network access configuration.
* `api_server_public_access_enabled` - (Optional) Cluster API Server public network access configuration, the value is `true` or `false`.
* `resource_public_access_default_enabled` - (Optional, ForceNew) Node public network access configuration, the value is `true` or `false`.

The `flannel_config` object supports the following:

* `max_pods_per_node` - (Optional, ForceNew) The maximum number of single-node Pod instances for a Flannel container network, the value can be `16` or `32` or `64` or `128` or `256`.
* `pod_cidrs` - (Optional, ForceNew) Pod CIDR for the Flannel container network.

The `log_setups` object supports the following:

* `log_type` - (Required) The currently enabled log type.
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

