---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_ca_certificates"
sidebar_current: "docs-volcengine-datasource-alb_ca_certificates"
description: |-
  Use this data source to query detailed information of alb ca certificates
---
# volcengine_alb_ca_certificates
Use this data source to query detailed information of alb ca certificates
## Example Usage
```hcl
data "volcengine_alb_ca_certificates" "foo" {
  ids = ["cert-1iidd2r9ii0hs74adhfeodxo1"]
}
```
## Argument Reference
The following arguments are supported:
* `ca_certificate_name` - (Optional) The name of the CA certificate.
* `ids` - (Optional) A list of CA certificate IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the CA certificate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `certificates` - The collection of CA certificates query.
    * `ca_certificate_id` - The ID of the CA certificate.
    * `ca_certificate_name` - The name of the CA certificate.
    * `certificate_type` - The type of the CA certificate.
    * `create_time` - The create time of the CA Certificate.
    * `description` - The description of the CA certificate.
    * `domain_name` - The domain name of the CA Certificate.
    * `expired_at` - The expire time of the CA Certificate.
    * `listeners` - The ID list of the CA Listener.
    * `project_name` - The ProjectName of the CA Certificate.
    * `status` - The status of the CA Certificate.
* `total_count` - The total count of query.


