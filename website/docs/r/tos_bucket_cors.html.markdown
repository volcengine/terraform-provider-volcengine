---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_cors"
sidebar_current: "docs-volcengine-resource-tos_bucket_cors"
description: |-
  Provides a resource to manage tos bucket cors
---
# volcengine_tos_bucket_cors
Provides a resource to manage tos bucket cors
## Example Usage
```hcl
resource "volcengine_tos_bucket" "foo" {
  bucket_name   = "tf-acc-test-bucket"
  public_acl    = "private"
  az_redundancy = "multi-az"
  project_name  = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_tos_bucket_cors" "foo" {
  bucket_name = volcengine_tos_bucket.foo.id
  cors_rules {
    allowed_origins = ["*"]
    allowed_methods = ["GET", "POST"]
    allowed_headers = ["Authorization"]
    expose_headers  = ["x-tos-request-id"]
    max_age_seconds = 1500
  }
  cors_rules {
    allowed_origins = ["*", "https://www.volcengine.com"]
    allowed_methods = ["POST", "PUT", "DELETE"]
    allowed_headers = ["Authorization"]
    expose_headers  = ["x-tos-request-id"]
    max_age_seconds = 2000
    response_vary   = true
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `cors_rules` - (Required) The CORS rules of the bucket.

The `cors_rules` object supports the following:

* `allowed_methods` - (Required) The list of HTTP methods that are allowed in a preflight request. Valid values: `PUT`, `POST`, `DELETE`, `GET`, `HEAD`.
* `allowed_origins` - (Required) The list of origins that are allowed to make requests to the bucket.
* `allowed_headers` - (Optional) The list of headers that are allowed in a preflight request.
* `expose_headers` - (Optional) The list of headers that are exposed in the response to a preflight request. It is recommended to add two expose headers, X-Tos-Request-Id and ETag.
* `max_age_seconds` - (Optional) The maximum amount of time that a preflight request can be cached. Unit: second. Default value: 3600.
* `response_vary` - (Optional) Indicates whether the bucket returns the 'Vary: Origin' header in the response to preflight requests. Default value: false.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketCors can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket_cors.default resource_id
```

