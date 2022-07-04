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
data "volcengine_vke_clusters" "default" {
  pods_config_pod_network_mode = "VpcCniShared"
  statuses {
    phase           = "Creating"
    conditions_type = "Progressing"
  }
}
```
## Argument Reference
The following arguments are supported:
* `create_client_token` - (Optional) ClientToken when successfully created.
* `delete_protection_enabled` - (Optional) The delete protection of the cluster.
* `ids` - (Optional) A list of Cluster IDs.
* `name_regex` - (Optional) A Name Regex of Cluster.
* `name` - (Optional) The name of the cluster.
* `output_file` - (Optional) File name where to save data source results.
* `page_number` - (Optional) The page number of clusters query.
* `page_size` - (Optional) The page size of clusters query.
* `pods_config_pod_network_mode` - (Optional) The network mode of the pod.
* `statuses` - (Optional) The statuses of the cluster.
* `update_client_token` - (Optional) ClientToken when the last update was successful.

The `statuses` object supports the following:

* `conditions_type` - (Optional) State conditions in the current primary state of the cluster.
* `phase` - (Optional) The status of cluster.

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
      * `access_source_ipsv4` - IPv4 public network access whitelist.
      * `public_access_network_config` - Public network access network configuration.
        * `bandwidth` - Peak bandwidth of public IP.
        * `billing_type` - Billing type of public IP.
        * `isp` - The ISP of public IP.
    * `api_server_public_access_enabled` - Cluster API Server public network access configuration.
    * `resource_public_access_default_enabled` - Node public network access configuration.
    * `security_group_ids` - The list of Security Group IDs.
    * `subnet_ids` - The list of Subnet IDs.
    * `vpc_id` - The VPC ID of the cluster control plane and the network of some nodes.
  * `create_time` - The create time of the Cluster.
  * `delete_protection_enabled` - The delete protection of the cluster.
  * `description` - The description of the cluster.
  * `eip_allocation_id` - Eip allocation Id.
  * `id` - The ID of the Cluster.
  * `kubeconfig_private` - Kubeconfig data with private network access, returned in BASE64 encoding.
  * `kubeconfig_public` - Kubeconfig data with public network access, returned in BASE64 encoding.
  * `kubernetes_version` - The version of Kubernetes specified when creating the VKE cluster.
  * `name` - The name of the cluster.
  * `node_statistics` - Statistics on the number of nodes corresponding to each master state in the cluster.
    * `creating_count` - Phase=Creating total number of nodes.
    * `deleting_count` - Phase=Deleting total number of nodes.
    * `failed_count` - Phase=Failed total number of nodes.
    * `running_count` - Phase=Running total number of nodes.
    * `stopped_count` - Phase=Stopped total number of nodes.
    * `total_count` - Total number of nodes.
    * `updating_count` - Phase=Updating total number of nodes.
  * `pods_config` - The config of the pods.
    * `flannel_config` - Flannel network configuration.
      * `max_pods_per_node` - Maximum number of Pod instances on a single node.
      * `pod_cidrs` - Container Pod Network CIDR.
    * `pod_network_mode` - Container Pod Network Type (CNI).
    * `vpc_cni_config` - VPC-CNI network configuration.
      * `subnet_ids` - List of subnets corresponding to the container Pod network.
      * `vpc_id` - Maximum number of Pod instances on a single node.
  * `services_config` - The config of the services.
    * `service_cidrsv4` - The IPv4 private network address exposed by the service.
  * `status` - The description of the cluster.
    * `conditions` - State conditions in the current primary state of the cluster.
      * `type` - State conditions in the current primary state of the cluster.
    * `phase` - The status of cluster.
  * `update_time` - The time the cluster was last admitted and executed/completed.
* `total_count` - The total count of Cluster query.


