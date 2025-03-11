---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_clusters"
sidebar_current: "docs-volcengine-datasource-veecp_clusters"
description: |-
  Use this data source to query detailed information of veecp clusters
---
# volcengine_veecp_clusters
Use this data source to query detailed information of veecp clusters
## Example Usage
```hcl
data "volcengine_veecp_clusters" "foo" {
  create_client_token          = ""
  delete_protection_enabled    = true
  ids                          = []
  name                         = ""
  pods_config_pod_network_mode = ""
  profiles                     = []
  update_client_token          = ""
}
```
## Argument Reference
The following arguments are supported:
* `create_client_token` - (Optional) ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
* `delete_protection_enabled` - (Optional) Cluster deletion protection. Values: true: Enable deletion protection. false: Disable deletion protection.
* `ids` - (Optional) Cluster ID. Supports exact matching. A maximum of 100 array elements can be filled in at a time. Note: When this parameter is an empty array, filtering is based on all clusters in the specified region under the account.
* `name_regex` - (Optional) A Name Regex of Cluster.
* `name` - (Optional) Cluster name.
* `output_file` - (Optional) File name where to save data source results.
* `pods_config_pod_network_mode` - (Optional) The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
* `profiles` - (Optional) Filter by cluster scenario: Cloud: non-edge cluster; Edge: edge cluster.
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
* `clusters` - The collection of query.
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
    * `create_client_token` - ClientToken when creation is successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
    * `create_time` - Cluster creation time. UTC+0 time in standard RFC3339 format.
    * `delete_protection_enabled` - The delete protection of the cluster, the value is `true` or `false`.
    * `description` - Cluster description information.
    * `id` - The ID of the cluster.
    * `kubernetes_version` - Kubernetes version information corresponding to the cluster, specific to the patch version.
    * `logging_config` - Cluster log configuration information.
        * `log_project_id` - The TLS log item ID of the collection target.
        * `log_setups` - Cluster logging options.
            * `enabled` - Whether to enable the log option, true means enable, false means not enable, the default is false. When Enabled is changed from false to true, a new Topic will be created.
            * `log_ttl` - The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.
            * `log_type` - The currently enabled log type.
    * `name` - Cluster name.
    * `node_statistics` - Statistics on the number of nodes corresponding to each master state in the cluster.
        * `creating_count` - Phase=Creating total number of nodes.
        * `deleting_count` - Phase=Deleting total number of nodes.
        * `failed_count` - Phase=Failed total number of nodes.
        * `running_count` - Phase=Running total number of nodes.
        * `starting_count` - Phase=Starting total number of nodes.
        * `stopped_count` - (**Deprecated**) This field has been deprecated and is not recommended for use. Phase=Stopped total number of nodes.
        * `stopping_count` - Phase=Stopping total number of nodes.
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
    * `status` - Cluster status. For detailed instructions, please refer to ClusterStatusResponse.
        * `conditions` - The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.
            * `type` - The state condition in the current main state of the cluster, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Balance`, `CreateError`, `ResourceCleanupFailed`, `Unknown`.
        * `phase` - Cluster status. The value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.
    * `tags` - Tags of the Cluster.
        * `key` - The Key of Tags.
        * `type` - The Type of Tags.
        * `value` - The Value of Tags.
    * `update_client_token` - ClientToken when the last update was successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
    * `update_time` - The time when the cluster last accepted a request and executed or completed execution. UTC+0 time in standard RFC3339 format.
* `total_count` - The total count of query.


