---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_cluster"
sidebar_current: "docs-volcengine-resource-veecp_cluster"
description: |-
  Provides a resource to manage veecp cluster
---
# volcengine_veecp_cluster
Provides a resource to manage veecp cluster
## Example Usage
```hcl
resource "volcengine_veecp_cluster" "foo" {
  name = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_config` - (Required, ForceNew) Network configuration of cluster control plane and nodes.
* `name` - (Required, ForceNew) Cluster name. Under the same region, the name must be unique. Supports upper and lower case English letters, Chinese characters, numbers, and hyphens (-). Numbers cannot be at the first position, and hyphens (-) cannot be at the first or last position. The length is limited to 2 to 64 characters.
* `pods_config` - (Required, ForceNew) Container (Pod) network configuration of the cluster.
* `services_config` - (Required, ForceNew) Cluster service (Service) network configuration.
* `client_token` - (Optional, ForceNew) ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.
* `description` - (Optional, ForceNew) Cluster description. Length is limited to within 300 characters.
* `kubernetes_version` - (Optional, ForceNew) Specify the Kubernetes version when creating a cluster. The format is x.xx. The default value is the latest version in the supported Kubernetes version list (currently 1.20).
* `logging_config` - (Optional, ForceNew) Cluster log configuration information.
* `profile` - (Optional, ForceNew) Edge cluster: Edge. Non-edge cluster: Cloud. When using edge hosting, set this item to Edge.
* `tags` - (Optional) Tags.

The `api_server_public_access_config` object supports the following:

* `public_access_network_config` - (Optional, ForceNew) Public network access network configuration.

The `cluster_config` object supports the following:

* `subnet_ids` - (Required, ForceNew) The subnet ID for communication within the private network (VPC) of the cluster control plane. You can call the private network API to obtain the subnet ID. Note: When creating a cluster, please ensure that all specified SubnetIds (including but not limited to this parameter) belong to the same private network. It is recommended that you choose subnets in different availability zones as much as possible to improve the high availability of the cluster control plane. Please note that this parameter is not supported to be modified after the cluster is created. Please configure it reasonably.
* `api_server_public_access_config` - (Optional, ForceNew) Cluster API Server public network access configuration information. It takes effect only when ApiServerPublicAccessEnabled=true.
* `api_server_public_access_enabled` - (Optional, ForceNew) Cluster API Server public network access configuration, values:
false: (default value). closed
true: opened.
* `resource_public_access_default_enabled` - (Optional, ForceNew) Node public network access configuration, values:
false: (default value). Do not enable public network access. Existing NAT gateways and rules are not affected. true: Enable public network access. After enabling, a NAT gateway is automatically created for the cluster's private network and corresponding rules are configured. Note: This parameter cannot be modified after the cluster is created. Please configure it reasonably.

The `flannel_config` object supports the following:

* `pod_cidrs` - (Required, ForceNew) Pod CIDR of Flannel model container network. Only configurable when PodNetworkMode=Flannel, but not mandatory. Note: The number of Pods in the cluster is limited by the number of IPs in this CIDR. This parameter cannot be modified after cluster creation. Please plan the Pod CIDR reasonably. Cannot conflict with the following network segments: private network network segments corresponding to ClusterConfig.SubnetIds. All clusters within the same private network's FlannelConfig.PodCidrs. All clusters within the same private network's ServiceConfig.ServiceCidrsv4. Different clusters within the same private network's FlannelConfig.PodCidrs cannot conflict.
* `max_pods_per_node` - (Optional, ForceNew) Upper limit of the number of single-node Pod instances in the Flannel model container network. Values: 64(default value), 16, 32, 128, 256.

The `log_setups` object supports the following:

* `log_type` - (Required, ForceNew) The current types of logs that can be enabled are:
Audit: Cluster audit logs.
KubeApiServer: kube-apiserver component logs.
KubeScheduler: kube-scheduler component logs.
KubeControllerManager: kube-controller-manager component logs.
* `enabled` - (Optional, ForceNew) Whether to enable the log option, true means enable, false means not enable, the default is false. When Enabled is changed from false to true, a new Topic will be created.
* `log_ttl` - (Optional, ForceNew) The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.

The `logging_config` object supports the following:

* `log_project_id` - (Optional, ForceNew) The TLS log item ID of the collection target.
* `log_setups` - (Optional, ForceNew) Cluster logging options. This structure can only be modified and added, and cannot be deleted. When encountering a `cannot be deleted` error, please query the log setups of the current cluster and fill in the current `tf` file.

The `pods_config` object supports the following:

* `pod_network_mode` - (Required, ForceNew) Container network model, values: Flannel: Flannel network model, an independent Underlay container network solution. Combined with the global routing capability of a private network (VPC), it realizes a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution implemented based on the elastic network interface (ENI) of a private network, with high network communication performance. Description: After the cluster is created, this parameter is not supported to be modified temporarily. Please configure it reasonably.
* `flannel_config` - (Optional, ForceNew) Flannel network configuration. It can be configured only when PodNetworkMode=Flannel, but it is not mandatory.
* `vpc_cni_config` - (Optional, ForceNew) VPC-CNI network configuration. PodNetworkMode=VpcCniShared, but it is not mandatory.

The `public_access_network_config` object supports the following:

* `bandwidth` - (Optional) The peak bandwidth of the public IP, unit: Mbps.
* `billing_type` - (Optional) Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.

The `services_config` object supports the following:

* `service_cidrsv4` - (Required, ForceNew) CIDR used by services within the cluster. It cannot conflict with the following network segments: FlannelConfig.PodCidrs. SubnetIds of all clusters within the same private network or FlannelConfig.VpcConfig.SubnetIds. ServiceConfig.ServiceCidrsv4 of all clusters within the same private network (this parameter).It is stated that currently only one array element is supported. When multiple values are specified, only the first value takes effect.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `vpc_cni_config` object supports the following:

* `subnet_ids` - (Required, ForceNew) A list of Pod subnet IDs for the VPC-CNI container network.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `delete_protection_enabled` - Cluster deletion protection. Values: false: (default value) Deletion protection is off. true: Enable deletion protection. The cluster cannot be directly deleted. After creating a cluster, when calling Delete edge cluster, configure the Force parameter and choose to forcibly delete the cluster.


## Import
VeecpCluster can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_cluster.default resource_id
```

