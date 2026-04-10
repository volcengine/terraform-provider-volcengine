---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_gateway_endpoints"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_gateway_endpoints"
description: |-
  Use this data source to query detailed information of privatelink vpc gateway endpoints
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_privatelink_vpc_gateway_endpoints
Use this data source to query detailed information of privatelink vpc gateway endpoints
## Example Usage
```hcl
data "volcengine_vpc_gateway_endpoints" "default" {
  ids = ["gwep-273yuq6q7bgn47fap8squ****"]
}

data "volcengine_vpc_gateway_endpoints" "foo" {
  vpc_id       = "vpc-bp15zkdt37pq72zv****"
  name_regex   = "^acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  output_file = "vpc_gateway_endpoints_output"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_name` - (Optional) The name of the gateway endpoint.
* `ids` - (Optional) A list of gateway endpoint IDs.
* `name_regex` - (Optional) A Name Regex of gateway endpoint.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the gateway endpoint.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The id of the vpc.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `vpc_gateway_endpoints` - The collection of query.
    * `creation_time` - The create time of the gateway endpoint.
    * `description` - The description of the gateway endpoint.
    * `endpoint_id` - The id of the gateway endpoint.
    * `endpoint_name` - The name of the gateway endpoint.
    * `id` - The id of the gateway endpoint.
    * `project_name` - The project name of the gateway endpoint.
    * `service_id` - The id of the gateway endpoint service.
    * `service_name` - The name of the gateway endpoint service.
    * `status` - The status of the gateway endpoint.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the gateway endpoint.
    * `vpc_id` - The id of the vpc.
    * `vpc_policy` - The vpc policy of the gateway endpoint.


