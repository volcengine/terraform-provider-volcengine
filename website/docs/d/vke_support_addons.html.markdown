---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_support_addons"
sidebar_current: "docs-volcengine-datasource-vke_support_addons"
description: |-
  Use this data source to query detailed information of vke support addons
---
# volcengine_vke_support_addons
Use this data source to query detailed information of vke support addons
## Example Usage
```hcl
data "volcengine_vke_support_addons" "default" {
  name       = "metrics-server"
  categories = ["Monitor"]
}
```
## Argument Reference
The following arguments are supported:
* `categories` - (Optional) The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.
* `deploy_modes` - (Optional) The deploy model, the value is `Managed` or `Unmanaged`.
* `deploy_node_types` - (Optional) The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.
* `kubernetes_versions` - (Optional) A list of Kubernetes Versions.
* `name` - (Optional) The name of the addon.
* `necessaries` - (Optional) The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.
* `output_file` - (Optional) File name where to save data source results.
* `pod_network_modes` - (Optional) The container network model, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `addons` - The collection of addons query.
    * `deploy_mode` - The deploy model.
    * `deploy_node_types` - The deploy node types.
    * `name` - The name of addon.
    * `pod_network_modes` - The network modes of pod.
    * `versions` - The version info of addon.
        * `compatibilities` - The compatible version list.
            * `kubernetes_version` - The Kubernetes Version of addon.
        * `compatible_versions` - The compatible version list.
        * `version` - The basic version info.
* `total_count` - The total count of addons query.


