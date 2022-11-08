---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_instances"
sidebar_current: "docs-volcengine-datasource-veenedge_instances"
description: |-
  Use this data source to query detailed information of veenedge instances
---
# volcengine_veenedge_instances
Use this data source to query detailed information of veenedge instances
## Example Usage
```hcl
data "volcengine_veenedge_instances" "default" {
  ids = ["veen28*****21", "veen177110*****172"]
}
```
## Argument Reference
The following arguments are supported:
* `cloud_server_ids` - (Optional) The list of cloud server ids.
* `ids` - (Optional) A list of instance IDs.
* `names` - (Optional) A list of instance names.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The list of instance status. The value can be `opening` or `starting` or `running` or `stopping` or `stop` or `rebooting` or `terminating`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of instance query.
    * `cloud_server_identity` - The id of cloud server.
    * `cloud_server_name` - The name of cloud server.
    * `cluster` - The cluster info.
        * `alias` - The alias of cluster.
        * `city` - The city of cluster.
        * `cluster_name` - The name of cluster.
        * `country` - The country of cluster.
        * `isp` - The isp of cluster.
        * `level` - The level of cluster.
        * `province` - The province of cluster.
        * `region` - The region of cluster.
    * `cpu` - The cpu of instance.
    * `create_time` - The create time of instance.
    * `creator` - The creator of instance.
    * `delete_time` - The delete time of instance.
    * `end_time` - The end time of instance.
    * `gpu` - The config of gpu.
        * `gpus` - The list gpu info.
            * `gpu_spec` - The spec of gpu.
                * `gpu_type` - The type of gpu.
            * `num` - The number of gpu.
    * `id` - The Id of instance.
    * `image` - The config of image.
        * `image_identity` - The id of image.
        * `image_name` - The name of image.
        * `property` - The property of system.
        * `system_arch` - The arch of system.
        * `system_bit` - The bit of system.
        * `system_type` - The type of system.
        * `system_version` - The version of system.
    * `instance_identity` - The Id of instance.
    * `instance_name` - The name of instance.
    * `mem` - The memory of instance.
    * `network` - The config of network.
        * `enable_ipv6` - Whether enable ipv6.
        * `external_interface` - The external interface of network.
            * `bandwidth_peak` - The peak of bandwidth.
            * `ip6_addr` - The ipv6 address.
            * `ip_addr` - The ip address.
            * `ips` - The ips of network.
                * `addr` - The ip address.
                * `ip_version` - The version of ip address.
                * `isp` - The isp info.
                * `mask` - The mask info.
            * `mac_addr` - The mac address.
            * `mask6` - The ipv6 mask info.
            * `mask` - The mask info.
        * `internal_interface` - The internal interface of network.
            * `bandwidth_peak` - The peak of bandwidth.
            * `ip6_addr` - The ipv6 address.
            * `ip_addr` - The ip address.
            * `ips` - The ips of network.
                * `addr` - The ip address.
                * `ip_version` - The version of ip address.
                * `isp` - The isp info.
                * `mask` - The mask info.
            * `mac_addr` - The mac address.
            * `mask6` - The ipv6 mask info.
            * `mask` - The mask info.
        * `vf_passthrough` - The passthrough info.
    * `spec_display` - The spec display of instance.
    * `spec` - The spec of instance.
    * `start_time` - The start time of instance.
    * `status` - The status of instance.
    * `storage` - The config of storage.
        * `data_disk_list` - The disk list info of data.
            * `capacity` - The capacity of storage.
            * `storage_type` - The type of storage.
        * `data_disk` - The disk info of data.
            * `capacity` - The capacity of storage.
            * `storage_type` - The type of storage.
        * `system_disk` - The disk info of system.
            * `capacity` - The capacity of storage.
            * `storage_type` - The type of storage.
    * `subnet_cidr` - The subnet cidr.
    * `update_time` - The update time of instance.
    * `vpc_identity` - The id of vpc.
* `total_count` - The total count of instance query.


