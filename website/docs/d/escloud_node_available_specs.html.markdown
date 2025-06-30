---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_node_available_specs"
sidebar_current: "docs-volcengine-datasource-escloud_node_available_specs"
description: |-
  Use this data source to query detailed information of escloud node available specs
---
# volcengine_escloud_node_available_specs
Use this data source to query detailed information of escloud node available specs
## Example Usage
```hcl
data "volcengine_escloud_node_available_specs" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The id of the instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `node_specs` - The collection of query.
    * `az_available_specs_sold_out` - The available specs sold out.
    * `configuration_code` - The configuration code.
    * `network_specs` - The network specs.
        * `network_role` - The network role.
        * `spec_name` - The spec name.
    * `node_available_specs` - The node available specs.
        * `resource_spec_names` - The resource spec names of node.
        * `storage_spec_names` - The storage spec names of node.
        * `type` - The type of node.
    * `resource_specs` - The resource specs.
        * `cpu` - The cpu of resource spec. Unit: Core.
        * `description` - The description of resource spec.
        * `display_name` - The display name of resource spec.
        * `memory` - The memory of resource spec. Unit: GiB.
        * `name` - The name of resource spec.
    * `storage_specs` - The storage specs.
        * `description` - The description of storage spec.
        * `display_name` - The display name of storage spec.
        * `max_size` - The max size of storage spec. Unit: GiB.
        * `min_size` - The min size of storage spec. Unit: GiB.
        * `name` - The name of storage spec.
        * `size` - The size of storage spec.
* `total_count` - The total count of query.


