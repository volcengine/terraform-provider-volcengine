---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_ssls"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_ssls"
description: |-
  Use this data source to query detailed information of rds postgresql instance ssls
---
# volcengine_rds_postgresql_instance_ssls
Use this data source to query detailed information of rds postgresql instance ssls
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_ssls" "example" {
  ids                  = ["postgres-72715e0d9f58", "postgres-0ac38a79fe35"]
  download_certificate = true
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Required) A list of the PostgreSQL instance IDs.
* `download_certificate` - (Optional) Whether to include SSL certificate raw bytes for each instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ssls` - The collection of query.
    * `address` - The protected addresses.
    * `certificate` - Raw byte stream array of certificate zip.
    * `force_encryption` - Whether to force encryption.
    * `instance_id` - The id of the postgresql Instance.
    * `is_valid` - Whether the SSL certificate is valid.
    * `ssl_enable` - Whether to enable SSL.
    * `ssl_expire_time` - The expiration time of the SSL certificate. The format is: yyyy-MM-ddTHH:mm:ss(UTC time).
    * `tls_version` - The supported TLS versions.
* `total_count` - The total count of query.


