---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_cloud_servers"
sidebar_current: "docs-volcengine-datasource-veenedge_cloud_servers"
description: |-
  Use this data source to query detailed information of veenedge cloud servers
---
# volcengine_veenedge_cloud_servers
Use this data source to query detailed information of veenedge cloud servers
## Example Usage
```hcl
data "volcengine_veenedge_cloud_servers" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of cloud server IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cloud_servers` - The collection of cloud servers query.
    * `billing_config` - The config of billing.
        * `bandwidth_billing_method` - The bandwidth billing method.
        * `computing_billing_method` - The computing billing method.
    * `cloud_server_identity` - The Id of cloud server.
    * `cpu` - The cpu info of cloud server.
    * `create_time` - The create time info.
    * `custom_data` - The config of custom data.
        * `data` - The data info.
    * `gpu` - The config of gpu.
        * `gpus` - The list gpu info.
            * `gpu_spec` - The spec of gpu.
                * `gpu_type` - The type of gpu.
            * `num` - The number of gpu.
    * `id` - The Id of cloud server.
    * `image` - The config of image.
        * `image_identity` - The id of image.
        * `image_name` - The name of image.
        * `property` - The property of system.
        * `system_arch` - The arch of system.
        * `system_bit` - The bit of system.
        * `system_type` - The type of system.
        * `system_version` - The version of system.
    * `instance_count` - The count of instances.
    * `instance_status` - The status of instances.
        * `instance_count` - The count of instance.
        * `status` - The status info.
    * `mem` - The memory info of cloud server.
    * `name` - The name of cloud server.
    * `network` - The config of network.
        * `bandwidth_peak` - The peak of bandwidth.
        * `enable_ipv6` - Whether enable ipv6.
        * `internal_bandwidth_peak` - The internal peak of bandwidth.
    * `schedule_strategy_configs` - The config of schedule strategy.
        * `price_strategy` - The price strategy.
        * `schedule_strategy` - The schedule strategy.
    * `secret_config` - The config of secret.
        * `secret_data` - The data of secret.
        * `secret_type` - The type of secret.
    * `server_area_count` - The server area count number.
    * `server_area_level` - The area level of cloud server.
    * `server_areas` - The server areas info.
        * `area` - The area info.
        * `instance_num` - The number of instance.
        * `isp` - The isp info.
    * `spec_display` - The Chinese spec info of cloud server.
    * `spec_sum` - The spec summary of cloud server.
    * `spec` - The spec info of cloud server.
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
    * `update_time` - The update time info.
* `total_count` - The total count of cloud servers query.


