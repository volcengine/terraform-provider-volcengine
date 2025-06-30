---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_resolver_endpoints"
sidebar_current: "docs-volcengine-datasource-private_zone_resolver_endpoints"
description: |-
  Use this data source to query detailed information of private zone resolver endpoints
---
# volcengine_private_zone_resolver_endpoints
Use this data source to query detailed information of private zone resolver endpoints
## Example Usage
```hcl
data "volcengine_private_zone_resolver_endpoints" "foo" {}
```
## Argument Reference
The following arguments are supported:
* `direction` - (Optional) The direction of the private zone resolver endpoint.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of the private zone resolver endpoint.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the private zone resolver endpoint.
* `status` - (Optional) The status of the private zone resolver endpoint.
* `tag_filters` - (Optional) List of tag filters.
* `vpc_id` - (Optional) The vpc ID of the private zone resolver endpoint.

The `tag_filters` object supports the following:

* `key` - (Optional) The key of the tag.
* `values` - (Optional) The values of the tag.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `endpoints` - The collection of query.
    * `created_at` - The created time of the endpoint.
    * `direction` - The direction of the endpoint.
    * `endpoint_id` - The endpoint id.
    * `id` - The id of the endpoint.
    * `ip_configs` - List of IP configurations.
        * `az_id` - The availability zone id of the endpoint.
        * `ip` - The IP address of the endpoint.
        * `subnet_id` - The subnet id of the endpoint.
    * `name` - The name of the endpoint.
    * `project_name` - The project name of the endpoint.
    * `security_group_id` - The security group id of the endpoint.
    * `status` - The status of the endpoint.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The updated time of the endpoint.
    * `vpc_id` - The vpc id of the endpoint.
    * `vpc_region` - The vpc region of the endpoint.
* `total_count` - The total count of query.


