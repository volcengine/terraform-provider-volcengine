---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_instance_types"
sidebar_current: "docs-volcengine-datasource-ecs_instance_types"
description: |-
  Use this data source to query detailed information of ecs instance types
---
# volcengine_ecs_instance_types
Use this data source to query detailed information of ecs instance types
## Example Usage
```hcl
data "volcengine_ecs_instance_types" "foo" {}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of instance type IDs. When the number of ids is greater than 10, only the first 10 are effective.
* `image_id` - (Optional) The id of image.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_types` - The collection of query.
    * `baseline_credit` - The CPU benchmark performance that can be provided steadily by on-demand instances is determined by the instance type.
    * `gpu` - The GPU device info of Instance.
        * `gpu_devices` - GPU device information list.
            * `count` - The Count of GPU device.
            * `memory` - Graphics memory information.
                * `encrypted_size` - The Encrypted Memory Size of GPU device.
                * `size` - The Memory Size of GPU device.
            * `product_name` - The Product Name of GPU device.
    * `initial_credit` - The CPU credits obtained at once when creating a on-demand performance instance are fixed at 30 credits per vCPU.
    * `instance_type_family` - The instance type family.
    * `instance_type_id` - The id of the instance type.
    * `local_volumes` - Local disk configuration information corresponding to instance specifications.
        * `count` - The number of local disks mounted on the instance.
        * `size` - The size of volume.
        * `volume_type` - The type of volume.
    * `memory` - Memory information of instance specifications.
        * `encrypted_size` - The Encrypted Memory Size of GPU device.
        * `size` - Memory size, unit: MiB.
    * `network` - Network information of instance specifications.
        * `baseline_bandwidth_mbps` - Network benchmark bandwidth capacity (out/in), unit: Mbps.
        * `maximum_bandwidth_mbps` - Peak network bandwidth capacity (out/in), unit: Mbps.
        * `maximum_network_interfaces` - Maximum number of elastic network interfaces supported for attachment.
        * `maximum_private_ipv4_addresses_per_network_interface` - Maximum number of IPv4 addresses for a single elastic network interface.
        * `maximum_queues_per_network_interface` - Maximum queue number for a single elastic network interface, including the queue number supported by the primary network interface and the auxiliary network interface.
        * `maximum_throughput_kpps` - Network packet sending and receiving capacity (in+out), unit: Kpps.
    * `processor` - CPU information of instance specifications.
        * `base_frequency` - CPU clock speed, unit: GHz.
        * `cpus` - The number of ECS instance CPU cores.
        * `model` - CPU model.
        * `turbo_frequency` - CPU Turbo Boost, unit: GHz.
    * `rdma` - RDMA Specification Information.
        * `rdma_network_interfaces` - Number of RDMA network cards.
    * `volume` - Cloud disk information for instance specifications.
        * `maximum_count` - The maximum number of volumes.
        * `supported_volume_types` - List of supported volume types.
* `total_count` - The total count of query.


