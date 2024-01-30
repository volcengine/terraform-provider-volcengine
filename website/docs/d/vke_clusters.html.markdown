---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_clusters"
sidebar_current: "docs-volcengine-datasource-vke_clusters"
description: |-
  Use this data source to query detailed information of vke clusters
---
# volcengine_vke_clusters
Use this data source to query detailed information of vke clusters
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

data "volcengine_vke_clusters" "foo" {
  ids = [volcengine_vke_cluster.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `create_client_token` - (Optional) ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
* `delete_protection_enabled` - (Optional) The delete protection of the cluster, the value is `true` or `false`.
* `ids` - (Optional) A list of Cluster IDs.
* `name_regex` - (Optional) A Name Regex of Cluster.
* `name` - (Optional) The name of the cluster.
* `output_file` - (Optional) File name where to save data source results.
* `page_number` - (Optional) The page number of clusters query.
* `page_size` - (Optional) The page size of clusters query.
* `pods_config_pod_network_mode` - (Optional) The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
* `statuses` - (Optional) Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
* `tags` - (Optional) Tags.
* `update_client_token` - (Optional) The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.

The `statuses` object supports the following:

* `conditions_type` - (Optional) The state condition in the current main state of the cluster, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Degraded`, `SetByProvider`, `Balance`, `Security`, `CreateError`, `ResourceCleanupFailed`, `LimitedByQuota`, `StockOut`,`Unknown`.
* `phase` - (Optional) The status of cluster. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Stopped`, `Failed`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `clusters` - The collection of VkeCluster query.
    * `cluster_config` - The config of the cluster.
        * `api_server_endpoints` - Endpoint information accessed by the cluster API Server.
            * `private_ip` - Endpoint address of the cluster API Server private network.
                * `ipv4` - Ipv4 address.
            * `public_ip` - Endpoint address of the cluster API Server public network.
                * `ipv4` - Ipv4 address.
        * `api_server_public_access_config` - Cluster API Server public network access configuration.
            * `access_source_ipsv4` - IPv4 public network access whitelist. A null value means all network segments (0.0.0.0/0) are allowed to pass.
            * `public_access_network_config` - Public network access network configuration.
                * `bandwidth` - The peak bandwidth of the public IP, unit: Mbps.
                * `billing_type` - Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.
                * `isp` - The ISP of public IP.
        * `api_server_public_access_enabled` - Cluster API Server public network access configuration, the value is `true` or `false`.
        * `resource_public_access_default_enabled` - Node public network access configuration, the value is `true` or `false`.
        * `security_group_ids` - The security group used by the cluster control plane and nodes.
        * `subnet_ids` - The subnet ID for the cluster control plane to communicate within the private network.
        * `vpc_id` - The ID of the private network (VPC) where the network of the cluster control plane and some nodes is located.
    * `create_time` - Cluster creation time. UTC+0 time in standard RFC3339 format.
    * `delete_protection_enabled` - The delete protection of the cluster, the value is `true` or `false`.
    * `description` - The description of the cluster.
    * `eip_allocation_id` - Eip allocation Id.
    * `id` - The ID of the Cluster.
    * `kubeconfig_private` - Kubeconfig data with private network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
    * `kubeconfig_public` - Kubeconfig data with public network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.
    * `kubernetes_version` - The Kubernetes version information corresponding to the cluster, specific to the patch version.
    * `logging_config` - Cluster log configuration information.
        * `log_project_id` - The TLS log item ID of the collection target.
        * `log_setups` - Cluster logging options.
            * `enabled` - Whether to enable the log option, true means enable, false means not enable, the default is false. When Enabled is changed from false to true, a new Topic will be created.
            * `log_ttl` - The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.
            * `log_type` - The currently enabled log type.
    * `name` - The name of the cluster.
    * `node_statistics` - Statistics on the number of nodes corresponding to each master state in the cluster.
        * `creating_count` - Phase=Creating total number of nodes.
        * `deleting_count` - Phase=Deleting total number of nodes.
        * `failed_count` - Phase=Failed total number of nodes.
        * `running_count` - Phase=Running total number of nodes.
        * `stopped_count` - (**Deprecated**) This field has been deprecated and is not recommended for use. Phase=Stopped total number of nodes.
        * `total_count` - Total number of nodes.
        * `updating_count` - Phase=Updating total number of nodes.
    * `pods_config` - The config of the pods.
        * `flannel_config` - Flannel network configuration.
            * `max_pods_per_node` - The maximum number of single-node Pod instances for a Flannel container network.
            * `pod_cidrs` - Pod CIDR for the Flannel container network.
        * `pod_network_mode` - Container Pod Network Type (CNI), the value is `Flannel` or `VpcCniShared`.
        * `vpc_cni_config` - VPC-CNI network configuration.
            * `subnet_ids` - A list of Pod subnet IDs for the VPC-CNI container network.
            * `vpc_id` - The private network where the cluster control plane network resides.
    * `services_config` - The config of the services.
        * `service_cidrsv4` - The IPv4 private network address exposed by the service.
    * `status` - The status of the cluster.
        * `conditions` - The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.
            * `type` - The state condition in the current main state of the cluster, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Balance`, `CreateError`, `ResourceCleanupFailed`, `Unknown`.
        * `phase` - The status of cluster. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Stopped`, `Failed`.
    * `tags` - Tags of the Cluster.
        * `key` - The Key of Tags.
        * `type` - The Type of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The last time a request was accepted by the cluster and executed or completed. UTC+0 time in standard RFC3339 format.
* `total_count` - The total count of Cluster query.


