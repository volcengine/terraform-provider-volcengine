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
resource "volcengine_vke_cluster" "foo" {
  name                      = "terraform-test-15"
  description               = "created by terraform"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids                       = ["subnet-2bzud0pbor8qo2dx0ee884y6h"]
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
    pod_network_mode = "Flannel"
    flannel_config {
      pod_cidrs         = ["172.27.224.0/19"]
      max_pods_per_node = 64
    }
    vpc_cni_config {
      subnet_ids = ["subnet-2bzud0pbor8qo2dx0ee884y6h"]
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/17"]
  }
}
```
## Argument Reference
The following arguments are supported:
* `cluster_config` - (Required) The config of the cluster.
* `name` - (Required) The name of the cluster.
* `pods_config` - (Required, ForceNew) The config of the pods.
* `services_config` - (Required, ForceNew) The config of the services.
* `client_token` - (Optional) ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.
* `delete_protection_enabled` - (Optional) The delete protection of the cluster.
* `description` - (Optional) The description of the cluster.
* `kubernetes_version` - (Optional, ForceNew) The version of Kubernetes specified when creating a VKE cluster (specified to patch version), if not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.

The `api_server_public_access_config` object supports the following:

* `public_access_network_config` - (Optional, ForceNew) Public network access network configuration.

The `cluster_config` object supports the following:

* `subnet_ids` - (Required, ForceNew) The list of Subnet IDs.
* `api_server_public_access_config` - (Optional) Cluster API Server public network access configuration.
* `api_server_public_access_enabled` - (Optional) Cluster API Server public network access configuration.
* `resource_public_access_default_enabled` - (Optional, ForceNew) Node public network access configuration.

The `flannel_config` object supports the following:

* `max_pods_per_node` - (Optional, ForceNew) Maximum number of Pod instances on a single node.
* `pod_cidrs` - (Optional, ForceNew) Container Pod Network CIDR.

The `pods_config` object supports the following:

* `pod_network_mode` - (Required, ForceNew) Container Pod Network Type (CNI).
* `flannel_config` - (Optional, ForceNew) Flannel network configuration.
* `vpc_cni_config` - (Optional, ForceNew) VPC-CNI network configuration.

The `public_access_network_config` object supports the following:

* `bandwidth` - (Optional) Peak bandwidth of public IP.
* `billing_type` - (Optional) Billing type of public IP.

The `services_config` object supports the following:

* `service_cidrsv4` - (Required, ForceNew) The IPv4 private network address exposed by the service.

The `vpc_cni_config` object supports the following:

* `subnet_ids` - (Optional, ForceNew) List of subnets corresponding to the container Pod network.
* `vpc_id` - (Optional, ForceNew) Maximum number of Pod instances on a single node.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `eip_allocation_id` - Eip allocation Id.
* `kubeconfig_private` - Kubeconfig data with private network access, returned in BASE64 encoding.
* `kubeconfig_public` - Kubeconfig data with public network access, returned in BASE64 encoding.


## Import
VkeCluster can be imported using the id, e.g.
```
$ terraform import volcengine_vke_cluster.default cc9l74mvqtofjnoj5****
```

