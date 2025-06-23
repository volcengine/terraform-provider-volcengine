---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_gateway_services"
sidebar_current: "docs-volcengine-datasource-apig_gateway_services"
description: |-
  Use this data source to query detailed information of apig gateway services
---
# volcengine_apig_gateway_services
Use this data source to query detailed information of apig gateway services
## Example Usage
```hcl
data "volcengine_apig_gateway_services" "foo" {
  gateway_id = "gd13d8c6eq1emkiunq6p0"
}
```
## Argument Reference
The following arguments are supported:
* `gateway_id` - (Optional) The gateway id of api gateway service.
* `name` - (Optional) The name of api gateway service. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The status of api gateway service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `gateway_services` - The collection of query.
    * `auth_spec` - The auth spec of the api gateway service.
        * `enable` - Whether the api gateway service enable auth.
    * `comments` - The comments of the api gateway service.
    * `create_time` - The create time of the api gateway service.
    * `custom_domains` - The custom domains of the api gateway service.
        * `domain` - The custom domain of the api gateway service.
        * `id` - The id of the custom domain.
    * `domains` - The domains of the api gateway service.
        * `domain` - The domain of the api gateway service.
        * `type` - The type of the domain.
    * `gateway_id` - The gateway id of the api gateway service.
    * `gateway_name` - The gateway name of the api gateway service.
    * `id` - The Id of the api gateway service.
    * `message` - The error message of the api gateway service.
    * `name` - The name of the api gateway service.
    * `protocol` - The protocol of the api gateway service.
    * `status` - The status of the api gateway service.
* `total_count` - The total count of query.


