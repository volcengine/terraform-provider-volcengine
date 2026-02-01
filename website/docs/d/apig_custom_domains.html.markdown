---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_custom_domains"
sidebar_current: "docs-volcengine-datasource-apig_custom_domains"
description: |-
  Use this data source to query detailed information of apig custom domains
---
# volcengine_apig_custom_domains
Use this data source to query detailed information of apig custom domains
## Example Usage
```hcl
data "volcengine_apig_custom_domains" "foo" {
  gateway_id = "gd13d8c6eq1emkiunq6p0"
  service_id = "sd142lm6kiaj519k4l640"
}
```
## Argument Reference
The following arguments are supported:
* `gateway_id` - (Optional) The id of api gateway.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `resource_type` - (Optional) The resource type of domain. Valid values: `Console`, `Ingress`.
* `service_id` - (Optional) The id of api gateway service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `custom_domains` - The collection of query.
    * `certificate_id` - The id of the certificate.
    * `comments` - The comments of the custom domain.
    * `create_time` - The create time of the custom domain.
    * `domain` - The custom domain of the api gateway service.
    * `id` - The id of the custom domain.
    * `protocol` - The protocol of the custom domain.
    * `resource_type` - The resource type of domain.
    * `service_id` - The id of the api gateway service.
    * `ssl_redirect` - Whether to redirect https.
    * `status` - The status of the custom domain.
    * `type` - The type of the domain.
    * `update_time` - The update time of the custom domain.
* `total_count` - The total count of query.


